package available_boards

import (
	"ChemBoard/all_boards"
	"ChemBoard/netcomms/session_info"
	"net/http"
	"text/template"
)

func Page(w http.ResponseWriter, r *http.Request) {
	if !session_info.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		tmpl, _ := template.ParseFiles("./templates/available_boards.html")
		tmpl.Execute(w, all_boards.SharedWithUser(session_info.GetUserID(r)))
	}
}
