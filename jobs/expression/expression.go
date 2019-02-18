package expression

import (
	"time"
)

// Expression represents an expression specifies a duty cycle.
type Expression interface {
	// Next returns the closest activated time from given time.
	Next(time.Time) time.Time
}
