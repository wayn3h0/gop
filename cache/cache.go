package cache

import (
	"github.com/wayn3h0/gop/cache/container"
	"github.com/wayn3h0/gop/errors"
)

// Cache represents a cache manager.
type Cache struct {
	container container.Container
}

// Clear removes all items from cache.
func (c *Cache) Clear() error {
	err := c.container.Clear()
	if err != nil {
		return errors.Wrap(err, "cache: could not clear cache items")
	}

	return nil
}

// Remove removes the cache item by given key.
func (c *Cache) Remove(key string) error {
	if len(key) == 0 {
		return errors.New("cache: key of item cannot be empty")
	}

	err := c.container.Remove(key)
	if err != nil {
		return errors.Wrapf(err, "cache: could not remove item with key %q", key)
	}

	return nil
}

func (c *Cache) Save(item *Item) error {
	if item == nil {
		return errors.New("cache: item cannot be nil")
	}

	err := c.container.Save(item.Key, item)
	if err != nil {
		return errors.Wrapf(err, "cache: could not save cache item with key %q to container", item.Key)
	}

	return nil
}

// Get returns the cache item by given key.
// It returns nil if cache item has expired or not found.
func (c *Cache) Get(key string) (*Item, error) {
	if len(key) == 0 {
		return nil, errors.New("cache: key of item cannot be empty")
	}

	v, err := c.container.Get(key)
	if err != nil {
		return nil, errors.Wrapf(err, "cache: could not get item with key %q", key)
	}
	if v == nil {
		return nil, nil
	}
	item := v.(*Item)
	if item.HasExpired() {
		err := c.container.Remove(key)
		if err != nil {
			return nil, errors.Wrapf(err, "cache: could not remove expired item with key %q", key)
		}

		return nil, nil
	}
	item.access()                          // update last accessed time
	err = c.container.Save(item.Key, item) // save the item to container
	if err != nil {
		return nil, errors.Wrapf(err, "cache: could not update item timestamp with key %q", key)
	}

	return item, nil
}

// New returns a new cache.
func New(container container.Container) (*Cache, error) {
	if container == nil {
		return nil, errors.New("cache: container of cache cannot be nil")
	}

	return &Cache{
		container: container,
	}, nil
}
