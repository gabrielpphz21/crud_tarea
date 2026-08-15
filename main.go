package main

import (
	"log"
	"net/http"
)

func main() {

	conn, err := conn()
	if err != nil {
		log.Fatal("Base de datos caída")
	}

	err1 := CreateDB(conn)
	if err1 != nil {
		log.Fatal("No se pudo crear la base de datos ( tablas )")
	}

	http.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		Task_general(conn, w, r)
	})
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		GetTaskHandler(conn, w, r)
	})

	err2 := http.ListenAndServe(":8080", nil)
	if err2 != nil {
		log.Fatal(err2)
	}

}
