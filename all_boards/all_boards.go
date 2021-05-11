package all_boards

import (
	"ChemBoard/database"
	"ChemBoard/utils/incrementor"
	"encoding/json"
	"strings"
	"sync"
	"time"
)

var BoardsArray = []*DataElem{}

//TODO check if userid is valid
//TODO lock mutexes while working + boardsArray chould be private
//TODO load user info when sending chat messages + proc on client side

type sAction struct {
	ID        uint64
	Type      uint64
	DrawingID uint64
}

type sBoard struct {
	ID           uint64
	AdminID      uint64
	Name         string
	Password     string
	ActionsIDs   []uint64
	ObserversIDs []uint64
	ChatHistory  []uint64
}

type sChmess struct {
	ID          uint64
	SenderID    uint64
	TimeStampID uint64
	ContentID   uint64
}

type sDrawing struct {
	ID   uint64
	Type uint64
	Data string
}

type sMesscont struct {
	ID   uint64
	Text string
}

type sObserver struct {
	ID          uint64
	UserID      uint64
	ActionsHist []uint64
	ChatHist    []uint64
}

type sTimestamp struct {
	ID     uint64
	Year   uint64
	Month  uint64
	Day    uint64
	Hour   uint64
	Minute uint64
}

func fa(inp []interface{}) []sAction {
	res := []sAction{}
	for _, el := range inp {
		res = append(res, el.(sAction))
	}
	return res
}

func fb(inp []interface{}) []sBoard {
	res := []sBoard{}
	for _, el := range inp {
		res = append(res, el.(sBoard))
	}
	return res
}

func fc(inp []interface{}) []sChmess {
	res := []sChmess{}
	for _, el := range inp {
		res = append(res, el.(sChmess))
	}
	return res
}

func fd(inp []interface{}) []sDrawing {
	res := []sDrawing{}
	for _, el := range inp {
		res = append(res, el.(sDrawing))
	}
	return res
}

func fm(inp []interface{}) []sMesscont {
	res := []sMesscont{}
	for _, el := range inp {
		res = append(res, el.(sMesscont))
	}
	return res
}

func fo(inp []interface{}) []sObserver {
	res := []sObserver{}
	for _, el := range inp {
		res = append(res, el.(sObserver))
	}
	return res
}

func ft(inp []interface{}) []sTimestamp {
	res := []sTimestamp{}
	for _, el := range inp {
		res = append(res, el.(sTimestamp))
	}
	return res
}

func init() {
	dbActions := fa(database.Query(`select * from "Actions"`, sAction{}))
	dbBoards := fb(database.Query(`select * from "Boards"`, sBoard{}))
	dbChmess := fc(database.Query(`select * from "ChatsMessages"`, sChmess{}))
	dbDrawings := fd(database.Query(`select * from "Drawings"`, sDrawing{}))
	dbMesscont := fm(database.Query(`select * from "MessagesContent"`, sMesscont{}))
	dbObservers := fo(database.Query(`select * from "Observers"`, sObserver{}))
	dbTimestamps := ft(database.Query(`select * from "TimeStamps"`, sTimestamp{}))

	// fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n%v\n\n", dbActions, dbBoards, dbChmess, dbDrawings, dbMesscont, dbObservers, dbTimestamps)
	loadRealData(dbActions, dbBoards, dbChmess, dbDrawings, dbMesscont, dbObservers, dbTimestamps)
}

func loadDrawing(dbDrawings []sDrawing, id uint64) Drawing {
	for _, dbdr := range dbDrawings {
		if dbdr.ID == id {
			var data interface{}
			err := json.Unmarshal([]byte(dbdr.Data), &data)
			if err != nil {
				panic(err)
			}
			return Drawing{dbdr.ID, dbdr.Type, data}
		}
	}
	return Drawing{}
}

func loadActions(dbActions []sAction, dbDrawings []sDrawing, actIds []uint64) []ActionMSG {
	res := []ActionMSG{}
	for _, dbac := range dbActions {
		cont := false
		for _, actid := range actIds {
			if actid == dbac.ID {
				cont = true
				break
			}
		}
		if cont {
			res = append(res, ActionMSG{dbac.ID, dbac.Type, loadDrawing(dbDrawings, dbac.DrawingID)})
		}
	}
	return res
}

func loadMesscont(dbMesscont []sMesscont, id uint64) ChatContent {
	for _, dbmc := range dbMesscont {
		if dbmc.ID == id {
			return ChatContent{dbmc.Text}
		}
	}
	return ChatContent{}
}

