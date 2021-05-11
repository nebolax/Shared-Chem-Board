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
	boardID() uint64
	userID() uint64
	mu() *sync.Mutex
}

type adminClient struct {
	dboardID uint64
	duserID  uint64
	dview    uint64
	dsock    *websocket.Conn
	dmu      *sync.Mutex
}

type observerClient struct {
	dboardID uint64
	duserID  uint64
	dview    bool // false - general board, true - personal board
	dsock    *websocket.Conn
	dmu      *sync.Mutex
}

func (cl adminClient) sock() *websocket.Conn { return cl.dsock }
func (cl adminClient) boardID() uint64       { return cl.dboardID }
func (cl adminClient) userID() uint64        { return cl.duserID }
func (cl adminClient) mu() *sync.Mutex       { return cl.dmu }
func (cl adminClient) isAdmin() bool         { return true }

func (cl observerClient) sock() *websocket.Conn { return cl.dsock }
func (cl observerClient) boardID() uint64       { return cl.dboardID }
func (cl observerClient) userID() uint64        { return cl.duserID }
func (cl observerClient) mu() *sync.Mutex       { return cl.dmu }
func (cl observerClient) isAdmin() bool         { return false }

type msgType uint64

const (
	tAction     msgType = iota
	tSetId      msgType = iota
	tObsStat    msgType = iota
	tChview     msgType = iota
	tInpChatMsg msgType = iota
	tOutChatMsg msgType = iota
)

type anyMSG struct {
	Type msgType     `json:"type"`
	Data interface{} `json:"data"`
}

type ChatMessage struct {
}

type chviewMSG struct {
	NView uint64 `json:"nview"`
}
type SetIdMSG struct {
	Property string `json:"property"`
	ID       uint64 `json:"id"`
}
type allObsStatMSG struct {
	Info []singleObsInfo `json:"allObsInfo"`
}
type singleObsInfo struct {
	UserID   uint64 `json:"userid"`
	UserName string `json:"username"`
}

type chatMessage struct {
	ID         uint64                 `json:"id"`
	SenderInfo interface{}            `json:"senderinfo"`
	TimeStamp  all_boards.TimeStamp   `json:"timestamp"`
	Content    all_boards.ChatContent `json:"content"`
}
