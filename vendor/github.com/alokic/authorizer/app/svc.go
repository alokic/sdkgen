package app

import (
	"context"

	"github.com/alokic/authorizer/app/proto"
	"github.com/alokic/gopkg/gid"
)

var (
	// SvcStateActive means svc is active.
	SvcStateActive = "active"
	// SvcStateBlocked means svc is blocked.
	SvcStateBlocked = "blocked"
	// SvcCookieAuthentication means service uses cookie for authentication token.
	SvcCookieAuthentication = "cookie"
	// SvcHeaderAuthentication means service uses headers for authentication token.
	SvcHeaderAuthentication = "header"
)

// Svc is domain object.
type Svc proto.Svc

// Svcs is list of domain object.
type Svcs []*Svc

// SvcRepository is interface for a database.
type SvcRepository interface {
	Create(context.Context, *Svc) (*Svc, error)
	Update(context.Context, *Svc) (*Svc, error)
	Get(context.Context, uint64) (*Svc, error)
	ListAll(context.Context) ([]*Svc, error)
}

// SvcBlockRequest is definition of block request.
type SvcBlockRequest proto.SvcBlockRequest

// SvcBlockReply is definition of block reply.
type SvcBlockReply proto.SvcBlockReply

// SvcActivateRequest is definition of actiavte user request.
type SvcActivateRequest proto.SvcActivateRequest

// SvcActivateReply is definition of activate reply.
type SvcActivateReply proto.SvcActivateReply

// SvcService provides interface for svc service.
type SvcService interface {
	Block(ctx context.Context, request *SvcBlockRequest) (*SvcBlockReply, error)
	Activate(context.Context, *SvcActivateRequest) (*SvcActivateReply, error)
}

// CreateSvc is a factory method to create a permission object.
func CreateSvc(name, domain string) (*Svc, error) {
	id, _ := gid.Get()

	return &Svc{
		Id:     id,
		Name:   name,
		Domain: domain,
		State:  SvcStateActive,
	}, nil
}

// ToProto converts permission domain object to the protobuf definition.
func (p *Svc) ToProto() *proto.Svc {
	pp := proto.Svc(*p)
	return &pp
}
