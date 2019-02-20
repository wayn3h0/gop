package arc

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// container represents a ARC cache container.
type container struct {
	Capacity int
	p        int // target size of t1
	t1       *list
	t2       *list
	b1       *list
	b2       *list
}

func (c *container) replace() {
	if c.t1.Count() >= max(1, c.p) { // t1's size exceeds target (t1 is too big)
		// grab from t1 and put to b1
		if key, val := c.t1.Discard(); len(key) > 0 {
			c.b1.Save(key, val)
		}
	} else {
		// grab from t2 and put to b2
		if key, val := c.t2.Discard(); len(key) > 0 {
			c.b2.Save(key, val)
		}
	}
}

func (c *container) Get(key string) (interface{}, error) {
	if c.t1.Contains(key) { // seen twice recently, put it to t2
		val := c.t1.Remove(key)
		c.t2.Save(key, val)
		return val, nil
	}

	if c.t2.Contains(key) {
		return c.t2.Get(key), nil
	}

	if c.b1.Contains(key) {
		c.p = min(c.Capacity, c.p+max(c.b2.Count()/c.b1.Count(), 1)) // adapt the target size of t1
		c.replace()
		val := c.b1.Remove(key)
		c.t2.Save(key, val) // seen twice recently, put it to t2
		return val, nil
	}

	if c.b2.Contains(key) {
		c.p = max(0, c.p-max(c.b1.Count()/c.b2.Count(), 1)) // adapt the target size of t1
		c.replace()
		val := c.b2.Remove(key)
		c.t2.Save(key, val) // seen twice recently, put it to t2
		return val, nil
	}

	return nil, nil
}

func (c *container) Save(key string, value interface{}) error {
	// remove the item anyway
	c.Remove(key)

	if c.t1.Count()+c.b1.Count() == c.Capacity { // b1 + t1 is full
		if c.t1.Count() < c.Capacity { // still room in t1
			c.b1.Discard()
			c.replace()
		} else {
			c.t1.Discard()
		}
	} else { //c.t1.Count()+c.b1.Count() < c.Capacity {
		total := c.t1.Count() + c.t2.Count() + c.b1.Count() + c.b2.Count()
		if total >= c.Capacity { // cache full
			if total == 2*c.Capacity {
				c.b2.Discard()
			}

			c.replace()
		}
	}

	c.t1.Save(key, value) // seen once recently, put it to t1

	return nil
}

func (c *container) Remove(key string) error {
	c.t1.Remove(key)
	c.t2.Remove(key)
	c.b1.Remove(key)
	c.b2.Remove(key)

	return nil
}

func (c *container) Clear() error {
	c.p = 0
	c.t1.Init()
	c.t2.Init()
	c.b1.Init()
	c.b2.Init()

	return nil
}
