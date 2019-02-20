package cache

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/wayn3h0/gop/cache/dependency"
	"github.com/wayn3h0/gop/errors"
)

// Item represents a cache item.
type Item struct {
	Key                     string
	Value                   interface{}
	CreatedAt               time.Time
	AccessedAt              time.Time
	AbsoluteExpirationTime  time.Time
	SlidingExpirationPeriod time.Duration
	Dependencies            []dependency.Dependency
}

// Access updates the last accessed timestamp.
func (i *Item) access() {
	i.AccessedAt = time.Now()
}

// SetAbsoluteExpiration sets the Absolute expiration for item.
func (i *Item) SetAbsoluteExpiration(Absolute time.Time) {
	i.AbsoluteExpirationTime = Absolute
}

// SetSlidingExpiration sets the Sliding expiration for item.
func (i *Item) SetSlidingExpiration(Sliding time.Duration) {
	i.SlidingExpirationPeriod = Sliding
}

// SetDependencies sets the Dependencies for item.
func (i *Item) SetDependencies(Dependencies ...dependency.Dependency) {
	i.Dependencies = Dependencies
}

// HasExpired reports whether the item has expired.
func (i *Item) HasExpired() bool {
	if !i.AbsoluteExpirationTime.IsZero() { // check absolute expiration time
		if i.AbsoluteExpirationTime.Before(time.Now()) {
			return true
		}
	}
	if i.SlidingExpirationPeriod > 0 { // check sliding expiration period
		if !i.AccessedAt.IsZero() { // accessed
			if i.AccessedAt.Add(i.SlidingExpirationPeriod).Before(time.Now()) {
				return true
			}
		} else { // never accessed
			if i.CreatedAt.Add(i.SlidingExpirationPeriod).Before(time.Now()) {
				return true
			}
		}
	}

	// check dependencies
	for _, dep := range i.Dependencies {
		if dep.HasChanged() {
			return true
		}
	}

	return false
}

// Marshal marshals the item to byte data by gob.
func (i *Item) MarshalGob() ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(i)
	if err != nil {
		return nil, errors.Wrap(err, "cache: could not marshal item to Gob data")
	}

	return buffer.Bytes(), nil
}

// Unmarshal unmarshals the item from byte data by gob.
func (i *Item) UnmarshalGob(data []byte) error {
	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	err := decoder.Decode(i)
	if err != nil {
		return errors.Wrap(err, "cache: could not unmarshal item from Gob data")
	}

	return nil
}

// registers types for gob encoding.
func init() {
	gob.Register(Item{})
}

// NewItem returns a new item.
func NewItem(key string, value interface{}) (*Item, error) {
	if len(key) == 0 {
		return nil, errors.New("cache: key of item cannot be empty")
	}

	return &Item{
		Key:       key,
		Value:     value,
		CreatedAt: time.Now(),
	}, nil
}

// MustNewItem is like as NewItem but panic if key is empty.
func MustNewItem(key string, value interface{}) *Item {
	item, err := NewItem(key, value)
	if err != nil {
		panic(err)
	}

	return item
}
