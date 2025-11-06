package main

import (
	"crud_go/storage"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("crud app")

	storage.InitDBConnetion()

	tempPerson := storage.GetUserByIdFromDB(1)

	fmt.Println(tempPerson.ID)
	fmt.Println(tempPerson.Name)
	fmt.Println(tempPerson.LastName)

	CrudGoController()

	fmt.Println("Servers starts: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
