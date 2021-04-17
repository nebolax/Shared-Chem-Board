package change_password

import (
	"ChemBoard/netcomms/session_info"
	"net/http"
	"text/template"
)

func Page(w http.ResponseWriter, r *http.Request) {
	if session_info.IsUserLoggedIn(r) {
		tmpl, _ := template.ParseFiles("./templates/change_password.html")
		tmpl.Execute(w, nil)
	}
}
