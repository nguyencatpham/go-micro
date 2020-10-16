package http

import (
	"errors"
	"net/http"

	"github.com/nguyencatpham/go-micro/selector"
)

type roundTripper struct {
	rt   http.RoundTripper
	st   selector.Selector
	opts Options
}

func (r *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	routes, err := r.opts.Router.Lookup(req.URL.Host)
	if err != nil {
		return nil, err
	}

	// rudimentary retry 3 times
	for _, route := range routes {
		req.URL.Host = route.Address
		w, err := r.rt.RoundTrip(req)
		if err != nil {
			continue
		}
		return w, nil
	}

	return nil, errors.New("failed request")
}
