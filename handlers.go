package main

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func Task_general(conn *pgx.Conn, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		RegisterTaskHandler(conn, w, r)
	}
	if r.Method == http.MethodDelete {
		DeleteTaskHandler(conn, w, r)
	}
	if r.Method == http.MethodPut {
		UpdateTaskHandler(conn, w, r)
	}

	http.Error(w, "Método no valido", http.StatusMethodNotAllowed)
}

func GetTaskHandler(conn *pgx.Conn, w http.ResponseWriter, r *http.Request) {
	tasks, err := GetTasks(conn)
	if err != nil {
		http.Error(w, "Error al traer las Tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func RegisterTaskHandler(conn *pgx.Conn, w http.ResponseWriter, r *http.Request) {
	var task Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "JSON de registro de Tasks inválido", http.StatusBadRequest)
		return
	}

	err1 := RegisterTask(task, conn)
	if err1 != nil {
		http.Error(w, "Error al registrar la Task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]bool{
		"success": true,
	}
	json.NewEncoder(w).Encode(resp)
}

func DeleteTaskHandler(conn *pgx.Conn, w http.ResponseWriter, r *http.Request) {
	var task_id string

	err := json.NewDecoder(r.Body).Decode(&task_id)
	if err != nil {
		http.Error(w, "JSON de Borrado de Task inválido", http.StatusBadRequest)
		return
	}

	err1 := DeleteTask(task_id, conn)
	if err1 != nil {
		http.Error(w, "Error al borrar la task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]bool{
		"success": true,
	}
	json.NewEncoder(w).Encode(resp)
}

func UpdateTaskHandler(conn *pgx.Conn, w http.ResponseWriter, r *http.Request) {
	var task Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "JSON de actualización de task inválido", http.StatusBadRequest)
		return
	}

	err1 := UpdateTask(task, conn)
	if err1 != nil {
		http.Error(w, "Error al actualizar la task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]bool{
		"success": true,
	}
	json.NewEncoder(w).Encode(resp)
}
