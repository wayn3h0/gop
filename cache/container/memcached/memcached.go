package memcached

import (
	"strings"

	"github.com/wayn3h0/gop/cache"
	ctn "github.com/wayn3h0/gop/cache/container"
	"github.com/wayn3h0/gop/errors"

	"github.com/rainycape/memcache"
)

type container struct {
	*memcache.Client
}

func (c *container) Clear() error {
	err := c.Client.Flush(0)
	if err != nil {
		return errors.Wrap(err, "memcached: could not clear container")
	}

	return nil
}

func (c *container) Remove(key string) error {
	err := c.Client.Delete(key)
	if err != nil && err != memcache.ErrCacheMiss {
		return errors.Wrapf(err, "memcached: could not remove item with key %q from container", key)
	}

	return nil
}

func (c *container) Save(key string, value interface{}) error {
	item := value.(*cache.Item)
	data, err := item.MarshalGob()
	if err != nil {
		return err
	}

	mci := &memcache.Item{
		Key:   key,
		Value: data,
	}
	err = c.Client.Set(mci)
	if err != nil {
		return errors.Wrapf(err, "memcached: could not save (insert/update) item with key %q to container", item.Key)
	}

	return nil
}

func (c *container) Get(key string) (interface{}, error) {
	mci, err := c.Client.Get(key)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return nil, nil
		}

		return nil, errors.Wrapf(err, "memcached: could not get item with key %s from container", key)
	}

	var item cache.Item
	err = item.UnmarshalGob(mci.Value)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// NewContainer returns a new memcached cache container.
func NewContainer(servers ...string) (ctn.Container, error) {
	client, err := memcache.New(servers...)
	if err != nil {
		return nil, errors.Wrapf(err, "memcached: could not connect to servers %q", strings.Join(servers, ","))
	}

	return &container{
		Client: client,
	}, nil
}

// Short to NewContainer func.
func New(servers ...string) (ctn.Container, error) {
	return NewContainer(servers...)
}
