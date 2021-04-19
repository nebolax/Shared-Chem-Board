package session_info

import (
	"ChemBoard/status"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("chemboard-secre1")
	store = sessions.NewCookieStore(key)
)

//GetUserInfo is func
func GetUserInfo(r *http.Request) (int, status.StatusCode) {
	id := GetUserID(r)

	return id, status.OK
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
