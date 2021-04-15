package all_boards

var boardsArray = []*Board{{ID: 2}}

func GetByID(id int) *Board {
	for _, board := range boardsArray {
		if board.ID == id {
			return board
		}
	}
	return nil
}
