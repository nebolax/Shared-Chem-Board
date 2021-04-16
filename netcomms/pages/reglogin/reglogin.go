package reglogin

import (
	"ChemBoard/netcomms/pages/reglogin/usersinc"
	"ChemBoard/netcomms/session_info"
	"ChemBoard/status"
	"html/template"
	"net/http"
)

type dbUser struct {
	ID       int
	login    string
	email    string
	password string
}

var users []*dbUser

func Logout(w http.ResponseWriter, r *http.Request) {
	LogoutUser(w, r)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func ProcRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	inpLogin := r.PostForm.Get("login")
	inpPwd := r.PostForm.Get("password")
	inpEmail := r.PostForm.Get("email")
	id, cs := RegUser(inpLogin, inpEmail, inpPwd)
	switch cs {
	case status.OK:
		session_info.SetSessionUserID(w, r, id)
		http.Redirect(w, r, "/myboards", http.StatusSeeOther)
	case status.UserAlreadyExists:
		tmpl, _ := template.ParseFiles("./templates/register.html")
		tmpl.Execute(w, "User already exists")
	}
}

func GetRegisterHTML(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./templates/register.html")
	tmpl.Execute(w, "")
}

func ProcLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	inpLogin := r.PostForm.Get("login")
	inpPwd := r.PostForm.Get("password")
	id, cs := LoginUser(inpLogin, inpPwd)
	tmpl, _ := template.ParseFiles("./templates/login.html")
	switch cs {
	case status.OK:
		session_info.SetSessionUserID(w, r, id)
		http.Redirect(w, r, "/myboards", http.StatusSeeOther)
	case status.NoSuchUser:
		tmpl.Execute(w, "user does not exist")
	case status.IncorrectPassword:
		tmpl.Execute(w, "incorrect password")
	}
}

func GetLoginHTML(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./templates/login.html")
	tmpl.Execute(w, "")
}

//LogoutUser is func
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	session_info.SetSessionUserID(w, r, 0)
}

//RegUser is func
func RegUser(login, email, pwd string) (int, status.StatusCode) {
	if userFromDB(login) != nil || userFromDB(email) != nil {
		return 0, status.UserAlreadyExists
	}
	id := usersinc.NewID()
	user := &dbUser{id, login, email, pwd}
	users = append(users, user)
	return id, status.OK
}

//LoginUser is func
func LoginUser(logmail, inpPwd string) (int, status.StatusCode) {
	if user := userFromDB(logmail); user == nil {
		return 0, status.NoSuchUser
	} else {
		if user.password != inpPwd {
			return 0, status.IncorrectPassword
		} else {
			return user.ID, status.OK
		}
	}
}

func userFromDB(logmail string) *dbUser {
	for _, user := range users {
		if user.login == logmail || user.email == logmail {
			return user
		}
	}

	return nil
}
