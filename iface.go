package presentation

import (
	"context"
	"time"
)
// TransactionManager contains public methods of package dtpc
type TransactionManager interface {

	// StartTransaction starts a transaction process
	StartTransaction(
		ctx context.Context,
		req Request,
		callbacks ...func() error) (*Response, error)

	// GetTransactions retrieves a list of transactions based on given condition
	GetTransactions(
		ctx context.Context,
		state TransactionState,
		query string) ([]*Transaction, error)

	// RecoverTransations recovers unfinished or failed transactions
	RecoverTransactions(ctx context.Context, recoverTime time.Time) error
}
// Ends OMIT

// TransactionHandler defines necessary methods of handling transaction data
type TransactionHandler interface {

	Insert(
		ctx context.Context,
		source, destination, reference string,
		data interface{}) (string, error)

	UpdateState(
		ctx context.Context,
		id string,
		newState TransactionState) (*Transaction, error)

	GetTransaction(ctx context.Context, id string) (*Transaction, error)

	GetTransactionsInState(
		ctx context.Context,
		state TransactionState,
		query string) ([]*Transaction, error)
}
// Ends OMIT

// AccountHandler defines necessary methods of handling account data
type AccountHandler interface {

	Get(ctx context.Context, accountID string, retval Account) error

	Put(ctx context.Context, doc Account) error

	Update(ctx context.Context, accountID, transactionID string, tr Request) error

	Rollback(ctx context.Context, accountID, transactionID string, tr Request) error

	Commit(ctx context.Context, accountID, transactionID string) error
}
// Ends OMIT