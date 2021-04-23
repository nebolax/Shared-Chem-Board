package board_page

import (
	"ChemBoard/all_boards"
)

var clients = make(map[int]sockClient)

func adminOfBoard(boardID int) *adminClient {
	for _, cl := range clients {
		if cl.boardID() == boardID {
			return cl.(*adminClient)
		}
	}
	return nil
}

func observersOfBoard(boardID int) []*observerClient {
	res := []*observerClient{}
	for _, cl := range clients {
		if cl.boardID() == boardID {
			res = append(res, cl.(*observerClient))
		}
	}
	return res
}

func newDrawing(boardID, viewID, exceptConn int, msg canvasMessage) {
	all_boards.NewDrawing(boardID, viewID, msg.Points)

	for _, cl := range clients {
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

func sendtoUserDevices(userID, exceptConn int, message interface{}) {
	for connID, client := range clients {
		if client.userID() == userID && connID != exceptConn {
			writeSingleMessage(connID, message)
		}
	}
}

func writeSingleMessage(connID int, message interface{}) {
	clients[connID].mu().Lock()
	err := clients[connID].sock().WriteJSON(message)
	if err != nil {
		delClient(connID)
	}
	clients[connID].mu().Unlock()
}

func delClient(connID int) {
	clients[connID].sock().Close()
	delete(clients, connID)
}

func readSingleMessage(connID int) (canvasMessage, bool) {
	var msg canvasMessage
	err := clients[connID].sock().ReadJSON(&msg)
	if err != nil {
		delClient(connID)
		return canvasMessage{}, false
	}
	return msg, true
}
