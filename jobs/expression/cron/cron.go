package cron

import (
	"github.com/wayn3h0/gop/errors"
	expr "github.com/wayn3h0/gop/jobs/expression"

	"github.com/gorhill/cronexpr"
)

// Expression represents a cron expression.
type expression struct {
	*cronexpr.Expression
}

func predefined(expr string) expr.Expression {
	return &expression{
		Expression: cronexpr.MustParse(expr),
	}
}

// Predefined cron expressions.
var (
	AnnuallyExpression = predefined("@annually") // every year at midnight in the morning of January 1
	YearlyExpression   = predefined("@yearly")   // every year at midnight in the morning of January 1
	MonthlyExpression  = predefined("@monthly")  // every month at midnight in the morning of the first of the month
	WeeklyExpression   = predefined("@weekly")   // every week at midnight in the morning of Sunday
	DailyExpression    = predefined("@daily")    // every day at midnight
	HourlyExpression   = predefined("@hourly")   // every hour at the beginning of the hour
)

// NewExpression returns a new cron expression.
// We use github.com/gorhill/cronexpr for parsing cron express inside.
func NewExpression(expr string) (expr.Expression, error) {
	cronexpr, err := cronexpr.Parse(expr)
	if err != nil {
		return nil, errors.Wrapf(err, "cron: could not parse cron expression %q", expr)
	}

	return &expression{
		Expression: cronexpr,
	}, nil
}
