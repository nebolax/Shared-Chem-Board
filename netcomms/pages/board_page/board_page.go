package board_page

import (
	"ChemBoard/all_boards"
	"ChemBoard/netcomms/pages/account_logic"
	"ChemBoard/utils/incrementor"
	"reflect"
	"sync"

	"github.com/gorilla/websocket"
)

func procIncomingMessages(connID int) {
	cl := clients[connID]
	for {
		msg, ok := readSingleMessage(connID)
		if ok {
			switch typesMap[reflect.TypeOf(msg)] {
			case tPoints:
				newDrawing(cl.boardID(), curView(cl), connID, msg.(pointsMSG))
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
				sendHistory(connID, cl.boardID(), curView(cl))
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

func regNewBoardObserver(ws *websocket.Conn, boardID, userID int) {
	connID := incrementor.Next("conns")
	if all_boards.IsAdmin(userID, boardID) {
		clients[connID] = adminClient{boardID, userID, 0, ws, &sync.Mutex{}}
	} else {
		clients[connID] = observerClient{boardID, userID, false, ws, &sync.Mutex{}}
		if adminID, admOn := isAdminOnline(boardID); admOn {
			if user, ok := account_logic.GetUserByID(userID); ok {
				sendtoUserDevices(adminID, 0, obsStatMSG{userID, user.Login})
			}
		}
	}
	sendHistory(connID, boardID, 0)

	go procIncomingMessages(connID)
}
