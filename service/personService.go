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
	json.NewEncoder(w).Encode(models.Persons)
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

	// ind := 0

	// for i := 0; i < len(models.Persons); i++ {
	// 	if models.Persons[i].ID == requested_id {
	// 		ind = i
	// 	}
	// }

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

	maxID := 0
	for _, p := range models.Persons {
		if p.ID > maxID {
			maxID = p.ID
		}
	}
	newPerson.ID = maxID + 1

	models.Persons = append(models.Persons, newPerson)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPerson)
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

	ind := 0

	for i := 0; i < len(models.Persons); i++ {
		if models.Persons[i].ID == requested_id {
			ind = i
		}
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(models.Persons[ind])

	models.Persons = append(models.Persons[:ind], models.Persons[ind+1:]...)
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

	ind := 0

	for i := 0; i < len(models.Persons); i++ {
		if models.Persons[i].ID == requested_id {
			ind = i
		}
	}

	var updates struct {
		Name     *string `json:"name,omitempty"`
		LastName *string `json:"lastName,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	models.Persons[ind].Name = *updates.Name

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(models.Persons[ind])

}
