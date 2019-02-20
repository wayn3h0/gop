package multilevel

import (
	ctn "github.com/wayn3h0/gop/cache/container"
	"github.com/wayn3h0/gop/errors"
)

type container struct {
	List []ctn.Container
}

func (c *container) Clear() error {
	for _, v := range c.List {
		err := v.Clear()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *container) Remove(key string) error {
	for _, v := range c.List {
		err := v.Remove(key)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *container) Save(key string, value interface{}) error {
	for _, v := range c.List {
		err := v.Save(key, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *container) Get(key string) (interface{}, error) {
	for _, v := range c.List {
		item, err := v.Get(key)
		if err != nil {
			return nil, err
		}

		if item != nil {
			return item, nil
		}
	}

	return nil, nil
}

// NewContainer returns a new multi-level cache container by warpping given containers.
func NewContainer(containers ...ctn.Container) (ctn.Container, error) {
	if len(containers) == 0 {
		return nil, errors.New("cache: containers cannot be empty")
	} else {
		for i, v := range containers {
			if v == nil {
				return nil, errors.Newf("cache: cannot accept nil container [index: %d]", i)
			}
		}
	}

	return &container{
		List: containers,
	}, nil
}
