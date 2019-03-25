package app

import (
	"context"

	"github.com/alokic/authorizer/app/proto"
)

// UserRole is domain object.
type UserRole proto.UserRole

// UserRoles is a list of domain objects.
type UserRoles []*UserRole

// UserRoleRepository is interface for a database
type UserRoleRepository interface {
	BatchUpdate(context.Context, []*UserRole) ([]*UserRole, error)
}

// UserRolesUpdateRequest is definition of update request.
type UserRolesUpdateRequest proto.UserRolesUpdateRequest

// UserRolesUpdateReply is definition of update reply.
type UserRolesUpdateReply proto.UserRolesUpdateReply

// UserRoleService provides interface for user_role service.
type UserRoleService interface {
	BatchUpdate(context.Context, *UserRolesUpdateRequest) (*UserRolesUpdateReply, error)
}

// CreateUserRole is factory method to create user_role domain object
func CreateUserRole(roleid, userid uint64) *UserRole {
	return &UserRole{
		UserId: userid,
		RoleId: roleid,
	}
}

// ToProto converts user_role domain object to protobuf definition.
func (ur *UserRole) ToProto() *proto.UserRole {
	pur := proto.UserRole(*ur)
	return &pur
}

// ToProto converts a list of user_roles domain objects to protobuf definition.
func (urs UserRoles) ToProto() []*proto.UserRole {
	var purs []*proto.UserRole
	for _, ur := range urs {
		purs = append(purs, ur.ToProto())
	}
	return purs
}
