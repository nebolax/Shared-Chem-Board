package boards_utils

import (
	"ChemBoard/all_boards"
	"ChemBoard/netcomms/pages/account_logic"
	"fmt"
	"html/template"
	"net/http"
)

func CreateBoardPage(w http.ResponseWriter, r *http.Request) {
	if !account_logic.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		tmpl, _ := template.ParseFiles("./static/boards_utils/board_creation.html")
		tmpl.Execute(w, nil)
	}
}

func ProcBoardCreation(w http.ResponseWriter, r *http.Request) {
	if !account_logic.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		r.ParseForm()
		bName := r.PostForm.Get("name")
		bPwd := r.PostForm.Get("pwd")
		nID := all_boards.CreateBoard(account_logic.GetUserID(r), bName, bPwd)
		http.Redirect(w, r, fmt.Sprintf("/board%d", nID), http.StatusSeeOther)
	}
}
