package boards_utils

import (
	"ChemBoard/all_boards"
	"ChemBoard/netcomms/pages/account_logic"
	"net/http"
	"text/template"
)

func AvailableBoardsPage(w http.ResponseWriter, r *http.Request) {
	if !account_logic.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		tmpl, _ := template.ParseFiles("./static/boards_utils/available_boards.html")
		tmpl.Execute(w, all_boards.SharedWithUser(account_logic.GetUserID(r)))
	}
}