func loadTimestamp(dbTimestamp []sTimestamp, id uint64) TimeStamp {
	for _, dbts := range dbTimestamp {
		if dbts.ID == id {
			return TimeStamp{dbts.Year, dbts.Month, dbts.Day, dbts.Hour, dbts.Minute}
		}
	}
	return TimeStamp{}
}

func loadChat(dbChmess []sChmess, dbMesscont []sMesscont, dbTimestamps []sTimestamp, msgIds []uint64) ChatHistory {
	res := ChatHistory{}
	for _, dbcm := range dbChmess {
		cont := false
		for _, msgid := range msgIds {
			if msgid == dbcm.ID {
				cont = true
				break
			}
		}
		if cont {
			res = append(res, ChatMessage{dbcm.ID, dbcm.SenderID, loadTimestamp(dbTimestamps, dbcm.TimeStampID), loadMesscont(dbMesscont, dbcm.ContentID)})
		}
	}
	return res
}

func loadObservers(dbObservers []sObserver, dbActions []sAction, dbDrawings []sDrawing, dbChmess []sChmess, dbMesscont []sMesscont, dbTimestamps []sTimestamp, obsIds []uint64) []*Observer {
	res := []*Observer{}
	for _, dbobs := range dbObservers {
		cont := false
		for _, obsid := range obsIds {
			if obsid == dbobs.ID {
				cont = true
				break
			}
		}
		if cont {
			res = append(res, &Observer{dbobs.ID, dbobs.UserID, loadActions(dbActions, dbDrawings, dbobs.ActionsHist), loadChat(dbChmess, dbMesscont, dbTimestamps, dbobs.ChatHist)})
		}
	}
	return res
}

func loadRealData(dbActions []sAction, dbBoards []sBoard, dbChmess []sChmess, dbDrawings []sDrawing, dbMesscont []sMesscont, dbObservers []sObserver, dbTimestamps []sTimestamp) {
	resBoards := []Board{}

	for _, dbbo := range dbBoards {
		resBoards = append(resBoards, Board{dbbo.ID, dbbo.AdminID, dbbo.Name, dbbo.Password,
			loadObservers(dbObservers, dbActions, dbDrawings, dbChmess, dbMesscont, dbTimestamps, dbbo.ObserversIDs),
			loadActions(dbActions, dbDrawings, dbbo.ActionsIDs),
			loadChat(dbChmess, dbMesscont, dbTimestamps, dbbo.ChatHistory)})
	}

	resDels := []*DataElem{}
	for _, el := range resBoards {
		resDels = append(resDels, &DataElem{el, sync.Mutex{}})
	}
	BoardsArray = resDels
}

func CreateBoard(adminID uint64, name, pwd string) uint64 {
	nID := incrementor.Next("boards", true)
	board := Board{nID, adminID, name, pwd, []*Observer{}, ActionsHistory{}, ChatHistory{}}
	database.Query(`insert into "Boards" values($1, $2, $3, $4, '', '', '')`, 0, nID, adminID, name, pwd)
	BoardsArray = append(BoardsArray, &DataElem{board, sync.Mutex{}})
	return nID
}

func BoardByID(id uint64) (Board, bool) {
	for _, el := range BoardsArray {
		if el.board.ID == id {
			return el.board, true
		}
	}
	return Board{}, false
}

func (b *Board) obspointerByID(userID uint64) *Observer {
	for _, obs := range b.Observers {
		if obs.UserID == userID {
			return obs
		}
	}
	return nil
}

func (b Board) ObserverByID(userID uint64) (Observer, bool) {
	for _, obs := range b.Observers {
		if obs.UserID == userID {
			return *obs, true
		}
	}
	return Observer{}, false
}

func boardPointerByID(boardID uint64) *Board {
	for _, el := range BoardsArray {
		if el.board.ID == boardID {
			return &el.board
		}
	}
	return nil
}

