package all_boards

import (
	"sync"
)

type drawingType int

const (
	tFreeMouse drawingType = iota
)

type DrawingMSG struct {
	Type drawingType `json:"type"`
	Data interface{} `json:"data"`
}

type DrawingsHistory []DrawingMSG

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
