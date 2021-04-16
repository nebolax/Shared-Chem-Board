package all_boards

import "sync"

type Board struct {
	ID       int
	Admin    int
	Name     string
	Password string
	Users    []int
	mu       sync.Mutex
}
