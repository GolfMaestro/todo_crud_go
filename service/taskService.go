package service

import (
	"crud_go/models"
	"crud_go/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func GetTasksByPersonId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	parts := strings.Split(path, "/")

	requested_id, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Println("Something went wrong")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.GetTasksByPersonIdFromDB(requested_id))

}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Wrong JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json") // maybe just pass person
	json.NewEncoder(w).Encode(storage.InsertNewTaskInDB(newTask))
}

func UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	parts := strings.Split(path, "/")

	requested_id, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Println("Something went wrong")
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(storage.UpdateTaskStatusDB(requested_id))

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	parts := strings.Split(path, "/")

	requested_id, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Println("Something went wrong")
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(storage.DeleteTaskFromDB(requested_id))

}
