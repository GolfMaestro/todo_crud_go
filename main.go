package main

import (
	"context"
	"crud_go/storage"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("crud app")

	storage.InitDBConnetion()

	var id, name, lastName string
	err := storage.Pool.QueryRow(context.Background(),
		"SELECT id, name, lastName FROM persons WHERE id = $1", 1,
	).Scan(&id, &name, &lastName)
	if err != nil {
		fmt.Println("Sonmething went wrong")
	}

	fmt.Println(id)
	fmt.Println(name)
	fmt.Println(lastName)

	CrudGoController()

	fmt.Println("Servers starts: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
