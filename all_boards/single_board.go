package all_boards

import "sync"

type Board struct {
	ID int
	mu sync.Mutex
}
