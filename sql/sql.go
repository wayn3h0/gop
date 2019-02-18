package sql

// Rows represents rows returned from SQL database.
type Rows interface {
	// Next moves the point to next row.
	Next() bool

	// Scan parses the data from current row.
	Scan(dest ...interface{}) error
}

// Transaction represents a SQL transaction.
type Transaction interface {
	// Execute executes the command and returns the number of rows affected in the transaction.
	Execute(statement string, args ...interface{}) (int64, error)

	// Query returns the rows selected in the transaction.
	Query(statement string, args ...interface{}) (Rows, error)

	// Rollback aborts the transaction.
	Rollback() error

	// Commit commits the transaction.
	Commit() error
}

// Database represents a SQL database.
type Database interface {
	// Execute executes the command and returns the number of rows affected.
	Execute(statement string, args ...interface{}) (int64, error)

	// Query returns the rows selected.
	Query(statement string, args ...interface{}) (Rows, error)

	// Begin starts a transaction.
	Begin() (Transaction, error)
}
