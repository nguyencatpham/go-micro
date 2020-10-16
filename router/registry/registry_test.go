package registry

import (
	"os"
	"testing"

	"github.com/nguyencatpham/go-micro/registry/memory"
	"github.com/nguyencatpham/go-micro/router"
)

func routerTestSetup() router.Router {
	r := memory.NewRegistry()
	return NewRouter(router.Registry(r))
}

func TestRouterClose(t *testing.T) {
	r := routerTestSetup()

	if err := r.Close(); err != nil {
		t.Errorf("failed to stop router: %v", err)
	}
	if len(os.Getenv("IN_TRAVIS_CI")) == 0 {
		t.Logf("TestRouterStartStop STOPPED")
	}
}
