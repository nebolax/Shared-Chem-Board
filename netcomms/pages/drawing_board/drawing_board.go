package drawing_board

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

type UInfo struct {
	ID     int
	Status string
}

func HandleSockets(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", http.StatusForbidden)
	} else {
		ws, _ := websocket.Upgrade(w, r, nil, 0, 0)
		vars := mux.Vars(r)
		boardID, _ := strconv.Atoi(vars["id"])
		if board := all_boards.GetByID(boardID); board != nil {
			RegNewBoardObserver(ws, board, session_info.GetUserID(r))
		}
	}
}

func Page(w http.ResponseWriter, r *http.Request) {
	if !session_info.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		vars := mux.Vars(r)
		boardID, _ := strconv.Atoi(vars["id"])
		if !all_boards.AvailableToUser(session_info.GetUserID(r), boardID) {
			http.Redirect(w, r, "/myboards", http.StatusSeeOther)
		} else {
			s := "observer"
			if all_boards.IsAdmin(session_info.GetUserID(r), boardID) {
				s = "admin"
			}
			tmpl, _ := template.ParseFiles("./templates/board.html")
			tmpl.Execute(w, UInfo{session_info.GetUserID(r), s})
		}
	}
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
	if err != nil {
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
