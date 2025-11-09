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
)

func TestReadTasks(t *testing.T) {

	TestConnection(t)

	storage.Pool.Exec(context.Background(), "TRUNCATE TABLE tasks CASCADE")
	storage.Pool.Exec(context.Background(), "TRUNCATE TABLE tasks CASCADE")
	storage.Pool.Exec(context.Background(), "INSERT INTO persons (id, name, lastName) VALUES (1, 'Mike', 'Black')")
	storage.Pool.Exec(context.Background(), "INSERT INTO tasks (person_id, title) VALUES (1, 'Learn Go')")

	req := httptest.NewRequest(http.MethodGet, "/tasks/1", nil)
	w := httptest.NewRecorder()

	service.GetTasksByPersonId(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("waiting 200, get:  %d", w.Code)
	}

	var tasks []models.Task

	if err := json.NewDecoder(w.Body).Decode(&tasks); err != nil {
		t.Fatal("Problem with decode json:", err)
	}

	isExist := false
	for i := range tasks {
		if tasks[i].Title == "Learn Go" {
			isExist = true
		}
	}

	if !isExist {
		t.Fatalf("Task \"Learn Go\" not in DB")
	}

}
