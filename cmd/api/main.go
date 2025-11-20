package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/snnyvrz/go-todo-crud-nethttp/internal/todos"
)

type application struct {
	todoService *todos.Service
}

func (app *application) listTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	todos, err := app.todoService.List()

	if err != nil {
		http.Error(w, "failed to list todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	app := &application{
		todoService: todos.NewService(),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/todos", app.listTodoHandler)

	log.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
