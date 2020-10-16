package client_test

import (
	"context"
	"testing"

	"github.com/nguyencatpham/go-micro/broker"
	bmemory "github.com/nguyencatpham/go-micro/broker/memory"
	"github.com/nguyencatpham/go-micro/client"
	"github.com/nguyencatpham/go-micro/client/grpc"
	tmemory "github.com/nguyencatpham/go-micro/network/transport/memory"
	rmemory "github.com/nguyencatpham/go-micro/registry/memory"
	"github.com/nguyencatpham/go-micro/router"
	rtreg "github.com/nguyencatpham/go-micro/router/registry"
	"github.com/nguyencatpham/go-micro/server"
	grpcsrv "github.com/nguyencatpham/go-micro/server/grpc"
	cw "github.com/nguyencatpham/go-micro/util/client"
)

type TestFoo struct {
}

type TestReq struct{}

type TestRsp struct {
	Data string
}

func (h *TestFoo) Bar(ctx context.Context, req *TestReq, rsp *TestRsp) error {
	rsp.Data = "pass"
	return nil
}

func TestStaticClient(t *testing.T) {
	var err error

	req := grpc.NewClient().NewRequest(
		"go.micro.service.foo",
		"TestFoo.Bar",
		&TestReq{},
		client.WithContentType("application/json"),
	)
	rsp := &TestRsp{}

	reg := rmemory.NewRegistry()
	brk := bmemory.NewBroker(broker.Registry(reg))
	tr := tmemory.NewTransport()
	rtr := rtreg.NewRouter(router.Registry(reg))

	srv := grpcsrv.NewServer(
		server.Broker(brk),
		server.Registry(reg),
		server.Name("go.micro.service.foo"),
		server.Address("127.0.0.1:0"),
		server.Transport(tr),
	)
	if err = srv.Handle(srv.NewHandler(&TestFoo{})); err != nil {
		t.Fatal(err)
	}

	if err = srv.Start(); err != nil {
		t.Fatal(err)
	}

	cli := grpc.NewClient(
		client.Router(rtr),
		client.Broker(brk),
		client.Transport(tr),
	)

	w1 := cw.Static("xxx_localhost:12345", cli)
	if err = w1.Call(context.TODO(), req, nil); err == nil {
		t.Fatal("address xxx_#localhost:12345 must not exists and call must be failed")
	}

	w2 := cw.Static(srv.Options().Address, cli)
	if err = w2.Call(context.TODO(), req, rsp); err != nil {
		t.Fatal(err)
	} else if rsp.Data != "pass" {
		t.Fatalf("something wrong with response: %#+v", rsp)
	}
}
