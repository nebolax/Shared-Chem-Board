package general_pages

import (
	"ChemBoard/netcomms/pages/account_logic"
	"html/template"
	"net/http"
)

func PersonalHomePage(w http.ResponseWriter, r *http.Request) {
	if !account_logic.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		tmpl, _ := template.ParseFiles("./static/general_pages/personal_home.html")
		tmpl.Execute(w, nil)
	}
}
