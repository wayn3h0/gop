package memory

import (
	ctn "github.com/wayn3h0/gop/cache/container"
	"github.com/wayn3h0/gop/errors"
)

// Container represents a memory cache container that implemented in another package.
type Container byte

// Memory Containers.
const (
	// FIFO represents a memory container using replacement algorithm using FIFO (first in first out).
	FIFO Container = iota + 1

	// LFU represents a memory container using replacement algorithm using LFU (least frequently used).
	LFU

	// LRU represents a memory container using replacement algorithm using LRU (least recently used).
	LRU

	// MRU represents a memory container using replacement algorithm using MRU (most recently used).
	MRU

	// ARC represents a memory container using replacement algorithm using ARC (adaptive/adjustable replacement cache).
	ARC

	maxOfContainers
)

var containers = make([]func(int) ctn.Container, maxOfContainers)

// Register registers the container cache.
// This is intended to be called from the init function in packages that implement container functions.
func (c Container) Register(function func(int) ctn.Container) {
	if c <= 0 && c >= maxOfContainers {
		panic(errors.New("cache: register of unknown memory container function"))
	}

	containers[c] = function
}

// Available reports whether the given container is linked into the binary.
func (c Container) Available() bool {
	return c > 0 && c < maxOfContainers && containers[c] != nil
}

// NewContainer returns a new memory cache.
func (c Container) NewContainer(capacity int) ctn.Container {
	if !c.Available() {
		panic(errors.Newf("cache: requested memory container function #%i is unavailable", int(c)))
	}

	return containers[c](capacity)
}
