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

//TODO check if user has permission to do actions

func prepChatMsg(dbChmsg all_boards.ChatMessage) chatMessage {
	uinfo, _ := account_logic.GetUserByID(dbChmsg.SenderID)
	return chatMessage{dbChmsg.ID,
		map[string]interface{}{
			"nickname": uinfo.Login,
		},
		dbChmsg.TimeStamp, dbChmsg.Content}
}

func procIncomingMessages(connID uint64) {
	ar := clients
	_ = ar
	for {
		cl := clients[connID]
		msg, ok := readSingleMessage(connID)
		if ok {
			switch typesMap[reflect.TypeOf(msg)] {
			case tAction:
				prm := msg.(all_boards.ActionMSG)
				newMsg, _ := all_boards.NewAction(cl.boardID(), curView(cl), prm)
				newGroupMessage(cl.boardID(), curView(cl), connID, newMsg)
				writeSingleMessage(connID, SetIdMSG{"action", newMsg.ID})
				writeSingleMessage(connID, SetIdMSG{"drawing", newMsg.Drawing.ID})
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

					newGroupMessage(cl.boardID(), curView(cl), 0, prepChatMsg(newMsg))
				}
			}
		} else {
			break
		}
	}
}

func curView(cl sockClient) uint64 {
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

func isAdminOnline(boardID uint64) (uint64, bool) {
	for _, cl := range clients {
		if all_boards.IsAdmin(cl.userID(), cl.boardID()) {
			return cl.userID(), true
		}
	}

	return 0, false
}

func updateObserversList(boardID uint64) {
	if adminID, admOn := isAdminOnline(boardID); admOn {
		sendtoUserDevices(adminID, 0, allObsStatMSG{curBoardObservers(boardID)})
	}
}

func curBoardObservers(boardID uint64) []singleObsInfo {
	ids := map[uint64]bool{}
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

func regNewBoardObserver(r *http.Request, ws *websocket.Conn, boardID, userID uint64) {
	connID := incrementor.Next("conns", false)
	if all_boards.IsAdmin(userID, boardID) {
		clients[connID] = adminClient{boardID, userID, 0, ws, &sync.Mutex{}}
		updateObserversList(boardID)
	} else {
		clients[connID] = observerClient{boardID, userID, false, ws, &sync.Mutex{}}
		updateObserversList(boardID)
	}
	sendHistory(connID, boardID, 0)

	go procIncomingMessages(connID)
}
