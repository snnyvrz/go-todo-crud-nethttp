package todos

import "sync"

type Service struct {
	mu     sync.Mutex
	todos  map[int]Todo
	nextID int
}

func NewService() *Service {
	return &Service{
		todos:  make(map[int]Todo),
		nextID: 1,
	}
}
