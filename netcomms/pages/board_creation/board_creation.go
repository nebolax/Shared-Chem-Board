package board_creation

import (
	"ChemBoard/all_boards"
	"ChemBoard/netcomms/session_info"
	"fmt"
	"html/template"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	if !session_info.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		tmpl, _ := template.ParseFiles("./templates/board_creation.html")
		tmpl.Execute(w, nil)
	}
}

func ProcCreation(w http.ResponseWriter, r *http.Request) {
	if !session_info.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		r.ParseForm()
		bName := r.PostForm.Get("name")
		bPwd := r.PostForm.Get("pwd")
		nID := all_boards.CreateBoard(session_info.GetUserID(r), bName, bPwd)
		http.Redirect(w, r, fmt.Sprintf("/board%d", nID), http.StatusSeeOther)
	}
}
