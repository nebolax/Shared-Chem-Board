package board_joining

import (
	"ChemBoard/all_boards"
	"ChemBoard/netcomms/session_info"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Page(w http.ResponseWriter, r *http.Request) {
	if !session_info.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		vars := mux.Vars(r)
		boardID, _ := strconv.Atoi(vars["id"])
		tmpl, _ := template.ParseFiles("./templates/board_joining.html")
		b, _ := all_boards.GetByID(boardID)
		tmpl.Execute(w, b)
	}
}

func ProcJoining(w http.ResponseWriter, r *http.Request) {
	if !session_info.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		r.ParseForm()
		inpPwd := r.PostForm.Get("pwd")
		vars := mux.Vars(r)
		boardID, _ := strconv.Atoi(vars["id"])
		if all_boards.AddObserver(boardID, session_info.GetUserID(r), inpPwd) {
			http.Redirect(w, r, fmt.Sprintf("/board%d", boardID), http.StatusSeeOther)
		}
	}
}
