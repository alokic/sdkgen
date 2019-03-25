package app

import (
	"context"
	"github.com/alokic/authorizer/app/proto"
)

// AccessPlcyPerm is a domain object
type AccessPlcyPerm proto.AccessPolicyPermission

// AccessPlcyPerms is a collection of domain objects
type AccessPlcyPerms []*AccessPlcyPerm

// AccessPlcyPermUpdateRequest is the request to batch update AccessPolicyPermission
type AccessPlcyPermUpdateRequest proto.AccessPolicyPermissionsUpdateRequest

// AccessPlcyPermUpdateReply is the response to batch update AccessPolicyPermissions
type AccessPlcyPermUpdateReply proto.AccessPolicyPermissionsUpdateReply

// AccessPolicyPermsService provides interface for access policy service
type AccessPolicyPermsService interface {
	BatchUpdate(context.Context, *AccessPlcyPermUpdateRequest) (
		*AccessPlcyPermUpdateReply, error,
	)
}

// AccessPlcyPermsRepository provides the interface for the database
type AccessPlcyPermsRepository interface {
	BatchUpdate(context.Context, []*AccessPlcyPerm, uint64) ([]*AccessPlcyPerm, error)
}

// CreateAccessPolicyPermission composes the AccessPlcyPerm domain object
func CreateAccessPolicyPermission(pID uint64, apID uint64) *AccessPlcyPerm {
	return &AccessPlcyPerm{
		AccessPolicyId: apID,
		PermissionId:   pID,
	}
}

// ToProto converts the domain object to the protobuf definition
func (ap *AccessPlcyPerm) ToProto() *proto.AccessPolicyPermission {
	pa := proto.AccessPolicyPermission(*ap)
	return &pa
}

// ToProto converts a collection of domain objects to their protobuf definitions
func (app AccessPlcyPerms) ToProto() []*proto.AccessPolicyPermission {
	var pa []*proto.AccessPolicyPermission
	for _, ap := range app {
		pa = append(pa, ap.ToProto())
	}

	return pa
}
