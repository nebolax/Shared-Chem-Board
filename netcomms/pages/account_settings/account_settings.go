package account_settings

import (
	"ChemBoard/netcomms/session_info"
	"html/template"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	if session_info.IsUserLoggedIn(r) {
		tmpl, _ := template.ParseFiles("./templates/account_settings.html")
		tmpl.Execute(w, nil)
	}
}
