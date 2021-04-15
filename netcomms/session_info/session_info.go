package session_info

import (
	"ChemBoard/statuses"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("chemboard-secret")
	store = sessions.NewCookieStore(key)
)

//GetSessionUserInfo is func
func GetSessionUserInfo(r *http.Request) (int, statuses.StatusCode) {
	id := GetSessionUserID(r)

	return id, statuses.OK
}

//IsUserLoggedIn is func
func IsUserLoggedIn(r *http.Request) bool {
	id := GetSessionUserID(r)
	return id != 0
}

//SetSessionUserID id func to set user id:D!!
func SetSessionUserID(w http.ResponseWriter, r *http.Request, id int) {
	session, _ := store.Get(r, "user-info")
	session.Values["userid"] = id
	session.Save(r, w)
}

//GetSessionUserID is func
func GetSessionUserID(r *http.Request) int {
	session, _ := store.Get(r, "user-info")
	id, ok := session.Values["userid"].(int)
	if !ok {
		return 0
	}
	return id
}
