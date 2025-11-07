package tests

import (
	"context"
	"crud_go/models"
	"crud_go/service"
	"crud_go/storage"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

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

	req := httptest.NewRequest(http.MethodGet, "/persons", nil)
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

func TestCreatePerson(t *testing.T) {

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

	values := strings.NewReader("{\"name\": \"Mike\", \"lastName\": \"Black\"}")
	req := httptest.NewRequest(http.MethodPost, "/persons", values)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	service.CreatePerson(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("waiting 201, get:  %d", w.Code)
	}

	var person models.Person
	if err := json.NewDecoder(w.Body).Decode(&person); err != nil {
		t.Fatalf("wrong JSON: %s", err)
	}

	if person.Name != "Mike" {
		t.Fatalf("Waiting Mike, get: %s", person.Name)
	}

	var count int
	storage.Pool.QueryRow(ctx, "SELECT COUNT(*) FROM persons WHERE name = $1", "Mike").Scan(&count)
	if count != 1 {
		t.Fatal("Person not saved in db")
	}

}
