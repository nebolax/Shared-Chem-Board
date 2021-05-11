package all_boards

import (
	"sync"
)

type Drawing struct {
	ID   uint64      `json:"id"`
	Type uint64      `json:"type"`
	Data interface{} `json:"data"`
}

type ActionMSG struct {
	ID      uint64  `json:"id"`
	Type    uint64  `json:"type"`
	Drawing Drawing `json:"drawing"`
}

type ActionsHistory []ActionMSG

type Board struct {
	ID        uint64
	Admin     uint64
	Name      string
	Password  string
	Observers []*Observer
	Actions   ActionsHistory
	Chat      ChatHistory
}

type Observer struct {
	DBID    uint64
	UserID  uint64
	Actions ActionsHistory
	Chat    ChatHistory
}

type DataElem struct {
	board Board
	mu    sync.Mutex
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}
