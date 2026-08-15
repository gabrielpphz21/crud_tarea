package main

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Task struct {
	id               string
	task_name        string
	task_date        string
	task_description string
}

func RegisterTask(task Task, sql_connection *pgx.Conn) error {
	query := "INSERT INTO tasks (task_name, task_date, task_description) VALUES ($1,$2, $3, $4)"
	_, err := sql_connection.Exec(context.Background(), query, task.task_name, task.task_date, task.task_description)

	return err

}

func GetTasks(sql_connection *pgx.Conn) ([]Task, error) {
	query := "SELECT * FROM tasks"
	rows, err := sql_connection.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	var response []Task

	for rows.Next() {
		var task Task

		err := rows.Scan(
			&task.id,
			&task.task_name,
			&task.task_date,
			&task.task_description,
		)
		if err != nil {
			return nil, err
		}

		response = append(response, task)

	}

	return response, nil

}

func UpdateTask(selected_T Task, sql_connection *pgx.Conn) error {
	query := `UPDATE tasks
			  SET task_name=$1, task_date=$2, task_description=$3
			  WHERE id=$4`
	_, err := sql_connection.Query(context.Background(), query)

	return err

}

func DeleteTask(task_id string, sql_connection *pgx.Conn) error {
	query := `DELETE FROM tasks
			  WHERE id=$1`
	_, err := sql_connection.Exec(context.Background(), query, task_id)

	return err

}

func CreateDB(sql_connection *pgx.Conn) error {
	query := `
			CREATE TABLE tasks(
			task_id INT PRIMARY KEY AUTOINCREMENT,
			task_name VARCHAR(200),
			task_date DATE,
			task_description TEXT)`

	_, err := sql_connection.Exec(context.Background(), query)

	return err
}
