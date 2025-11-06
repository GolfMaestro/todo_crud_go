package main

import (
	"crud_go/service"
	"fmt"
	"net/http"
)

func CrudGoController() {
	http.HandleFunc("/hello", hello_handler)

	http.HandleFunc("/persons/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			service.GetPersonById(w, r)
		} else if r.Method == http.MethodDelete {
			service.DeletePersonById(w, r)
		}
		// else {
		// 	service.UpdatePersonNameById(w, r)
		// }
	})

	http.HandleFunc("/persons", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			service.CreatePerson(w, r)
		} else {
			service.GetPersons(w, r)
		}
	})
}

func hello_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello go")
}
