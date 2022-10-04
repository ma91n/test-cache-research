package sample

import (
	"context"
	"testing"

	"github.com/future-architect/go-exceltesting"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

func TestResetStatus(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "サンプル",
			wantErr: false,
		},
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, "postgres://sample:password@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()

	db := stdlib.OpenDB(*pool.Config().ConnConfig)
	defer db.Close()

	e := exceltesting.New(db)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e.Load(t, exceltesting.LoadRequest{
				TargetBookPath: "../testdata/input.xlsx",
			})

			if err := ResetStatus(ctx, pool); (err != nil) != tt.wantErr {
				t.Errorf("ResetStatus() error = %v, wantErr %v", err, tt.wantErr)
			}

			e.Compare(t, exceltesting.CompareRequest{
				TargetBookPath: "../testdata/want.xlsx",
				IgnoreColumns:  []string{"created_at", "updated_at"},
			})
		})
	}
}
