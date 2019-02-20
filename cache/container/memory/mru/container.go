package mru

// container represents a MRU caching container.
type container struct {
	Capacity int
	list     *list
}

func (c *container) Get(key string) (interface{}, error) {
	return c.list.Get(key), nil
}

func (c *container) Save(key string, value interface{}) error {
	if c.Capacity > 0 && c.list.Count() == c.Capacity && !c.list.Contains(key) {
		c.list.Discard()
	}

	c.list.Save(key, value)

	return nil
}

func (c *container) Remove(key string) error {
	c.list.Remove(key)

	return nil
}

func (c *container) Clear() error {
	c.list.Init()

	return nil
}
