package tests

import (
	"context"
	"crud_go/storage"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

// run tests: go test ./... -v

func TestConnection(t *testing.T) {

	ctx := context.Background()

	config, err := pgxpool.ParseConfig("postgres://postgres:2004@localhost:5432/test_crud_go")
	if err != nil {
		t.Fatal(err)
	}
	config.MaxConns = 5
	config.MinConns = 1

	storage.Pool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		t.Fatal(err)
	}
}
