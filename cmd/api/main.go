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

func (app *application) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input todos.CreateTodoRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if input.Title == "" {
		http.Error(w, "title cannot be empty", http.StatusBadRequest)
		return
	}

	todo, err := app.todoService.Create(input.Title, input.Completed)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(todo)
}

func main() {
	app := &application{
		todoService: todos.NewService(),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			app.listTodoHandler(w, r)
			return
		}

		if r.Method == http.MethodPost {
			app.createTodoHandler(w, r)
			return
		}

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	log.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
