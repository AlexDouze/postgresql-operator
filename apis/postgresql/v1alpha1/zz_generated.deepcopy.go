//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/easymile/postgresql-operator/apis/postgresql/common"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseModulesList) DeepCopyInto(out *DatabaseModulesList) {
	*out = *in
	if in.List != nil {
		in, out := &in.List, &out.List
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseModulesList.
func (in *DatabaseModulesList) DeepCopy() *DatabaseModulesList {
	if in == nil {
		return nil
	}
	out := new(DatabaseModulesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlDatabase) DeepCopyInto(out *PostgresqlDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlDatabase.
func (in *PostgresqlDatabase) DeepCopy() *PostgresqlDatabase {
	if in == nil {
		return nil
	}
	out := new(PostgresqlDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlDatabaseList) DeepCopyInto(out *PostgresqlDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgresqlDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlDatabaseList.
func (in *PostgresqlDatabaseList) DeepCopy() *PostgresqlDatabaseList {
	if in == nil {
		return nil
	}
	out := new(PostgresqlDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlDatabaseSpec) DeepCopyInto(out *PostgresqlDatabaseSpec) {
	*out = *in
	in.Schemas.DeepCopyInto(&out.Schemas)
	in.Extensions.DeepCopyInto(&out.Extensions)
	if in.EngineConfiguration != nil {
		in, out := &in.EngineConfiguration, &out.EngineConfiguration
		*out = new(common.CRLink)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlDatabaseSpec.
func (in *PostgresqlDatabaseSpec) DeepCopy() *PostgresqlDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresqlDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlDatabaseStatus) DeepCopyInto(out *PostgresqlDatabaseStatus) {
	*out = *in
	out.Roles = in.Roles
	if in.Schemas != nil {
		in, out := &in.Schemas, &out.Schemas
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlDatabaseStatus.
func (in *PostgresqlDatabaseStatus) DeepCopy() *PostgresqlDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(PostgresqlDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlEngineConfiguration) DeepCopyInto(out *PostgresqlEngineConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlEngineConfiguration.
func (in *PostgresqlEngineConfiguration) DeepCopy() *PostgresqlEngineConfiguration {
	if in == nil {
		return nil
	}
	out := new(PostgresqlEngineConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlEngineConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlEngineConfigurationList) DeepCopyInto(out *PostgresqlEngineConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgresqlEngineConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlEngineConfigurationList.
func (in *PostgresqlEngineConfigurationList) DeepCopy() *PostgresqlEngineConfigurationList {
	if in == nil {
		return nil
	}
	out := new(PostgresqlEngineConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlEngineConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlEngineConfigurationSpec) DeepCopyInto(out *PostgresqlEngineConfigurationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlEngineConfigurationSpec.
func (in *PostgresqlEngineConfigurationSpec) DeepCopy() *PostgresqlEngineConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresqlEngineConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlEngineConfigurationStatus) DeepCopyInto(out *PostgresqlEngineConfigurationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlEngineConfigurationStatus.
func (in *PostgresqlEngineConfigurationStatus) DeepCopy() *PostgresqlEngineConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(PostgresqlEngineConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUser) DeepCopyInto(out *PostgresqlUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUser.
func (in *PostgresqlUser) DeepCopy() *PostgresqlUser {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserList) DeepCopyInto(out *PostgresqlUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgresqlUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserList.
func (in *PostgresqlUserList) DeepCopy() *PostgresqlUserList {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserRole) DeepCopyInto(out *PostgresqlUserRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserRole.
func (in *PostgresqlUserRole) DeepCopy() *PostgresqlUserRole {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlUserRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserRoleList) DeepCopyInto(out *PostgresqlUserRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgresqlUserRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserRoleList.
func (in *PostgresqlUserRoleList) DeepCopy() *PostgresqlUserRoleList {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlUserRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserRolePrivilege) DeepCopyInto(out *PostgresqlUserRolePrivilege) {
	*out = *in
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(common.CRLink)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserRolePrivilege.
func (in *PostgresqlUserRolePrivilege) DeepCopy() *PostgresqlUserRolePrivilege {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserRolePrivilege)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserRoleSpec) DeepCopyInto(out *PostgresqlUserRoleSpec) {
	*out = *in
	if in.Privileges != nil {
		in, out := &in.Privileges, &out.Privileges
		*out = make([]*PostgresqlUserRolePrivilege, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PostgresqlUserRolePrivilege)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserRoleSpec.
func (in *PostgresqlUserRoleSpec) DeepCopy() *PostgresqlUserRoleSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserRoleStatus) DeepCopyInto(out *PostgresqlUserRoleStatus) {
	*out = *in
	if in.OldPostgresRoles != nil {
		in, out := &in.OldPostgresRoles, &out.OldPostgresRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserRoleStatus.
func (in *PostgresqlUserRoleStatus) DeepCopy() *PostgresqlUserRoleStatus {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserSpec) DeepCopyInto(out *PostgresqlUserSpec) {
	*out = *in
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(common.CRLink)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserSpec.
func (in *PostgresqlUserSpec) DeepCopy() *PostgresqlUserSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserStatus) DeepCopyInto(out *PostgresqlUserStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserStatus.
func (in *PostgresqlUserStatus) DeepCopy() *PostgresqlUserStatus {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusPostgresRoles) DeepCopyInto(out *StatusPostgresRoles) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusPostgresRoles.
func (in *StatusPostgresRoles) DeepCopy() *StatusPostgresRoles {
	if in == nil {
		return nil
	}
	out := new(StatusPostgresRoles)
	in.DeepCopyInto(out)
	return out
}
