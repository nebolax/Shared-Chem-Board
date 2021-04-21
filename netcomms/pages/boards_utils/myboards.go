package boards_utils

import (
	"ChemBoard/all_boards"
	"ChemBoard/netcomms/pages/account_logic"
	"html/template"
	"net/http"
)

func MyboardsPage(w http.ResponseWriter, r *http.Request) {
	if !account_logic.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		userId := account_logic.GetUserID(r)
		tmpl, _ := template.ParseFiles("./static/boards_utils/myboards.html")
		tmpl.Execute(w, all_boards.UserAdmin(userId))
	}
}
