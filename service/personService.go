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

func GetPersons(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.GetUsersFromDB())
}

func GetPersonById(w http.ResponseWriter, r *http.Request) {

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

	json.NewEncoder(w).Encode(storage.GetUserByIdFromDB(requested_id))
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newPerson models.Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		http.Error(w, "Wrong JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json") // maybe just pass person
	json.NewEncoder(w).Encode(storage.InsertNewPersonInDB(newPerson.Name, newPerson.LastName))
}

func DeletePersonById(w http.ResponseWriter, r *http.Request) {

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

	json.NewEncoder(w).Encode(storage.DeleteUserByIDFromDB(requested_id))

}

func UpdatePersonNameById(w http.ResponseWriter, r *http.Request) {

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

	var updates struct {
		Name     *string `json:"name,omitempty"`
		LastName *string `json:"lastName,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(storage.UpdatePersonNameById(requested_id, *updates.Name))

}
