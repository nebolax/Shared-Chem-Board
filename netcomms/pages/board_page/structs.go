package board_page

import (
	"ChemBoard/all_boards"
	"sync"

	"github.com/gorilla/websocket"
)

//TODO lock structs while messaging

type sockClient interface {
	isAdmin() bool
	sock() *websocket.Conn
	boardID() int
	userID() int
	mu() *sync.Mutex
}

type adminClient struct {
	dboardID int
	duserID  int
	dview    int
	dsock    *websocket.Conn
	dmu      *sync.Mutex
}

type observerClient struct {
	dboardID int
	duserID  int
	dview    bool // false - general board, true - personal board
	dsock    *websocket.Conn
	dmu      *sync.Mutex
}

func (cl adminClient) sock() *websocket.Conn { return cl.dsock }
func (cl adminClient) boardID() int          { return cl.dboardID }
func (cl adminClient) userID() int           { return cl.duserID }
func (cl adminClient) mu() *sync.Mutex       { return cl.dmu }
func (cl adminClient) isAdmin() bool         { return true }

func (cl observerClient) sock() *websocket.Conn { return cl.dsock }
func (cl observerClient) boardID() int          { return cl.dboardID }
func (cl observerClient) userID() int           { return cl.duserID }
func (cl observerClient) mu() *sync.Mutex       { return cl.dmu }
func (cl observerClient) isAdmin() bool         { return false }

type msgType int

const (
	tPoints     msgType = iota
	tObsStat    msgType = iota
	tChview     msgType = iota
	tInpChatMsg msgType = iota
	tOutChatMsg msgType = iota
)

type anyMSG struct {
	Type msgType     `json:"type"`
	Data interface{} `json:"data"`
}
type pointsMSG struct {
	Points []all_boards.Point `json:"points"`
}

type chviewMSG struct {
	NView int `json:"nview"`
}

type allObsStatMSG struct {
	Info []singleObsInfo `json:"allObsInfo"`
}
type singleObsInfo struct {
	UserID   int    `json:"userid"`
	UserName string `json:"username"`
}
