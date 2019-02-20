package arc

import (
	ctn "github.com/wayn3h0/gop/cache/container"
	"github.com/wayn3h0/gop/cache/container/memory"
)

// NewContainer returns a new in-memory cache container using ARC (adaptive/adjustable replacement cache) arithmetic.
func NewContainer(capacity int) ctn.Container {
	return &container{
		Capacity: capacity,
		p:        0,
		t1:       new(list).Initialize(),
		t2:       new(list).Initialize(),
		b1:       new(list).Initialize(),
		b2:       new(list).Initialize(),
	}
}

// register the container.
func init() {
	memory.ARC.Register(NewContainer)
}
