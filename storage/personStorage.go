package storage

import (
	"context"
	"crud_go/models"
	"fmt"
)

func GetUserByIdFromDB(req_id int) models.Person {

	var id, name, lastName string
	err := Pool.QueryRow(context.Background(),
		"SELECT id, name, lastName FROM persons WHERE id = $1", req_id,
	).Scan(&id, &name, &lastName)
	if err != nil {
		fmt.Println("Sonmething went wrong")
	}

	tempPerson := models.Person{
		ID:       req_id, // Лучше переконвертировать id в инт я думаю
		Name:     name,
		LastName: lastName,
	}

	return tempPerson

}
