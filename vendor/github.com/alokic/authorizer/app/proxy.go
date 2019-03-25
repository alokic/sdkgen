package app

import (
	"context"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/alokic/authorizer/app/proto"
)

// Proxy is domain object.
type Proxy proto.Proxy

// Proxys is list of domain objects.
type Proxys []*Proxy

// ProxyLoadRequest is definition of load permissions request.
type ProxyLoadRequest proto.ProxyLoadRequest

// ProxyLoadReply is definition of load permissions reply.
type ProxyLoadReply proto.ProxyLoadReply

// ProxyLoadAsyncRequest is definition of load permissions asynchronously request.
type ProxyLoadAsyncRequest proto.ProxyLoadAsyncRequest

// ProxyLoadAsyncReply is definition of load permissions asynchronously reply.
type ProxyLoadAsyncReply proto.ProxyLoadAsyncReply

// ProxyGetRequest is definition of get proxy request.
type ProxyGetRequest proto.ProxyGetRequest

// ProxyGetReply is definition of get proxy reply.
type ProxyGetReply struct {
	proto.ProxyGetReply
	Router http.Handler
}

// ProxyCachePool is interface to cache connect pool.
type ProxyCachePool interface {
	Get() ProxyCacheConn
}

// ProxyCacheConn is interface to cache connection.
type ProxyCacheConn = redis.Conn

// ProxyService provides interface for proxy service.
type ProxyService interface {
	Load(context.Context, *ProxyLoadRequest) (*ProxyLoadReply, error)
	LoadAsync(context.Context, *ProxyLoadAsyncRequest) (*ProxyLoadAsyncReply, error)
	Get(context.Context, *ProxyGetRequest) (*ProxyGetReply, error)
}

// ToProto converts report domain object to protobuf definition.
func (s *Proxy) ToProto() *proto.Proxy {
	report := proto.Proxy(*s)
	return &report
}

// ToProto converts list of report domain object to protobuf definition.
func (roles Proxys) ToProto() []*proto.Proxy {
	var pr []*proto.Proxy
	for _, r := range roles {
		pr = append(pr, r.ToProto())
	}
	return pr
}
