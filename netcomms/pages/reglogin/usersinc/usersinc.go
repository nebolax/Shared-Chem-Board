package usersinc

var curID = 0

func NewID() int {
	curID++
	return curID
}
