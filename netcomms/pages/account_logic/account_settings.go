package account_logic

import (
	"html/template"
	"net/http"
)

func AccSettingsPage(w http.ResponseWriter, r *http.Request) {
	if IsUserLoggedIn(r) {
		tmpl, _ := template.ParseFiles("./static/account_logic/account_settings.html")
		tmpl.Execute(w, nil)
	}
}

func ChangePasswordPage(w http.ResponseWriter, r *http.Request) {
	if IsUserLoggedIn(r) {
		tmpl, _ := template.ParseFiles("./static/account_logic/change_password.html")
		tmpl.Execute(w, nil)
	}
}
