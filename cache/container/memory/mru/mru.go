package mru

import (
	ctn "github.com/wayn3h0/gop/cache/container"
	"github.com/wayn3h0/gop/cache/container/memory"
)

// NewContainer returns a new in-memory cache Container using MRU (most recently used) arithmetic.
func NewContainer(capacity int) ctn.Container {
	return &container{
		Capacity: capacity,
		list:     new(list).Initialize(),
	}
}

// register the container.
func init() {
	memory.MRU.Register(NewContainer)
}
