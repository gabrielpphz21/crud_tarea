package main

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Task struct {
	TaskId          string `json:"task_id"`
	TaskName        string `json:"task_name"`
	TaskDate        string `json:"task_date"`
	TaskDescription string `json:"task_description"`
}

func RegisterTask(task Task, sql_connection *pgx.Conn) error {
	query := "INSERT INTO tasks (task_name, task_date, task_description) VALUES ($1,$2, $3)"
	_, err := sql_connection.Exec(context.Background(), query, task.TaskName, task.TaskDate, task.TaskDescription)

	return err

}

func GetTasks(sql_connection *pgx.Conn) ([]Task, error) {
	query := "SELECT * FROM tasks"
	rows, err := sql_connection.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var response []Task

	for rows.Next() {
		var task Task
		var time_t time.Time

		err := rows.Scan(
			&task.TaskId,
			&task.TaskName,
			&time_t,
			&task.TaskDescription,
		)

		if err != nil {
			return nil, err
		}
		task.TaskDate = time_t.Format("2006-01-02")

		response = append(response, task)

	}

	return response, nil

}

func UpdateTask(selected_T Task, sql_connection *pgx.Conn) error {
	query := `UPDATE tasks
			  SET task_name=$1, task_date=$2, task_description=$3
			  WHERE task_id=$4`
	_, err := sql_connection.Exec(context.Background(), query, selected_T.TaskName, selected_T.TaskDate, selected_T.TaskDescription, selected_T.TaskId)

	return err

}

func DeleteTask(task_id int, sql_connection *pgx.Conn) error {
	query := `DELETE FROM tasks
			  WHERE task_id=$1`
	_, err := sql_connection.Exec(context.Background(), query, task_id)

	return err

}

func CreateDB(sql_connection *pgx.Conn) error {
	query := `
			DROP TABLE IF EXISTS tasks;
			CREATE TABLE tasks(
			task_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			task_name VARCHAR(200),
			task_date DATE,
			task_description TEXT)`

	_, err := sql_connection.Exec(context.Background(), query)

	return err
}
