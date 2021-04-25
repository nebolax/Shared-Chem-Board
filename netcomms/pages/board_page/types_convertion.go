package board_page

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
)

var (
	typesMap = map[reflect.Type]msgType{
		reflect.TypeOf(pointsMSG{}):     tPoints,
		reflect.TypeOf(allObsStatMSG{}): tObsStat,
		reflect.TypeOf(chviewMSG{}):     tChview,
	}
)

func decodeMessage(msg anyMSG) (interface{}, bool) {
	switch msg.Type {
	case tPoints:
		var mstr pointsMSG
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
	default:
		return 0, false
	}
}

func encodeMessage(msg interface{}) (anyMSG, bool) {
	switch typesMap[reflect.TypeOf(msg)] {
	case tPoints:
		return anyMSG{tPoints, msg}, true
	case tChview:
		return anyMSG{tChview, msg}, true
	case tObsStat:
		return anyMSG{tObsStat, msg}, true
	}
	return anyMSG{}, false
}
