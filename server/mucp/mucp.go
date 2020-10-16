// Package mucp provides a transport agnostic RPC server
package mucp

import (
	"github.com/nguyencatpham/go-micro/server"
)

var (
	DefaultRouter = newRpcRouter()
)

// NewServer returns a micro server interface
func NewServer(opts ...server.Option) server.Server {
	return newServer(opts...)
}
