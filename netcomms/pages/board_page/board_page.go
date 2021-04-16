package board_page

import (
	"ChemBoard/all_boards"
	"ChemBoard/netcomms/connsinc"
	"ChemBoard/netcomms/session_info"
	"html/template"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func HandleSockets(w http.ResponseWriter, r *http.Request) {
	println("new user")
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", http.StatusForbidden)
	} else {
		ws, _ := websocket.Upgrade(w, r, nil, 0, 0)
		vars := mux.Vars(r)
		boardID, _ := strconv.Atoi(vars["id"])
		if board := all_boards.GetByID(boardID); board != nil {
			RegNewBoardObserver(ws, board, session_info.GetSessionUserID(r))
		}
	}
}

func BoardPage(w http.ResponseWriter, r *http.Request) {
	if !session_info.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	}
	tmpl, _ := template.ParseFiles("./templates/board.html")
	tmpl.Execute(w, nil)
}

type sockClient struct {
	userID int
	board  *all_boards.Board
	sock   *websocket.Conn
	mu     sync.Mutex
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

//canvasMessage is a struct
type canvasMessage struct {
	Points []Point `json:"points"`
}

type boardPageSetup struct {
	Type         string    `json:"type"`
	History      []float64 `json:"history"`
	PublicOffers int       `json:"publicOffers"`
	PersonOffers int       `json:"personOffers"`
}

var clients = make(map[int]*sockClient)

//SendtoBoardObservers is func
func SendtoBoardObservers(boardID int, message interface{}) {
	for _, client := range clients {
		if client.board.ID == boardID {
			sendtoUserDevices(client.userID, message)
		}
	}
}

//sendtoUserDevices is func
func sendtoUserDevices(userID int, message interface{}) {
	for connID, client := range clients {
		if client.userID == userID {
			writeSingleMessage(connID, message)
		}
	}
}

//writeSingleMessage is func
func writeSingleMessage(connID int, message interface{}) {
	client := clients[connID]
	println(message)
	println(client.userID)
	client.mu.Lock()
	err := client.sock.WriteJSON(message)
	if err != nil {
		delClient(connID)
	}
	client.mu.Unlock()
}

func delClient(connID int) {
	clients[connID].sock.Close()
	delete(clients, connID)
}

//ReadSingleMessage is func
func readSingleMessage(connID int) (canvasMessage, bool) {
	var msg canvasMessage
	err := clients[connID].sock.ReadJSON(&msg)
	println("read")
	if err != nil {
		println(err.Error())
		delClient(connID)
		return canvasMessage{}, false
	}
	return msg, true
}

func procIncomingMessages(connID int) {
	client := clients[connID]
	cls := clients
	_ = cls
	for {
		msg, ok := readSingleMessage(connID)
		if ok {
			SendtoBoardObservers(client.board.ID, msg)
			println("sent")
		} else {
			break
		}
	}
}

//RegNewBoardObserver is func
func RegNewBoardObserver(ws *websocket.Conn, board *all_boards.Board, userID int) {
	connID := connsinc.NewID()
	clients[connID] = &sockClient{userID: userID, board: board, sock: ws}
	//writeSingleMessage(connID)
	go procIncomingMessages(connID)
}
