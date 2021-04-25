package board_page

import (
	"ChemBoard/all_boards"
)

var clients = make(map[int]sockClient)

func sendHistory(connID, boardID, viewID int) {
	if b, ok := all_boards.BoardByID(boardID); ok {
		var hist [][]all_boards.Point
		if viewID == 0 {
			hist = b.History
		} else {
			if obs, ok := b.ObserverByID(viewID); ok {
				hist = obs.History
			}
		}
		for _, pack := range hist {
			writeSingleMessage(connID, pointsMSG{pack})
		}
	}
}

func newDrawing(boardID, viewID, exceptConn int, msg pointsMSG) {
	all_boards.NewDrawing(boardID, viewID, msg.Points)

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

func sendtoUserDevices(userID, exceptConn int, message interface{}) {
	for connID, client := range clients {
		if client.userID() == userID && connID != exceptConn {
			writeSingleMessage(connID, message)
		}
	}
}

func writeSingleMessage(connID int, msg interface{}) {
	clients[connID].mu().Lock()
	defer clients[connID].mu().Unlock()

	if enc, ok := encodeMessage(msg); ok {
		err := clients[connID].sock().WriteJSON(enc)
		if err != nil {
			delClient(connID)
		}
	}
}

func delClient(connID int) {
	clients[connID].sock().Close()
	delete(clients, connID)
}

func readSingleMessage(connID int) (interface{}, bool) {
	var msg anyMSG
	err := clients[connID].sock().ReadJSON(&msg)
	if err != nil {
		delClient(connID)
		return 0, false
	}
	return decodeMessage(msg)
}
