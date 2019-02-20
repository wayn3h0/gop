package jobs

import (
	"sort"
	"time"

	"github.com/wayn3h0/gop/jobs/expression"
)

// job represents a background job.
type job struct {
	Function   func()
	Expression expression.Expression
	Name       string
	Previous   time.Time
	Next       time.Time
}

// jobs represents a sortable collection of job.
type jobs []*job

// Len implements sort.Interface.
func (j jobs) Len() int {
	return len(j)
}

// Swap implements sort.Interface.
func (j jobs) Swap(x, y int) {
	j[x], j[y] = j[y], j[x]
}

// Less implements sort.Interface.
func (j jobs) Less(x, y int) bool {
	if j[x].Next.IsZero() {
		return false
	}

	if j[y].Next.IsZero() {
		return true
	}

	return j[x].Next.Before(j[y].Next)
}

// Sort sorts the jobs.
func (j jobs) Sort() {
	sort.Sort(j)
}
