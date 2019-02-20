package fifo

import (
	ctn "github.com/wayn3h0/gop/cache/container"
	"github.com/wayn3h0/gop/cache/container/memory"
)

// NewContainer returns a new in-memory cache container using FIFO (first in first out) arithmetic.
func NewContainer(capacity int) ctn.Container {
	return &container{
		Capacity: capacity,
		list:     new(list).Initialize(),
	}
}

// register the container.
func init() {
	memory.FIFO.Register(NewContainer)
}
