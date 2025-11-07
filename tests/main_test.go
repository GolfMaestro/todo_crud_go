package tests

import (
	"context"
	"crud_go/models"
	"crud_go/service"
	"crud_go/storage"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

func TestReadPersons(t *testing.T) {

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

	storage.Pool.Exec(context.Background(), "TRUNCATE TABLE persons CASCADE")
	storage.Pool.Exec(context.Background(), "INSERT INTO persons (name, lastName) VALUES ('Mike', 'Black')")

	req := httptest.NewRequest("GET", "/persons", nil)
	w := httptest.NewRecorder()

	service.GetPersons(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("waiting 200, get:  %d", w.Code)
	}

	var persons []models.Person

	if err := json.NewDecoder(w.Body).Decode(&persons); err != nil {
		t.Fatal("Problem with decode json:", err)
	}

	isExist := false
	for i := range persons {
		if persons[i].Name == "Mike" {
			isExist = true
		}
	}

	if !isExist {
		t.Fatalf("Mike dont exists in DB")
	}

}
