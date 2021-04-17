package search_board

import (
	"ChemBoard/all_boards"
	"ChemBoard/netcomms/session_info"
	"html/template"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	if !session_info.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		tmpl, _ := template.ParseFiles("./templates/search_board.html")
		tmpl.Execute(w, nil)
	}
}

func ProcSearching(w http.ResponseWriter, r *http.Request) {
	if !session_info.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		r.ParseForm()
		searchKey := r.PostForm.Get("key")
		searchRes := all_boards.BoardsWithoutUser(searchKey, session_info.GetUserID(r))
		tmpl, _ := template.ParseFiles("./templates/search_board.html")
		tmpl.Execute(w, searchRes)
	}
}
