package sample

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ResetStatus(ctx context.Context, conn *pgxpool.Pool) error {
	_, err := conn.Exec(ctx, `
update company
set status     = 2,
    updated_at = current_timestamp,
    revision   = revision + 1
`)

	if err != nil {
		return fmt.Errorf("update exec: %w", err)
	}
	return nil
}
