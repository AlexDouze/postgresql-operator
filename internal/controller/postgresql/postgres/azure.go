package postgres

import (
	"context"
	"fmt"
	"strings"
)

type azurepg struct {
	serverName string
	pg
}

const MinUserSplit = 1

func newAzurePG(postgres *pg) PG {
	splitUser := strings.Split(postgres.user, "@")
	serverName := ""

	if len(splitUser) > MinUserSplit {
		serverName = splitUser[1]
	}

	return &azurepg{
		serverName,
		*postgres,
	}
}

func (azpg *azurepg) CreateUserRole(ctx context.Context, role, password string, attributes *RoleAttributes) (string, error) {
	returnedRole, err := azpg.pg.CreateUserRole(ctx, role, password, attributes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s@%s", returnedRole, azpg.serverName), nil
}

func (azpg *azurepg) GetRoleForLogin(login string) string {
	splitUser := strings.Split(azpg.user, "@")
	if len(splitUser) > MinUserSplit {
		return splitUser[0]
	}

	return login
}

func (azpg *azurepg) CreateDB(ctx context.Context, dbname, role string) error {
	// Have to add the master role to the group role before we can transfer the database owner
	err := azpg.GrantRole(ctx, role, azpg.GetRoleForLogin(azpg.user), false)
	if err != nil {
		return err
	}

	return azpg.pg.CreateDB(ctx, dbname, role)
}
