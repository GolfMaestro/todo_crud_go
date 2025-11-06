package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("crud app")

	CrudGoController()

	fmt.Println("Servers starts: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
