package all_boards

import (
	"sync"
)

type DrawingsHistory [][]Point

type Board struct {
	ID              int
	Admin           int
	Name            string
	Password        string
	Observers       []*Observer
	DrawingsHistory DrawingsHistory
	ChatHistory     ChatHistory
}

type Observer struct {
	UserID          int
	DrawingsHistory DrawingsHistory
	ChatHistory     ChatHistory
}

type DataElem struct {
	board *Board
	mu    sync.Mutex
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}
