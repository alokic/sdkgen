package app

import (
	"context"
	"errors"
	"regexp"

	"github.com/alokic/authorizer/app/proto"
	"github.com/alokic/gopkg/gid"
)

const (
	mobileRegex = "^(?:(?:\\+|0{0,2})\\d\\d(\\s*[\\-]\\s*)?|[0]?)?[789]\\d{9}$"
	emailRegex  = "^[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?$"
)

var (
	// ErrUserInvalidEmail means email is in invalid format(check regex) or blank.
	ErrUserInvalidEmail = errors.New("invalid email")
	// ErrUserInvalidMobile means mobile is in invalid format(check regex) or blank.
	ErrUserInvalidMobile = errors.New("invalid mobile")
	// ErrUserNameBlank means name is blank.
	ErrUserNameBlank = errors.New("name is blank")
)

var (
	// UserStateActive means user is active in system.
	UserStateActive = "active"
	// UserStateBlocked means user is blocked.
	UserStateBlocked = "blocked"
)

// User is domain object.
type User proto.User

// Users is list of domain objects.
type Users []*User

// UserRepository is interface for a user store
type UserRepository interface {
	Create(context.Context, *User) (*User, error)
	Get(context.Context, uint64) (*User, error)
	GetByExternalID(context.Context, uint64) (*User, error)
	GetByEmail(context.Context, string) (*User, error)
	List(context.Context, []uint64) ([]*User, error)
	Update(context.Context, *User) error
}

// UserCreateRequest is definition of create request.
type UserCreateRequest proto.UserCreateRequest

// UserCreateReply is definition of create reply.
type UserCreateReply proto.UserCreateReply

// UserListRequest is definition of list request.
type UserListRequest proto.UserListRequest

// UserListReply is definition of list reply.
type UserListReply proto.UserListReply

// UserGetRequest is definition of get request.
type UserGetRequest proto.UserGetRequest

// UserGetReply is definition of get reply.
type UserGetReply proto.UserGetReply

// UserBlockRequest is definition of block user request.
type UserBlockRequest proto.UserBlockRequest

// UserBlockReply is definition of block reply.
type UserBlockReply proto.UserBlockReply

// UserActivateRequest is definition of actiavte user request.
type UserActivateRequest proto.UserActivateRequest

// UserActivateReply is definition of activate reply.
type UserActivateReply proto.UserActivateReply

// UserSignInRequest is definition of sign_in request.
type UserSignInRequest proto.UserSignInRequest

// UserSignInReply is definition of sign_in request.
type UserSignInReply proto.UserSignInReply

// UserSignOutRequest is definition of sign_out reply.
type UserSignOutRequest proto.UserSignOutRequest

// UserSignOutReply is definition of sign_out reply.
type UserSignOutReply proto.UserSignOutReply

// UserIsAllowedRequest is definition of is_allowed reply.
type UserIsAllowedRequest proto.UserIsAllowedRequest

// UserIsAllowedReply is definition of is_allowed reply.
type UserIsAllowedReply proto.UserIsAllowedReply

// UserService provides interface for user service.
type UserService interface {
	Create(context.Context, *UserCreateRequest) (*UserCreateReply, error)
	List(context.Context, *UserListRequest) (*UserListReply, error)
	Get(context.Context, *UserGetRequest) (*UserGetReply, error)
	Block(context.Context, *UserBlockRequest) (*UserBlockReply, error)
	Activate(context.Context, *UserActivateRequest) (*UserActivateReply, error)
	SignIn(context.Context, *UserSignInRequest) (*UserSignInReply, error)
	SignOut(context.Context, *UserSignOutRequest) (*UserSignOutReply, error)
	IsAllowed(context.Context, *UserIsAllowedRequest) (*UserIsAllowedReply, error)
}

// CreateUser is factory method to create user domain object
func CreateUser(au *User) (*User, error) {
	if au.Email != "" && !valid("email", au.Email) {
		return nil, ErrUserInvalidEmail
	}

	if !valid("name", au.Name) {
		return nil, ErrUserNameBlank
	}

	id, _ := gid.Get()
	return &User{
		Id:         id,
		Name:       au.Name,
		Email:      au.Email,
		ExternalId: au.ExternalId,
		State:      UserStateActive,
	}, nil
}

// ToProto converts user domain object to protobuf definition.
func (u *User) ToProto() *proto.User {
	user := proto.User(*u)
	return &user
}

// ToProto converts list of user domain object to protobuf definition.
func (us Users) ToProto() []*proto.User {
	var ps []*proto.User
	for _, u := range us {
		ps = append(ps, u.ToProto())
	}
	return ps
}

func valid(key, val string) bool {
	switch key {
	case "mobile":
		b, _ := regexp.MatchString(mobileRegex, val)
		return b
	case "email":
		b, _ := regexp.MatchString(emailRegex, val)
		return b
	case "name":
		if val != "" {
			return true
		}
	}
	return false
}