func SharedWithUser(userID uint64) []Board {
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

func AvailableToUser(userID, boardID uint64) bool {
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

func UserAdmin(userID uint64) []Board {
	res := []Board{}
	for _, el := range BoardsArray {
		if el.board.Admin == userID {
			res = append(res, el.board)
		}
	}

	return res
}

func IsAdmin(userID, boardID uint64) bool {
	if b, ok := BoardByID(boardID); ok && b.Admin == userID {
		return true
	}
	return false
}

func f1(obss []*Observer) []uint64 {
	res := []uint64{}
	for _, el := range obss {
		res = append(res, el.UserID)
	}
	return res
}

func AddObserver(boardID, userID uint64, pwd string) bool {
	if b := boardPointerByID(boardID); b != nil {
		if b.Password == pwd {
			nID := incrementor.Next("db-obsid", true)
			b.Observers = append(b.Observers, &Observer{nID, userID, ActionsHistory{}, ChatHistory{}})
			database.Query(`update "Boards" set "ObserversIDs" = $1 where "ID" = $2`, 0, f1(b.Observers), boardID)
			database.Query(`insert into "Observers" values($1, $2, '', '')`, 0, nID, userID)
			return true
		}
	}
	return false
}

func BoardsWithoutUser(key string, userID uint64) []Board {
	res := []Board{}
	for _, el := range BoardsArray {
		if strings.Contains(el.board.Name, key) && !AvailableToUser(userID, el.board.ID) {
			res = append(res, el.board)
		}
	}

	return res
}

func f2(actions []ActionMSG) []uint64 {
	res := []uint64{}
	for _, act := range actions {
		res = append(res, act.ID)
	}
	return res
}

func NewAction(boardID, viewID uint64, naction ActionMSG) (ActionMSG, bool) {
	if b := boardPointerByID(boardID); b != nil {
		actionID := incrementor.Next("Board-action", true)
		naction.ID = actionID
		if naction.Type == 0 {
			drawingID := incrementor.Next("Board-drawing", true)
			naction.Drawing.ID = drawingID
			drCont, err := json.Marshal(naction.Drawing.Data)
			if err != nil {
				panic(err)
			}
			database.Query(`insert into "Drawings" values($1, $2, $3)`, 0, naction.Drawing.ID, naction.Drawing.Type, drCont)
		} else {
			database.Query(`delete from "Drawings" where "ID" = $1`, 0, naction.Drawing.ID)
		}
		database.Query(`insert into "Actions" values($1, $2, $3)`, 0, naction.ID, naction.Type, naction.Drawing.ID)

		if viewID == 0 {
			b.Actions = append(b.Actions, naction)
			database.Query(`update "Boards" set "ActionsIDs" = $1 where "ID" = $2`, 0, f2(b.Actions), boardID)
		} else {
			if obs := b.obspointerByID(viewID); obs != nil {
				obs.Actions = append(obs.Actions, naction)
				database.Query(`update "Observers" set "ActionsHist" = $1 where "ID" = $2`, 0, f2(obs.Actions), obs.DBID)
			}
		}

		b := BoardsArray
		_ = b
		return naction, true
	} else {
		return ActionMSG{}, false
	}
}

func curTimeStamp() TimeStamp {
	ct := time.Now()
	return TimeStamp{
		uint64(ct.Year()),
		uint64(ct.Month()),
		uint64(ct.Day()),
		uint64(ct.Hour()),
		uint64(ct.Minute()),
	}
}

func f3(progChat ChatHistory) []uint64 {
	res := []uint64{}
	for _, el := range progChat {
		res = append(res, el.ID)
	}
	return res
}

func NewChatMessage(boardID, viewID, senderID uint64, content ChatContent) (ChatMessage, bool) {
	timeStamp := curTimeStamp()
	msgID := incrementor.Next("chat-message", true)
	msg := ChatMessage{msgID, senderID, timeStamp, content}

	if b := boardPointerByID(boardID); b != nil {
		if viewID == 0 {
			b.Chat = append(b.Chat, msg)
			database.Query(`update "Boards" set "ChatHistory" = $1 where "ID" = $2`, 0, f3(b.Chat), boardID)
		} else {
			if obs := b.obspointerByID(viewID); obs != nil {
				obs.Chat = append(obs.Chat, msg)
				database.Query(`update "Observers" set "ChatHist" = $1 where "ID" = $2`, 0, f3(obs.Chat), obs.DBID)
			} else {
				return ChatMessage{}, false
			}
		}
	} else {
		return ChatMessage{}, false
	}
	tsID := incrementor.Next("timestamps", true)
	mcID := incrementor.Next("mcontent", true)
	ts := msg.TimeStamp
	database.Query(`insert into "TimeStamps" values($1, $2, $3, $4, $5, $6)`, 0, tsID, ts.Year, ts.Month, ts.Day, ts.Hour, ts.Minute)
	database.Query(`insert into "MessagesContent" values($1, $2)`, 0, mcID, msg.Content.Text)
	database.Query(`insert into "ChatsMessages" values($1, $2, $3, $4)`, 0, msg.ID, senderID, tsID, mcID)
	return msg, true
}
