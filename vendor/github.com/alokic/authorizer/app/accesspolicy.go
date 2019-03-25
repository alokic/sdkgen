package app

import (
	"context"

	"github.com/alokic/authorizer/app/proto"
	"github.com/alokic/gopkg/gid"
)

var (
	// AccessPolicyStateActive means svc is active.
	AccessPolicyStateActive = "active"
	// AccessPolicyStateBlocked means svc is blocked.
	AccessPolicyStateBlocked = "blocked"
)

// AccessPolicy is a domain object.
type AccessPolicy proto.AccessPolicy

// AccessPolicies is a list of domain objects.
type AccessPolicies []*AccessPolicy

// AccessPolicyRepository is the interface for the database.
type AccessPolicyRepository interface {
	Create(context.Context, *AccessPolicy) (*AccessPolicy, error)
	Get(context.Context, uint64) (*AccessPolicy, error)
	Update(context.Context, *AccessPolicy) (*AccessPolicy, error)
	Delete(context.Context, uint64) error
	List(context.Context, []uint64, uint64) ([]*AccessPolicy, error)
	ListByUserPermission(context.Context, uint64, uint64) ([]uint64, error)
}

// AccessPolicyCreateRequest is the definition of access policy create request.
type AccessPolicyCreateRequest proto.AccessPolicyCreateRequest

// AccessPolicyCreateReply is the definition of access policy create reply.
type AccessPolicyCreateReply proto.AccessPolicyCreateReply

// AccessPolicyListRequest is the request to list access policies.
type AccessPolicyListRequest proto.AccessPolicyListRequest

// AccessPolicyListReply is the reply to list access policies.
type AccessPolicyListReply proto.AccessPolicyListReply

// AccessPolicyUpdateRequest is the request to update access policies.
type AccessPolicyUpdateRequest proto.AccessPolicyUpdateRequest

// AccessPolicyUpdateReply is the response to update access policies.
type AccessPolicyUpdateReply proto.AccessPolicyUpdateReply

// AccessPolicyDeleteRequest is the request to delete access policies.
type AccessPolicyDeleteRequest proto.AccessPolicyDeleteRequest

// AccessPolicyDeleteReply is the reply for deleting access policies.
type AccessPolicyDeleteReply proto.AccessPolicyDeleteReply

// AccessPolicyBlockRequest is definition of block user request.
type AccessPolicyBlockRequest proto.AccessPolicyBlockRequest

// AccessPolicyBlockReply is definition of block reply.
type AccessPolicyBlockReply proto.AccessPolicyBlockReply

// AccessPolicyActivateRequest is definition of actiavte user request.
type AccessPolicyActivateRequest proto.AccessPolicyActivateRequest

// AccessPolicyActivateReply is definition of activate reply.
type AccessPolicyActivateReply proto.AccessPolicyActivateReply

// AccessPolicyService provides interface for access policy service.
type AccessPolicyService interface {
	Create(context.Context, *AccessPolicyCreateRequest) (*AccessPolicyCreateReply, error)
	Update(context.Context, *AccessPolicyUpdateRequest) (*AccessPolicyUpdateReply, error)
	List(context.Context, *AccessPolicyListRequest) (*AccessPolicyListReply, error)
	Delete(context.Context, *AccessPolicyDeleteRequest) (*AccessPolicyDeleteReply, error)
	Block(context.Context, *AccessPolicyBlockRequest) (*AccessPolicyBlockReply, error)
	Activate(context.Context, *AccessPolicyActivateRequest) (*AccessPolicyActivateReply, error)
}

// CreateAccessPolicy is to create an access policy domain object.
func CreateAccessPolicy(name string, id uint64) (*AccessPolicy, error) {
	if id == 0 {
		id, _ = gid.Get()
	}
	return &AccessPolicy{
		Id:    id,
		Name:  name,
		State: AccessPolicyStateActive,
	}, nil
}

// ToProto converts report domain object to protobuf definition.
func (a *AccessPolicy) ToProto() *proto.AccessPolicy {
	ap := proto.AccessPolicy(*a)
	return &ap
}

// ToProto converts list of report domain object to protobuf definition.
func (aa AccessPolicies) ToProto() []*proto.AccessPolicy {
	var ap []*proto.AccessPolicy
	for _, a := range aa {
		ap = append(ap, a.ToProto())
	}

	return ap
}
