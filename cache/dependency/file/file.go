package file

import (
	"os"
	"time"

	dep "github.com/wayn3h0/gop/cache/dependency"
)

// dependency represents a file dependency.
type dependency struct {
	Path             string
	LastModifiedTime time.Time
}

// HasChanged reports whether the file has changed.
func (d *dependency) HasChanged() bool {
	fi, err := os.Stat(d.Path)
	if err != nil {
		return true
	}

	return fi.ModTime().After(d.LastModifiedTime)
}

// NewDependency returns a new file dependency.
func NewDependency(path string) dep.Dependency {
	var lastModifiedTime time.Time
	fi, err := os.Stat(path)
	if err != nil {
		lastModifiedTime = time.Now()
	} else {
		lastModifiedTime = fi.ModTime()
	}

	return &dependency{
		Path:             path,
		LastModifiedTime: lastModifiedTime,
	}
}
