package memory

import "github.com/nguyencatpham/go-micro/store"

// Options which are used to configure the in-memory stream
type Options struct {
	Store store.Store
}

// Option is a function which configures options
type Option func(o *Options)

// Store sets the store to use
func Store(s store.Store) Option {
	return func(o *Options) {
		o.Store = s
	}
}
