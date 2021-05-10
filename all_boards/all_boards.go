package all_boards

import (
	"ChemBoard/database"
	"ChemBoard/netcomms/pages/account_logic"
	"ChemBoard/utils/incrementor"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/lib/pq"
)

var BoardsArray = []*DataElem{}

//TODO check if userid is valid
//TODO lock mutexes while working + boardsArray chould be private

func init() {
	dbActions := database.Query(`select * from "Actions"`, struct {
		ID        int64
		Type      int64
		DrawingID int64
	}{})
	dbBoards := database.Query(`select * from "Boards"`, struct {
		ID           int64
		AdminID      int64
		Name         string
		DrawingsIDs  pq.Int64Array
		Password     string
		ChatHistory  []uint8
		ObserversIDs []uint8
	}{})
	dbChmess := database.Query(`select * from "ChatsMessages"`, struct {
		ID          int64
		SenderID    int64
		TimeStampID int64
		ContentID   int64
	}{})
	dbDrawings := database.Query(`select * from "Drawings"`, struct {
		ID   int64
		Type int64
		Data string
	}{})
	dbMesscont := database.Query(`select * from "MessagesContent"`, struct {
		ID   int64
		Text string
	}{})
	dbObservers := database.Query(`select * from "Observers"`, struct {
		ID          int64
		DrawingHist int64
		ChatHist    int64
	}{})
	dbTimestamps := database.Query(`select * from "TimeStamps"`, struct {
		ID     int64
		Year   int64
		Month  int64
		Day    int64
		Hour   int64
		Minute int64
	}{})
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n%v\n\n", dbActions, dbBoards, dbChmess, dbDrawings, dbMesscont, dbObservers, dbTimestamps)
}

func CreateBoard(adminID int, name, pwd string) int {
	nID := incrementor.Next("boards", true)
	board := Board{nID, adminID, name, pwd, []*Observer{}, ActionsHistory{}, ChatHistory{}}
	BoardsArray = append(BoardsArray, &DataElem{board, sync.Mutex{}})
	return nID
}

func BoardByID(id int) (Board, bool) {
	for _, el := range BoardsArray {
		if el.board.ID == id {
			return el.board, true
		}
	}
	return Board{}, false
}

func (b *Board) obspointerByID(userID int) *Observer {
	for _, obs := range b.Observers {
		if obs.UserID == userID {
			return obs
		}
	}
	return nil
}

func (b Board) ObserverByID(userID int) (Observer, bool) {
	for _, obs := range b.Observers {
		if obs.UserID == userID {
			return *obs, true
		}
	}
	return Observer{}, false
}

func boardPointerByID(boardID int) *Board {
	for _, el := range BoardsArray {
		if el.board.ID == boardID {
			return &el.board
		}
	}
	return nil
}

func SharedWithUser(userID int) []Board {
	res := []Board{}
	for _, el := range BoardsArray {
		for _, obs := range el.board.Observers {
			if obs.UserID == userID {
				res = append(res, el.board)
				break
			}
		}
	}

	return res
}

func AvailableToUser(userID, boardID int) bool {
	userBoards := SharedWithUser(userID)

	if IsAdmin(userID, boardID) {
		return true
	}

	for _, b := range userBoards {
		if b.ID == boardID {
			return true
		}
	}

	return false
}

func UserAdmin(userID int) []Board {
	res := []Board{}
	for _, el := range BoardsArray {
		if el.board.Admin == userID {
			res = append(res, el.board)
		}
	}

	return res
}

func IsAdmin(userID, boardID int) bool {
	if b, ok := BoardByID(boardID); ok && b.Admin == userID {
		return true
	}
	return false
}

func AddObserver(boardID, userID int, pwd string) bool {
	if b := boardPointerByID(boardID); b != nil {
		if b.Password == pwd {
			b.Observers = append(b.Observers, &Observer{userID, ActionsHistory{}, ChatHistory{}})
			return true
		}
	}

	return false
}

func BoardsWithoutUser(key string, userID int) []Board {
	res := []Board{}
	for _, el := range BoardsArray {
		if strings.Contains(el.board.Name, key) && !AvailableToUser(userID, el.board.ID) {
			res = append(res, el.board)
		}
	}

	return res
}

func NewDrawing(boardID, viewID int, msg ActionMSG) (ActionMSG, bool) {
	actionID := incrementor.Next(fmt.Sprintf("Board%d-action", boardID), true)
	drawingID := incrementor.Next(fmt.Sprintf("Board%d-drawing", boardID), true)
	msg.ID = actionID
	msg.Drawing.ID = drawingID
	bar := BoardsArray
	_ = bar
	if b := boardPointerByID(boardID); b != nil {
		if viewID == 0 {
			b.Actions = append(b.Actions, msg)
		} else {
			if obs := b.obspointerByID(viewID); obs != nil {
				obs.Actions = append(obs.Actions, msg)
			}
		}
		return msg, true
	} else {
		return ActionMSG{}, false
	}
}

func DeleteDrawing(boardID, viewID, drawingID int) {
	if b := boardPointerByID(boardID); b != nil {
		if viewID == 0 {
			res := ActionsHistory{}
			for _, el := range b.Actions {
				if el.ID != drawingID {
					res = append(res, el)
				}
			}
			b.Actions = res
		} else {
			if obs := b.obspointerByID(viewID); obs != nil {
				res := ActionsHistory{}
				for _, el := range obs.Actions {
					if el.ID != drawingID {
						res = append(res, el)
					}
				}
				obs.Actions = res
			}
		}
	}
}

func curTimeStamp() TimeStamp {
	ct := time.Now()
	return TimeStamp{
		ct.Year(),
		int(ct.Month()),
		ct.Day(),
		ct.Hour(),
		ct.Minute(),
	}
}

func NewChatMessage(boardID, viewID, senderID int, content ChatContent) (ChatMessage, bool) {
	if user, ok := account_logic.GetUserByID(senderID); ok {
		timeStamp := curTimeStamp()
		msgID := incrementor.Next("chat-message", true)
		msg := ChatMessage{msgID, user, timeStamp, content}

		if b := boardPointerByID(boardID); b != nil {
			if viewID == 0 {
				b.Chat = append(b.Chat, msg)
			} else {
				if obs := b.obspointerByID(viewID); obs != nil {
					obs.Chat = append(obs.Chat, msg)
				} else {
					return ChatMessage{}, false
				}
			}
		} else {
			return ChatMessage{}, false
		}

		return msg, true
	}
	return ChatMessage{}, false
}
