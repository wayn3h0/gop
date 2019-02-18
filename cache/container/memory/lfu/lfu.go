package lfu

import (
	ctn "github.com/wayn3h0/gop/cache/container"
	"github.com/wayn3h0/gop/cache/container/memory"
)

// NewContainer returns a new in-memory cache container using LFU (least frequently used) arithmetic.
func NewContainer(capacity int) ctn.Container {
	return &container{
		Capacity: capacity,
		heap:     new(heap).Initialize(),
	}
}

// register the container.
func init() {
	memory.LFU.Register(NewContainer)
}
