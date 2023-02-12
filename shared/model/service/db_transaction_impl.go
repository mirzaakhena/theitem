package service

import (
	"context"
	"theitem/shared/model/repository"
)

// WithTransaction is helper function that simplify the transaction execution handling
func WithTransaction[T any](ctx context.Context, trx repository.WithTransactionDB, trxFunc func(dbCtx context.Context) (*T, error)) (*T, error) {
	dbCtx, err := trx.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			err = trx.RollbackTransaction(dbCtx)
			panic(p)

		} else if err != nil {
			err = trx.RollbackTransaction(dbCtx)

		} else {
			err = trx.CommitTransaction(dbCtx)

		}
	}()

	var t *T
	t, err = trxFunc(dbCtx)
	if err != nil {
		return nil, err
	}

	return t, err
}
