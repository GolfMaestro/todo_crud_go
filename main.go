package main

import (
	"crud_go/storage"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("crud app")

	storage.InitDBConnetion()

	CrudGoController()

	fmt.Println("Servers starts: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
