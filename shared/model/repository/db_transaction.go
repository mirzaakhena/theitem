package repository

import "context"

// WithTransactionDB used for common transaction handling
// all the context must use the same database session.
type WithTransactionDB interface {
	BeginTransaction(ctx context.Context) (context.Context, error)
	CommitTransaction(ctx context.Context) error
	RollbackTransaction(ctx context.Context) error
}
