package helper

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CommitOrRollback(err error, ctx context.Context, tx pgx.Tx) error {
	if err != nil {
		return tx.Rollback(ctx)
	}
	return tx.Commit(ctx)
}
