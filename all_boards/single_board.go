package all_boards

import "sync"

type Board struct {
	ID    int
	Name  string
	Users []int
	mu    sync.Mutex
}
