package random

import (
	"testing"

	"github.com/nguyencatpham/go-micro/selector"
)

func TestRandom(t *testing.T) {
	selector.Tests(t, NewSelector())
}
