package landing

import (
	"ChemBoard/netcomms/session_info"
	"net/http"
	"text/template"
)

func Page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./templates/landing.html")
	isAuthed := session_info.IsUserLoggedIn(r)
	tmpl.Execute(w, isAuthed)
}
