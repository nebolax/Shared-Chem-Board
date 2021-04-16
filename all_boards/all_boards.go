package all_boards

import (
	"ChemBoard/all_boards/boardsinc"
	"sync"
)

var BoardsArray = []*Board{
	// {1, 1, "First", "x", []int{1, 2, 3}, sync.Mutex{}},
	// {2, 2, "Second", "y", []int{1, 2}, sync.Mutex{}},
	// {3, 1, "Third", "z", []int{2, 3}, sync.Mutex{}},
}

//TODO check if userid is valid
//TODO lock mutexes while working + boardsArray chould be private

func CreateBoard(adminID int, name, pwd string) int {
	nID := boardsinc.NewID()
	board := &Board{nID, adminID, name, pwd, []int{}, sync.Mutex{}}
	BoardsArray = append(BoardsArray, board)
	return nID
}

func GetByID(id int) *Board {
	for _, board := range BoardsArray {
		if board.ID == id {
			return board
		}
	}
	return nil
}

func SharedWithUser(userID int) []*Board {
	res := []*Board{}
	for _, b := range BoardsArray {
		for _, ux := range b.Users {
			if ux == userID {
				res = append(res, b)
				break
			}
		}
	}

	return res
}

func AvailableToUser(userID, boardID int) bool {
	userBoards := SharedWithUser(userID)

	if GetByID(boardID).Admin == userID {
		return true
	}

	for _, b := range userBoards {
		if b.ID == boardID {
			return true
		}
	}

	return false
}

func UserAdmin(userID int) []*Board {
	res := []*Board{}
	for _, b := range BoardsArray {
		if b.Admin == userID {
			res = append(res, b)
		}
	}

	return res
}

func IsAdmin(userID, boardID int) bool {
	return GetByID(boardID).Admin == userID
}
