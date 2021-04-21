package general_pages

import (
	"ChemBoard/netcomms/pages/account_logic"
	"net/http"
	"text/template"
)

func LandingPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./static/general_pages/landing.html")
	isAuthed := account_logic.IsUserLoggedIn(r)
	tmpl.Execute(w, isAuthed)
}
