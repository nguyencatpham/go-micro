package http

import (
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	"github.com/nguyencatpham/go-micro/v3/registry"
	"github.com/nguyencatpham/go-micro/v3/registry/memory"
	"github.com/nguyencatpham/go-micro/v3/router"
	regRouter "github.com/nguyencatpham/go-micro/v3/router/registry"
)

func TestRoundTripper(t *testing.T) {
	m := memory.NewRegistry()
	r := regRouter.NewRouter(router.Registry(m))

	rt := NewRoundTripper(WithRouter(r))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`hello world`))
	})

	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	go http.Serve(l, nil)

	m.Register(&registry.Service{
		Name: "example.com",
		Nodes: []*registry.Node{
			{
				Id:      "1",
				Address: l.Addr().String(),
			},
		},
	})

	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	w, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	w.Body.Close()

	if string(b) != "hello world" {
		t.Fatal("response is", string(b))
	}

	// test http request
	c := &http.Client{
		Transport: rt,
	}

	rsp, err := c.Get("http://example.com")
	if err != nil {
		t.Fatal(err)
	}

	b, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		t.Fatal(err)
	}
	rsp.Body.Close()

	if string(b) != "hello world" {
		t.Fatal("response is", string(b))
	}

}
