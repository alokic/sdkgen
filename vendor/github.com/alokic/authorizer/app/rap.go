package app

import (
	"context"
	"github.com/alokic/authorizer/app/proto"
)

// RolePlcyUpdateRequest is the definition of an update request.
type RolePlcyUpdateRequest proto.RoleAccessPoliciesUpdateRequest

// RolePlcyUpdateReply is the definition of an update reply.
type RolePlcyUpdateReply proto.RoleAccessPoliciesUpdateReply

// RoleAccessPolicy is the domain object.
type RoleAccessPolicy proto.RoleAccessPolicy

// RoleAccessPolicies is the collection of domain objects.
type RoleAccessPolicies []*RoleAccessPolicy

// RolePlcyRepository is the interface for the database.
type RolePlcyRepository interface {
	BatchUpdate(context.Context, []*RoleAccessPolicy, uint64) ([]*RoleAccessPolicy, error)
}

// RolePlcyService is the interface for the service.
type RolePlcyService interface {
	BatchUpdate(context.Context, *RolePlcyUpdateRequest) (*RolePlcyUpdateReply, error)
}

// CreateRoleAccessPolicy composes the domain object.
func CreateRoleAccessPolicy(apID uint64, rID uint64) *RoleAccessPolicy {
	return &RoleAccessPolicy{RoleId: rID, AccessPolicyId: apID}
}

// ToProto converts domain object to the protobuf definition.
func (r *RoleAccessPolicy) ToProto() *proto.RoleAccessPolicy {
	pr := proto.RoleAccessPolicy(*r)
	return &pr
}

// ToProto converts a collection of domain objects to protobuf definition.
func (rr RoleAccessPolicies) ToProto() []*proto.RoleAccessPolicy {
	var pr []*proto.RoleAccessPolicy

	{
		for _, r := range rr {
			pr = append(pr, r.ToProto())
		}
	}
	return pr
}
