package all_boards

import (
	"ChemBoard/all_boards/boardsinc"
	"strings"
	"sync"
)

var BoardsArray = []*DataElem{
	// {1, 1, "First", "x", []int{1, 2, 3}, sync.Mutex{}},
	// {2, 2, "Second", "y", []int{1, 2}, sync.Mutex{}},
	// {3, 1, "Third", "z", []int{2, 3}, sync.Mutex{}},
}

//TODO check if userid is valid
//TODO lock mutexes while working + boardsArray chould be private

func CreateBoard(adminID int, name, pwd string) int {
	nID := boardsinc.NewID()
	board := &Board{nID, adminID, name, pwd, []Observer{}, [][]Point{}}
	BoardsArray = append(BoardsArray, &DataElem{board, sync.Mutex{}})
	return nID
}

func GetByID(id int) (Board, bool) {
	for _, el := range BoardsArray {
		if el.board.ID == id {
			return *el.board, true
		}
	}
	return Board{}, false
}

func pointerByID(id int) *Board {
	for _, el := range BoardsArray {
		if el.board.ID == id {
			return el.board
		}
	}
	return nil
}

func SharedWithUser(userID int) []Board {
	res := []Board{}
	for _, el := range BoardsArray {
		for _, obs := range el.board.Observers {
			if obs.UserID == userID {
				res = append(res, *el.board)
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
			res = append(res, *el.board)
		}
	}

	return res
}

func IsAdmin(userID, boardID int) bool {
	if b, ok := GetByID(boardID); ok && b.Admin == userID {
		return true
	}
	return false
}

func AddObserver(boardID, userID int, pwd string) bool {
	if b := pointerByID(boardID); b != nil {
		if b.Password == pwd {
			b.Observers = append(b.Observers, Observer{UserID: userID})
			return true
		}
	}

	return false
}

func BoardsWithoutUser(key string, userID int) []Board {
	res := []Board{}
	for _, el := range BoardsArray {
		if strings.Contains(el.board.Name, key) && !AvailableToUser(userID, el.board.ID) {
			res = append(res, *el.board)
		}
	}

	return res
}

func NewDrawing(boardID int, points []Point) {
	if b := pointerByID(boardID); b != nil {
		b.History = append(b.History, points)
	}
}
