package main

import (
	"log"
	"net/http"

	"github.com/snnyvrz/go-todo-crud-nethttp/internal/todos"
)

func main() {
	todoService := todos.NewService()

	mux := http.NewServeMux()

	log.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
