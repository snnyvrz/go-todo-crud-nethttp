package todos

import "sync"

type Service struct {
	mu     sync.Mutex
	todos  map[int]Todo
	nextID int
}

func (s *Service) List() ([]Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	items := make([]Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		items = append(items, todo)
	}
	return items, nil
}

func NewService() *Service {
	return &Service{
		todos:  make(map[int]Todo),
		nextID: 1,
	}
}
