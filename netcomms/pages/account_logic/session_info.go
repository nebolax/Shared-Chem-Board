package account_logic

import (
	"ChemBoard/utils/configs"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   []byte
	store *sessions.CookieStore
)

func init() {
	conf_key := configs.Get("coockies-key")
	if conf_key == nil {
		nkey := 1000000000000000
		configs.Set("coockies-key", nkey)
		conf_key = nkey
	}
	num_key := conf_key.(float64)
	key = []byte(fmt.Sprint(num_key))
	store = sessions.NewCookieStore(key)
	num_key++
	configs.Set("coockies-key", num_key)
}

func SetUserInfo(w http.ResponseWriter, r *http.Request, info map[interface{}]interface{}) {
	session, _ := store.Get(r, "user-info")
	for key, val := range info {
		session.Values[key] = val
	}
	session.Save(r, w)
}

//GetUserInfo is func
func GetUserInfo(r *http.Request) map[interface{}]interface{} {
	session, _ := store.Get(r, "user-info")
	return session.Values
}

//IsUserLoggedIn is func
func IsUserLoggedIn(r *http.Request) bool {
	id := GetUserID(r)
	return id != 0
}

//SetUserID id func to set user id:D!!
func SetUserID(w http.ResponseWriter, r *http.Request, id int) {
	session, _ := store.Get(r, "user-info")
	session.Values["userid"] = id
	session.Save(r, w)
}

//GetUserID is func
func GetUserID(r *http.Request) int {
	session, _ := store.Get(r, "user-info")
	id, ok := session.Values["userid"].(int)
	if !ok {
		return 0
	}
	return id
}
