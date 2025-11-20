package todos

import (
	"errors"
	"sync"
)

type Service struct {
	mu     sync.Mutex
	todos  map[int]Todo
	nextID int
}

var ErrNotFound = errors.New("todo not found")

func (s *Service) List() ([]Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	items := make([]Todo, 0, len(s.todos))

	for _, todo := range s.todos {
		items = append(items, todo)
	}

	return items, nil
}

func (s *Service) Get(id int) (Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, ok := s.todos[id]

	if !ok {
		return Todo{}, ErrNotFound
	}

	return todo, nil
}

func (s *Service) Create(title string, completed bool) (Todo, error) {
	if title == "" {
		return Todo{}, errors.New("title cannot be empty")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.nextID

	todo := Todo{
		ID:        id,
		Title:     title,
		Completed: completed,
	}

	s.todos[id] = todo

	s.nextID++

	return todo, nil
}

func (s *Service) Update(id int, title string, completed bool) (Todo, error) {
	if title == "" {
		return Todo{}, errors.New("title cannot be empty")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	todo, ok := s.todos[id]

	if !ok {
		return Todo{}, ErrNotFound
	}

	todo.Title = title
	todo.Completed = completed

	s.todos[id] = todo

	return todo, nil
}

func NewService() *Service {
	return &Service{
		todos:  make(map[int]Todo),
		nextID: 1,
	}
}
