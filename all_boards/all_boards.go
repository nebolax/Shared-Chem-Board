package all_boards

import "sync"

var BoardsArray = []*Board{
	{1, "First", []int{1, 2, 3}, sync.Mutex{}},
	{2, "Second", []int{1, 2}, sync.Mutex{}},
	{3, "Third", []int{2, 3}, sync.Mutex{}},
}

func GetByID(id int) *Board {
	for _, board := range BoardsArray {
		if board.ID == id {
			return board
		}
	}
	return nil
}

func BoardsOfUser(userID int) []*Board {
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
	userBoards := BoardsOfUser(userID)

	for _, b := range userBoards {
		if b.ID == boardID {
			return true
		}
	}

	return false
}
