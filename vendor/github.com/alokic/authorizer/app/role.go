package app

import (
	"context"

	"github.com/alokic/authorizer/app/proto"
	"github.com/alokic/gopkg/gid"
)

var (
	// RoleStateActive means user is active in system.
	RoleStateActive = "active"
	// RoleStateBlocked means user is blocked.
	RoleStateBlocked = "blocked"
	// RoleSuperUser has admin + authorizer-admin privileges.
	RoleSuperUser = "superuser"
)

// Role is domain object.
type Role proto.Role

// Roles is list of domain object.
type Roles []*Role

// RoleRepository is interface for a database
type RoleRepository interface {
	Create(context.Context, *Role) (*Role, error)
	Get(context.Context, uint64) (*Role, error)
	Update(context.Context, *Role) (*Role, error)
	List(context.Context, []uint64, uint64) ([]*Role, error)
	ListRoles(context.Context, []uint64) ([]*Role, error)
	ListUserRoles(context.Context, uint64) ([]*Role, error)
}

// RoleCreateRequest is definition of create request.
type RoleCreateRequest proto.RoleCreateRequest

// RoleCreateReply is definition of create reply.
type RoleCreateReply proto.RoleCreateReply

// RoleListRequest is definition of list request.
type RoleListRequest proto.RoleListRequest

// RoleListReply is definition of list reply.
type RoleListReply proto.RoleListReply

// RoleGetRequest is definition of get request.
type RoleGetRequest proto.RoleGetRequest

// RoleGetReply is definition of get reply.
type RoleGetReply proto.RoleGetReply

// RoleUpdateRequest is definition of update request.
type RoleUpdateRequest proto.RoleUpdateRequest

// RoleUpdateReply is definition of update reply.
type RoleUpdateReply proto.RoleUpdateReply

// RoleBlockRequest is definition of block user request.
type RoleBlockRequest proto.RoleBlockRequest

// RoleBlockReply is definition of block reply.
type RoleBlockReply proto.RoleBlockReply

// RoleActivateRequest is definition of actiavte user request.
type RoleActivateRequest proto.RoleActivateRequest

// RoleActivateReply is definition of activate reply.
type RoleActivateReply proto.RoleActivateReply

// RoleService provides interface for role service.
type RoleService interface {
	Create(context.Context, *RoleCreateRequest) (*RoleCreateReply, error)
	List(ctx context.Context, request *RoleListRequest) (*RoleListReply, error)
	Get(ctx context.Context, request *RoleGetRequest) (*RoleGetReply, error)
	Block(context.Context, *RoleBlockRequest) (*RoleBlockReply, error)
	Activate(context.Context, *RoleActivateRequest) (*RoleActivateReply, error)
}

// CreateRole is factory method to create role domain object
func CreateRole(name, description string, serviceID uint64) *Role {
	id, _ := gid.Get()
	return &Role{
		Id:          id,
		Name:        name,
		Description: description,
		ServiceId:   serviceID,
		State:       RoleStateActive,
	}
}

// ToProto converts role domain object to protobuf definition.
func (r *Role) ToProto() *proto.Role {
	role := proto.Role(*r)
	return &role
}

// ToProto converts role domain object to protobuf definition.
func (roles Roles) ToProto() []*proto.Role {
	var pr []*proto.Role
	for _, r := range roles {
		pr = append(pr, r.ToProto())
	}
	return pr
}
