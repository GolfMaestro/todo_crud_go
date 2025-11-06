package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Person struct {
	ID       int
	Name     string
	LastName string
}

var persons = []Person{
	{ID: 1, Name: "Ivan", LastName: "Ivanov"},
	{ID: 2, Name: "Maria", LastName: "Petrova"},
}

func main() {

	fmt.Println("crud app")

	http.HandleFunc("/hello", hello_handler)

	http.HandleFunc("/persons/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			getPersonById(w, r)
		} else if r.Method == http.MethodDelete {
			deletePersonById(w, r)
		} else {
			updatePersonNameById(w, r)
		}
	})

	http.HandleFunc("/persons", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			createPerson(w, r)
		} else {
			getPersons(w, r)
		}
	})

	fmt.Println("Servers starts: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func hello_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello go")
}

func getPersons(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func getPersonById(w http.ResponseWriter, r *http.Request) {

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

	ind := 0

	for i := 0; i < len(persons); i++ {
		if persons[i].ID == requested_id {
			ind = i
		}
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(persons[ind])
}

func createPerson(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newPerson Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		http.Error(w, "Wrong JSON", http.StatusBadRequest)
		return
	}

	maxID := 0
	for _, p := range persons {
		if p.ID > maxID {
			maxID = p.ID
		}
	}
	newPerson.ID = maxID + 1

	persons = append(persons, newPerson)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPerson)
}

func deletePersonById(w http.ResponseWriter, r *http.Request) {

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

	for i := 0; i < len(persons); i++ {
		if persons[i].ID == requested_id {
			ind = i
		}
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(persons[ind])

	persons = append(persons[:ind], persons[ind+1:]...)
}

func updatePersonNameById(w http.ResponseWriter, r *http.Request) {

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

	for i := 0; i < len(persons); i++ {
		if persons[i].ID == requested_id {
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

	persons[ind].Name = *updates.Name

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(persons[ind])

}
