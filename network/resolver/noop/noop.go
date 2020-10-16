// Package noop is a noop resolver
package noop

import (
	"github.com/nguyencatpham/go-micro/network/resolver"
)

type Resolver struct{}

// Resolve returns the list of nodes
func (r *Resolver) Resolve(name string) ([]*resolver.Record, error) {
	return []*resolver.Record{}, nil
}
