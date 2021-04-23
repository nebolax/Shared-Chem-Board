package board_page

import (
	"ChemBoard/all_boards"
	"ChemBoard/netcomms/pages/account_logic"
	"ChemBoard/utils/incrementor"
	"sync"

	"github.com/gorilla/websocket"
)

func procIncomingMessages(connID int) {
	cl := clients[connID]
	for {
		msg, ok := readSingleMessage(connID)
		if ok {
			var viewID int
			if cl.isAdmin() {
				viewID = cl.(adminClient).dview
			} else {
				if !cl.(observerClient).dview {
					viewID = 0
				} else {
					viewID = cl.userID()
				}
			}
			newDrawing(cl.boardID(), viewID, connID, msg)
		} else {
			break
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

func regNewBoardObserver(ws *websocket.Conn, boardID, userID int) {
	connID := incrementor.Next("conns")
	if all_boards.IsAdmin(userID, boardID) {
		clients[connID] = adminClient{boardID, userID, 0, ws, &sync.Mutex{}}
		if b, ok := all_boards.BoardByID(boardID); ok {
			for _, pack := range b.History {
				writeSingleMessage(connID, canvasMessage{"points", pack})
			}
		}
	} else {
		clients[connID] = observerClient{boardID, userID, false, ws, &sync.Mutex{}}
		if adminID, admOn := isAdminOnline(boardID); admOn {
			if user, ok := account_logic.GetUserByID(userID); ok {
				sendtoUserDevices(adminID, 0, newObserver{"newObserver", userID, user.Login})
			}
		}
	}

	go procIncomingMessages(connID)
}
