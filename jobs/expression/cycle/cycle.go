package cycle

import (
	"time"

	expr "github.com/wayn3h0/gop/jobs/expression"
)

// expression represents a cyclic expression.
type expression struct {
	Interval time.Duration
}

// Next implements jobs.Expression interface.
func (e *expression) Next(from time.Time) time.Time {
	return from.Add(e.Interval)
}

// NewExpression returns a new cycle expression.
// A duration string supported, check time.ParseDuration.
// Less than a second are not supported (will round up to 1 second).
func NewExpression(dur time.Duration) expr.Expression {
	if dur < time.Second {
		dur = time.Second
	}

	return &expression{
		Interval: dur,
	}
}
