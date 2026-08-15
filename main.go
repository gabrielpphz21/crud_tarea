package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	conn, err := conn()
	if err != nil {
		log.Fatal("Base de datos caída")
	}
	http.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		Task_general(conn, w, r)
	})
	http.HandleFunc("tasks", func(w http.ResponseWriter, r *http.Request) {
		GetTaskHandler(conn, w, r)
	})

	log.Fatal(app.Listen(":3000"))
}
