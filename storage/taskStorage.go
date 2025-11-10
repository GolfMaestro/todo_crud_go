package storage

import (
	"context"
	"crud_go/models"
	"fmt"
	"time"
)

func GetTasksByPersonIdFromDB(req_person_id int) []models.Task {

	rows, err := Pool.Query(context.Background(),
		"SELECT id, person_id, title, is_complete, created_at FROM tasks WHERE person_id = $1",
		req_person_id)

	if err != nil {
		fmt.Println("Sonmething went wrong in fuction GetTasksByPersonIdFromDB")
	}

	var tasks []models.Task

	defer rows.Close()

	for rows.Next() {
		var t models.Task
		temp_err := rows.Scan(&t.ID, &t.PersonID, &t.Title, &t.IsComplete, &t.CreatedAt)
		if temp_err != nil {
			fmt.Println("Something went wrong")
		}
		tasks = append(tasks, t)
	}

	return tasks

}

func InsertNewTaskInDB(newTask models.Task) models.Task {

	var taskID int
	err := Pool.QueryRow(context.Background(),
		"INSERT INTO tasks (person_id, title) VALUES ($1, $2) RETURNING id",
		newTask.PersonID, newTask.Title).Scan(&taskID)

	if err != nil {
		fmt.Println("Something went wrong in funciton InsertNewPersonInDB")
	}

	temp_task := models.Task{
		ID:         taskID,
		PersonID:   newTask.ID,
		Title:      newTask.Title,
		IsComplete: false,
		CreatedAt:  time.Now(),
	}

	return temp_task

}

func UpdateTaskStatusDB(req_id int) int {

	_, err := Pool.Exec(context.Background(),
		"UPDATE tasks SET is_complete = true WHERE id = $1", req_id)
	if err != nil {
		fmt.Println("Something went wrong in funciton UpdateTaskStatusDB")
	}

	return req_id

}
func DeleteTaskFromDB(req_id int) int {

	_, err := Pool.Exec(context.Background(),
		"DELETE FROM tasks WHERE id = $1", req_id)

	if err != nil {
		fmt.Println("Something went wrong in funciton DeleteTaskFromDB")
	}

	return req_id

}
