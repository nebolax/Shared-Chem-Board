package all_boards

import "sync"

type Board struct {
	ID        int
	Admin     int
	Name      string
	Password  string
	Observers []int
	History   [][]Point
}

type DataElem struct {
	board *Board
	mu    sync.Mutex
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}
