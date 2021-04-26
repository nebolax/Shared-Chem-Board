package board_page

import (
	"ChemBoard/all_boards"
	"ChemBoard/netcomms/pages/account_logic"
	"ChemBoard/utils/incrementor"
	"net/http"
	"reflect"
	"sync"

	"github.com/gorilla/websocket"
)

func procIncomingMessages(connID int) {
	ar := clients
	_ = ar
	for {
		cl := clients[connID]
		msg, ok := readSingleMessage(connID)
		if ok {
			switch typesMap[reflect.TypeOf(msg)] {
			case tPoints:
				all_boards.NewDrawing(cl.boardID(), curView(cl), msg.(pointsMSG).Points)
				newGroupMessage(cl.boardID(), curView(cl), connID, msg.(pointsMSG))
			case tChview:
				tms := msg.(chviewMSG)
				if cl.isAdmin() {
					nc := cl.(adminClient)
					nc.dview = tms.NView
					clients[connID] = nc
				} else {
					nc := cl.(observerClient)
					nc.dview = tms.NView == 1
					clients[connID] = nc
				}
				sendHistory(connID, cl.boardID(), curView(clients[connID]))
			case tInpChatMsg:
				msgContent := msg.(all_boards.ChatContent)
				if newMsg, ok := all_boards.NewChatMessage(cl.boardID(), curView(cl), cl.userID(), msgContent); ok {
					newGroupMessage(cl.boardID(), curView(cl), 0, newMsg)
				}
			}
		} else {
			break
		}
	}
}

func curView(cl sockClient) int {
	if cl.isAdmin() {
		return cl.(adminClient).dview
	} else {
		if !cl.(observerClient).dview {
			return 0
		} else {
			return cl.userID()
		}
	}
}

func isAdminOnline(boardID int) (int, bool) {
	for _, cl := range clients {
		if all_boards.IsAdmin(cl.userID(), cl.boardID()) {
			return cl.userID(), true
		}
	}

	return 0, false
}

func updateObserversList(boardID int) {
	if adminID, admOn := isAdminOnline(boardID); admOn {
		sendtoUserDevices(adminID, 0, allObsStatMSG{curBoardObservers(boardID)})
	}
}

func curBoardObservers(boardID int) []singleObsInfo {
	ids := map[int]bool{}
	for _, cl := range clients {
		if cl.boardID() == boardID && !cl.isAdmin() {
			if _, ok := ids[cl.userID()]; !ok {
				ids[cl.userID()] = true
			}
		}
	}
	res := []singleObsInfo{}
	for id, _ := range ids {
		res = append(res, singleObsInfo{id, account_logic.UserLogin(id)})
	}
	return res
}

func regNewBoardObserver(r *http.Request, ws *websocket.Conn, boardID, userID int) {
	connID := incrementor.Next("conns")
	if all_boards.IsAdmin(userID, boardID) {
		clients[connID] = adminClient{boardID, userID, 0, ws, &sync.Mutex{}}
	} else {
		clients[connID] = observerClient{boardID, userID, false, ws, &sync.Mutex{}}
		updateObserversList(boardID)
	}
	sendHistory(connID, boardID, 0)

	go procIncomingMessages(connID)
}
