package board_page

import (
	"ChemBoard/all_boards"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

var (
	typesMap = map[reflect.Type]msgType{
		reflect.TypeOf(SetIdMSG{}):               tSetId,
		reflect.TypeOf(all_boards.ActionMSG{}):   tDrawing,
		reflect.TypeOf(allObsStatMSG{}):          tObsStat,
		reflect.TypeOf(chviewMSG{}):              tChview,
		reflect.TypeOf(all_boards.ChatContent{}): tInpChatMsg,
		reflect.TypeOf(all_boards.ChatMessage{}): tOutChatMsg,
	}
)

func decodeMessage(msg anyMSG) (interface{}, bool) {
	switch msg.Type {
	case tDrawing:
		var mstr all_boards.ActionMSG
		err := mapstructure.Decode(msg.Data, &mstr)
		if err != nil {
			println(err.Error())
			return 0, false
		}
		return mstr, true
	case tChview:
		var mstr chviewMSG
		err := mapstructure.Decode(msg.Data, &mstr)
		if err != nil {
			println(err.Error())
			return 0, false
		}
		return mstr, true
	case tInpChatMsg:
		var mstr all_boards.ChatContent
		err := mapstructure.Decode(msg.Data, &mstr)
		if err != nil {
			println(err.Error())
			return 0, false
		}
		return mstr, true
	default:
		return 0, false
	}
}

func encodeMessage(msg interface{}) (anyMSG, bool) {
	switch typesMap[reflect.TypeOf(msg)] {
	case tSetId:
		return anyMSG{tSetId, msg}, true
	case tDrawing:
		return anyMSG{tDrawing, msg}, true
	case tChview:
		return anyMSG{tChview, msg}, true
	case tObsStat:
		return anyMSG{tObsStat, msg}, true
	case tOutChatMsg:
		return anyMSG{tOutChatMsg, msg}, true
	}
	return anyMSG{}, false
}
