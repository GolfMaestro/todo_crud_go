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

func TestCreateTask(t *testing.T) {

	TestConnection(t)

	storage.Pool.Exec(context.Background(), "TRUNCATE TABLE tasks CASCADE")
	storage.Pool.Exec(context.Background(), "TRUNCATE TABLE tasks CASCADE")
	storage.Pool.Exec(context.Background(), "INSERT INTO persons (id, name, lastName) VALUES (1, 'Mike', 'Black')")

	values := strings.NewReader("{\"personId\": 1, \"title\": \"Learn GO\"}")
	req := httptest.NewRequest(http.MethodPost, "/tasks", values)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	service.CreateTask(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("waiting 201, get:  %d", w.Code)
	}

	var task models.Task
	if err := json.NewDecoder(w.Body).Decode(&task); err != nil {
		t.Fatalf("wrong JSON: %s", err)
	}

	if task.Title != "Learn GO" {
		t.Fatalf("Waiting Learn Go, get: %s", task.Title)
	}

	var count int
	storage.Pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM tasks WHERE person_id = $1", "1").Scan(&count)
	if count != 1 {
		t.Fatal("Task not saved in db")
	}

}

func TestUpdateTask(t *testing.T) {

	TestConnection(t)

	storage.Pool.Exec(context.Background(), "TRUNCATE TABLE tasks CASCADE")
	storage.Pool.Exec(context.Background(), "TRUNCATE TABLE tasks CASCADE")
	storage.Pool.Exec(context.Background(), "INSERT INTO persons (id, name, lastName) VALUES (1, 'Mike', 'Black')")
	storage.Pool.Exec(context.Background(), "INSERT INTO tasks (id, person_id, title) VALUES (1, 1, 'Learn Go')")

	req := httptest.NewRequest(http.MethodPut, "/tasks/1", nil)
	w := httptest.NewRecorder()

	service.UpdateTaskStatus(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("waiting 200, get:  %d", w.Code)
	}

	var status bool
	storage.Pool.QueryRow(context.Background(),
		"SELECT is_complete FROM tasks WHERE id = $1", 1,
	).Scan(&status)

	if !status {
		t.Fatal("Task not updated")
	}

}
