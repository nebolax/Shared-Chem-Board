package boards_utils

import (
	"ChemBoard/all_boards"
	"ChemBoard/netcomms/pages/account_logic"
	"html/template"
	"net/http"
)

func SearchBoardPage(w http.ResponseWriter, r *http.Request) {
	if !account_logic.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		tmpl, _ := template.ParseFiles("./static/boards_utils/search_board.html")
		searchRes := all_boards.BoardsWithoutUser("", account_logic.GetUserID(r))

		tmpl.Execute(w, searchRes)
	}
}

func ProcBoardSearching(w http.ResponseWriter, r *http.Request) {
	if !account_logic.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		r.ParseForm()
		searchKey := r.PostForm.Get("key")
		searchRes := all_boards.BoardsWithoutUser(searchKey, account_logic.GetUserID(r))
		tmpl, _ := template.ParseFiles("./static/boards_utils/search_board.html")
		tmpl.Execute(w, searchRes)
	}
}
