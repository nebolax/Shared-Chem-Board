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

func HandleSockets(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", http.StatusForbidden)
	} else {
		ws, _ := websocket.Upgrade(w, r, nil, 0, 0)
		vars := mux.Vars(r)
		boardID, _ := strconv.Atoi(vars["id"])
		if board, ok := all_boards.GetByID(boardID); ok {
			RegNewBoardObserver(ws, board.ID, session_info.GetUserID(r))
		}
	}
}

type AdminInfo struct {
	IsAdmin   bool
	Observers []all_boards.Observer
}

type ObserverInfo struct {
	IsAdmin bool
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
			var info interface{}
			if all_boards.IsAdmin(session_info.GetUserID(r), boardID) {
				if b, ok := all_boards.GetByID(boardID); ok {
					info = AdminInfo{true, b.Observers}
				}
			} else {
				info = ObserverInfo{false}
			}
			tmpl, _ := template.ParseFiles("./templates/drawing_board.html")
			tmpl.Execute(w, info)
		}
	}
}

type sockClient struct {
	userID  int
	boardID int
	sock    *websocket.Conn
	mu      sync.Mutex
}

//canvasMessage is a struct
type canvasMessage struct {
	Points []all_boards.Point `json:"points"`
}

type boardPageSetup struct {
	Type         string    `json:"type"`
	History      []float64 `json:"history"`
	PublicOffers int       `json:"publicOffers"`
	PersonOffers int       `json:"personOffers"`
}

var clients = make(map[int]*sockClient)

//SendtoBoardObservers is func
func SendtoBoardObservers(boardID, originUser int, message interface{}) {
	for _, client := range clients {
		if client.boardID == boardID && client.userID != originUser {
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
			all_boards.NewDrawing(client.boardID, msg.Points)
			SendtoBoardObservers(client.boardID, client.userID, msg)
		} else {
			break
		}
	}
}

//RegNewBoardObserver is func
func RegNewBoardObserver(ws *websocket.Conn, boardID, userID int) {
	connID := connsinc.NewID()
	clients[connID] = &sockClient{userID: userID, boardID: boardID, sock: ws}
	if b, ok := all_boards.GetByID(boardID); ok {
		for _, p := range b.History {
			writeSingleMessage(connID, canvasMessage{Points: p})
		}
	}
	go procIncomingMessages(connID)
}
