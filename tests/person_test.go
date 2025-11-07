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
)

func TestReadPersons(t *testing.T) {

	TestConnection(t)

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

func TestReadPersonById(t *testing.T) {

	TestConnection(t)

	storage.Pool.Exec(context.Background(), "TRUNCATE TABLE persons CASCADE")
	storage.Pool.Exec(context.Background(), "INSERT INTO persons (id, name, lastName) VALUES (1, 'Mike', 'Black')")

	req := httptest.NewRequest(http.MethodGet, "/persons/1", nil)
	w := httptest.NewRecorder()

	service.GetPersonById(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("waiting 200, get:  %d", w.Code)
	}

	var person models.Person

	if err := json.NewDecoder(w.Body).Decode(&person); err != nil {
		t.Fatal("Problem with decode json:", err)
	}

	if person.Name != "Mike" {
		t.Fatalf("Mike dont exists in DB")
	}

}

func TestCreatePerson(t *testing.T) {

	TestConnection(t)

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
	storage.Pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM persons WHERE name = $1", "Mike").Scan(&count)
	if count != 1 {
		t.Fatal("Person not saved in db")
	}

}

func TestUpdatePerson(t *testing.T) {

	TestConnection(t)

	storage.Pool.Exec(context.Background(), "TRUNCATE TABLE persons CASCADE")
	storage.Pool.Exec(context.Background(), "INSERT INTO persons (id, name, lastName) VALUES (1, 'Mike', 'Black')")

	values := strings.NewReader("{\"name\": \"Sandy\"}")
	req := httptest.NewRequest(http.MethodPut, "/persons/1", values)
	w := httptest.NewRecorder()

	service.UpdatePersonNameById(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("waiting 200, get:  %d", w.Code)
	}

	var id int
	var name, lastName string
	storage.Pool.QueryRow(context.Background(),
		"SELECT id, name, lastName FROM persons WHERE id = $1", 1,
	).Scan(&id, &name, &lastName)

	if name != "Sandy" {
		t.Fatal("Person not updated")
	}

}
