package file

import (
	"context"

	"github.com/nguyencatpham/go-micro/store"
)

type dirKey struct{}

// WithDir sets the directory to store the files in
func WithDir(dir string) store.Option {
	return func(o *store.Options) {
		o.Context = context.WithValue(o.Context, dirKey{}, dir)
	}
}
