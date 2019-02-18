package container

// Container represents a cache container where store the data.
type Container interface {
	// Clear removes all items.
	Clear() error

	// Remove removes the item by given key.
	Remove(key string) error

	// Save inserts/updates the item.
	Save(key string, value interface{}) error

	// Get returns the item by given key.
	Get(key string) (interface{}, error)
}
