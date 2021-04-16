package personal_home

import (
	"ChemBoard/netcomms/session_info"
	"html/template"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	if !session_info.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		tmpl, _ := template.ParseFiles("./templates/personal_home.html")
		tmpl.Execute(w, nil)
	}
}
