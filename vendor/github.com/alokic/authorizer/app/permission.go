package app

import (
	"context"

	"github.com/alokic/authorizer/app/proto"
	"github.com/alokic/gopkg/gid"
)

var (
	// PermissionStateActive means permission is active.
	PermissionStateActive = "active"
	// PermissionStateBlocked means permission is blocked.
	PermissionStateBlocked = "blocked"
)

// Permission is a domain object.
type Permission proto.Permission

// Permissions is a list of domain objects.
type Permissions []*Permission

// PermissionRepository is the interface for the database.
type PermissionRepository interface {
	Create(context.Context, *Permission) (*Permission, error)
	Get(context.Context, uint64) (*Permission, error)
	Update(context.Context, *Permission) (*Permission, error)
	ListAll(context.Context) ([]*Permission, error)
	ListInInterval(context.Context, int64) ([]*Permission, error)
	GetState(context.Context, uint64) (*Permission, *Svc, error)
	List(context.Context, uint64) ([]*Permission, error)
	ListByAccessPolicy(context.Context, uint64) ([]*Permission, error)
}

// PermissionCreateRequest is the definition of a create request.
type PermissionCreateRequest proto.PermissionCreateRequest

// PermissionCreateReply is the definition of create reply.
type PermissionCreateReply proto.PermissionCreateReply

// PermissionUpdateRequest is the definition of an update request.
type PermissionUpdateRequest proto.PermissionUpdateRequest

// PermissionUpdateReply is the definition of an update reply.
type PermissionUpdateReply proto.PermissionUpdateReply

// PermissionBlockRequest is the definition of a delete request.
type PermissionBlockRequest proto.PermissionBlockRequest

// PermissionBlockReply is the definition of a delete reply.
type PermissionBlockReply proto.PermissionBlockReply

// PermissionActivateRequest is definition of actiavte user request.
type PermissionActivateRequest proto.PermissionActivateRequest

// PermissionActivateReply is definition of activate reply.
type PermissionActivateReply proto.PermissionActivateReply

//PermissionListRequest is the definition of a list request
type PermissionListRequest proto.PermissionListRequest

//PermissionListReply is the definition of a list reply
type PermissionListReply proto.PermissionListReply

// PermissionService provides the interface for permission service.
type PermissionService interface {
	Create(context.Context, *PermissionCreateRequest) (*PermissionCreateReply, error)
	Update(context.Context, *PermissionUpdateRequest) (*PermissionUpdateReply, error)
	Block(context.Context, *PermissionBlockRequest) (*PermissionBlockReply, error)
	Activate(context.Context, *PermissionActivateRequest) (*PermissionActivateReply, error)
	List(context.Context, *PermissionListRequest) (*PermissionListReply, error)
}

// CreatePermission is a factory method to create a permission object.
func CreatePermission(name, path, action, description string, serviceid uint64) (*Permission, error) {
	id, _ := gid.Get()

	return &Permission{
		Id:          id,
		ServiceId:   serviceid,
		Name:        name,
		Path:        path,
		Action:      action,
		State:       PermissionStateActive,
		Description: description,
	}, nil
}

// ToProto converts permission domain object to the protobuf definition.
func (p *Permission) ToProto() *proto.Permission {
	pp := proto.Permission(*p)
	return &pp
}

// ToProto converts collection of permission domain objects to their protobuf definitions.
func (p Permissions) ToProto() []*proto.Permission {
	var pp []*proto.Permission
	for _, p := range p {
		pp = append(pp, p.ToProto())
	}

	return pp
}
