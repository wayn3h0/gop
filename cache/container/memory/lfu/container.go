package lfu

// container represents a LFU caching container.
type container struct {
	Capacity int
	heap     *heap
}

func (c *container) Get(key string) (interface{}, error) {
	return c.heap.Get(key), nil
}

func (c *container) Save(key string, value interface{}) error {
	if c.Capacity > 0 && c.heap.Count() == c.Capacity && !c.heap.Contains(key) {
		c.heap.Discard()
	}

	c.heap.Save(key, value)

	return nil
}

func (c *container) Remove(key string) error {
	c.heap.Remove(key)

	return nil
}

func (c *container) Clear() error {
	c.heap.Initialize()

	return nil
}
