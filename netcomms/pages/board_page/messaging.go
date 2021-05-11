package board_page

import (
	"ChemBoard/all_boards"
)

var clients = make(map[uint64]sockClient)

func sendHistory(connID, boardID, viewID uint64) {
	if b, ok := all_boards.BoardByID(boardID); ok {
		var drawingsHist all_boards.ActionsHistory
		var chatHist all_boards.ChatHistory
		if viewID == 0 {
			drawingsHist = b.Actions
			chatHist = b.Chat
		} else {
			if obs, ok := b.ObserverByID(viewID); ok {
				drawingsHist = obs.Actions
				chatHist = obs.Chat
			}
		}
		for _, pack := range drawingsHist {
			writeSingleMessage(connID, pack)
		}
		for _, pack := range chatHist {
			writeSingleMessage(connID, prepChatMsg(pack))
		}
	}
}

func newGroupMessage(boardID, viewID, exceptConn uint64, msg interface{}) {
	for _, cl := range clients {
		if cl.boardID() == boardID {
			if cl.isAdmin() {
				if cl.(adminClient).dview == viewID {
					sendtoUserDevices(cl.userID(), exceptConn, msg)
				}
			} else {
				if (!cl.(observerClient).dview && viewID == 0) || (cl.(observerClient).dview && viewID == cl.(observerClient).duserID) {
					sendtoUserDevices(cl.userID(), exceptConn, msg)
				}
			}
		}
	}
}

func sendtoUserDevices(userID, exceptConn uint64, message interface{}) {
	for connID, client := range clients {
		if client.userID() == userID && connID != exceptConn {
			writeSingleMessage(connID, message)
		}
	}
}

func writeSingleMessage(connID uint64, msg interface{}) {
	clients[connID].mu().Lock()
	defer clients[connID].mu().Unlock()

	if enc, ok := encodeMessage(msg); ok {
		err := clients[connID].sock().WriteJSON(enc)
		if err != nil {
			delClient(connID)
		}
	}
}

func delClient(connID uint64) {
	boardID := clients[connID].boardID()
	clients[connID].sock().Close()
	delete(clients, connID)
	updateObserversList(boardID)
}

func readSingleMessage(connID uint64) (interface{}, bool) {
	var msg anyMSG
	err := clients[connID].sock().ReadJSON(&msg)
	if err != nil {
		delClient(connID)
		return 0, false
	}
	return decodeMessage(msg)
}
