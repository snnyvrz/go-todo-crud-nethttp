package todos

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type CreateTodoRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
