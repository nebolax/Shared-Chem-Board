package account_logic

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   []byte
	store *sessions.CookieStore
)

func init() {
	key = []byte("sh-chemboard0103")
	store = sessions.NewCookieStore(key)
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
func SetUserID(w http.ResponseWriter, r *http.Request, id uint64) {
	session, _ := store.Get(r, "user-info")
	session.Values["userid"] = id
	session.Save(r, w)
}

//GetUserID is func
func GetUserID(r *http.Request) uint64 {
	session, _ := store.Get(r, "user-info")
	id, ok := session.Values["userid"].(uint64)
	if !ok {
		return 0
	}
	return id
}
