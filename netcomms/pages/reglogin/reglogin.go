package reglogin

import (
	"ChemBoard/netcomms/pages/reglogin/usersinc"
	"ChemBoard/netcomms/session_info"
	"ChemBoard/status"
	"html/template"
	"net/http"
)

type DBUser struct {
	ID       int
	Login    string
	Email    string
	Password string
}

var users []*DBUser

func Logout(w http.ResponseWriter, r *http.Request) {
	session_info.SetUserID(w, r, 0)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ProcRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	inpLogin := r.PostForm.Get("login")
	inpPwd := r.PostForm.Get("password")
	inpEmail := r.PostForm.Get("email")
	id, cs := RegUser(inpLogin, inpEmail, inpPwd)
	switch cs {
	case status.OK:
		session_info.SetUserID(w, r, id)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	case status.UserAlreadyExists:
		tmpl, _ := template.ParseFiles("./templates/register.html")
		tmpl.Execute(w, "User already exists")
	}
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
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
		session_info.SetUserID(w, r, id)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	case status.NoSuchUser:
		tmpl.Execute(w, "user does not exist")
	case status.IncorrectPassword:
		tmpl.Execute(w, "incorrect password")
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./templates/login.html")
	tmpl.Execute(w, "")
}

//RegUser is func
func RegUser(login, email, pwd string) (int, status.StatusCode) {
	if userFromDB(login) != nil || userFromDB(email) != nil {
		return 0, status.UserAlreadyExists
	}
	id := usersinc.NewID()
	user := &DBUser{id, login, email, pwd}
	users = append(users, user)
	return id, status.OK
}

//LoginUser is func
func LoginUser(logmail, inpPwd string) (int, status.StatusCode) {
	if user := userFromDB(logmail); user == nil {
		return 0, status.NoSuchUser
	} else {
		if user.Password != inpPwd {
			return 0, status.IncorrectPassword
		} else {
			return user.ID, status.OK
		}
	}
}

func userFromDB(logmail string) *DBUser {
	for _, user := range users {
		if user.Login == logmail || user.Email == logmail {
			return user
		}
	}

	return nil
}
