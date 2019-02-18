package cache

import (
	"time"

	"github.com/wayn3h0/gop/cache/dependency"
	"github.com/wayn3h0/gop/errors"
)

// Item represents a cache item.
type Item struct {
	key                     string
	value                   interface{}
	createdAt               time.Time
	accessedAt              time.Time
	absoluteExpirationTime  time.Time
	slidingExpirationPeriod time.Duration
	dependencies            []dependency.Dependency
}

// Key returns the key of item.
func (i *Item) Key() string {
	return i.key
}

// Value returns the value of item.
func (i *Item) Value() interface{} {
	return i.value
}

// CreatedAt returns the creation timestamp of item.
func (i *Item) CreatedAt() time.Time {
	return i.createdAt
}

// AccessedAt returns the last accessed timestamp of item.
func (i *Item) AccessedAt() time.Time {
	return i.accessedAt
}

// Access updates the last accessed timestamp.
func (i *Item) Access() {
	i.accessedAt = time.Now()
}

// SetValue sets the value of item.
func (i *Item) SetValue(v interface{}) *Item {
	i.value = v
	return i
}

// SetAbsoluteExpirationTime sets the absolute expiration time for item.
func (i *Item) SetAbsoluteExpirationTime(v time.Time) *Item {
	i.absoluteExpirationTime = v
	return i
}

// SetSlidingExpirationPeriod sets the sliding expiration duration for item.
func (i *Item) SetSlidingExpirationPeriod(v time.Duration) *Item {
	i.slidingExpirationPeriod = v
	return i
}

// SetDependencies sets the Dependencies for item.
func (i *Item) SetDependencies(deps ...dependency.Dependency) *Item {
	i.dependencies = deps
	return i
}

// HasExpired reports whether the item has expired.
func (i *Item) HasExpired() bool {
	if !i.absoluteExpirationTime.IsZero() { // check absolute expiration time
		if i.absoluteExpirationTime.Before(time.Now()) {
			return true
		}
	}
	if i.slidingExpirationPeriod > 0 { // check sliding expiration period
		if !i.accessedAt.IsZero() { // accessed
			if i.accessedAt.Add(i.slidingExpirationPeriod).Before(time.Now()) {
				return true
			}
		} else { // never accessed
			if i.createdAt.Add(i.slidingExpirationPeriod).Before(time.Now()) {
				return true
			}
		}
	}

	// check dependencies
	for _, dep := range i.dependencies {
		if dep.HasChanged() {
			return true
		}
	}

	return false
}

// NewItem returns a new item.
func NewItem(key string, value interface{}) (*Item, error) {
	if len(key) == 0 {
		return nil, errors.New("cache: key of item cannot be empty")
	}

	return &Item{
		key:       key,
		value:     value,
		createdAt: time.Now(),
	}, nil
}
