package myboards

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
		userId := session_info.GetUserID(r)
		tmpl, _ := template.ParseFiles("./templates/myboards.html")
		tmpl.Execute(w, all_boards.UserAdmin(userId))
	}
}
