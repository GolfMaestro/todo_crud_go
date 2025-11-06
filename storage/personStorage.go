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
		fmt.Println("Sonmething went wrong in fuction GetUserByIdFromDB")
	}

	tempPerson := models.Person{
		ID:       req_id, // Лучше переконвертировать id в инт я думаю
		Name:     name,
		LastName: lastName,
	}

	return tempPerson

}

func GetUsersFromDB() []models.Person {

	rows, err := Pool.Query(context.Background(),
		"SELECT id, name, lastName FROM persons")

	if err != nil {
		fmt.Println("Something went wrong in function GetUsersFromDB")
	}

	defer rows.Close()

	var persons []models.Person

	for rows.Next() {
		var t models.Person
		temp_err := rows.Scan(&t.ID, t.Name, t.LastName)
		if temp_err != nil {
			fmt.Println("Something went wrong")
		}
		persons = append(persons, t)
	}

	return persons

}

func InsertNewPersonInDB(name string, lastName string) models.Person {

	var userID int
	err := Pool.QueryRow(context.Background(),
		"INSERT INTO persons (name, lastName) VALUES ($1, $2) RETURNING id", name, lastName).Scan(&userID)

	if err != nil {
		fmt.Println("Something went wrong")
	}

	tempPerson := models.Person{
		ID:       userID,
		Name:     name,
		LastName: lastName,
	}

	return tempPerson

}
