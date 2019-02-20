package concurrent

import (
	"sync"

	ctn "github.com/wayn3h0/gop/cache/container"
	"github.com/wayn3h0/gop/errors"
)

type container struct {
	Locker sync.RWMutex
	Inner  ctn.Container
}

func (c *container) Clear() error {
	c.Locker.Lock()
	defer c.Locker.Unlock()

	return c.Inner.Clear()
}

func (c *container) Remove(key string) error {
	c.Locker.Lock()
	defer c.Locker.Unlock()

	return c.Inner.Remove(key)
}

func (c *container) Save(key string, value interface{}) error {
	c.Locker.Lock()
	defer c.Locker.Unlock()

	return c.Inner.Save(key, value)
}

func (c *container) Get(key string) (interface{}, error) {
	c.Locker.RLock()
	defer c.Locker.RUnlock()

	return c.Inner.Get(key)
}

// NewContainer returns a new container for safe concurrent access.
func NewContainer(inner ctn.Container) (ctn.Container, error) {
	if inner == nil {
		return nil, errors.New("cache: inner container cannot be nil")
	}

	return &container{
		Inner: inner,
	}, nil
}
