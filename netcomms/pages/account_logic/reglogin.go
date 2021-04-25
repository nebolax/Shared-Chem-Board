package account_logic

import (
	"ChemBoard/utils/incrementor"
	"ChemBoard/utils/status"
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

func GetUserByID(userID int) (DBUser, bool) {
	for _, user := range users {
		if user.ID == userID {
			return *user, true
		}
	}
	return DBUser{}, false
}

func Logout(w http.ResponseWriter, r *http.Request) {
	SetUserID(w, r, 0)
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
		SetUserID(w, r, id)
		SetUserInfo(w, r, map[interface{}]interface{}{"login": inpLogin, "email": inpEmail})
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	case status.UserAlreadyExists:
		tmpl, _ := template.ParseFiles("./static/account_logic/register.html")
		tmpl.Execute(w, "User already exists")
	}
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./static/account_logic/register.html")
	tmpl.Execute(w, "")
}

func ProcLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	inpLogmail := r.PostForm.Get("login")
	inpPwd := r.PostForm.Get("password")
	id, cs := LoginUser(inpLogmail, inpPwd)
	tmpl, _ := template.ParseFiles("./static/account_logic/login.html")
	switch cs {
	case status.OK:
		SetUserID(w, r, id)
		u := userFromDB(inpLogmail)
		SetUserInfo(w, r, map[interface{}]interface{}{"login": u.Login, "email": u.Email})
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	case status.NoSuchUser:
		tmpl.Execute(w, "user does not exist")
	case status.IncorrectPassword:
		tmpl.Execute(w, "incorrect password")
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./static/account_logic/login.html")
	tmpl.Execute(w, "")
}

//RegUser is func
func RegUser(login, email, pwd string) (int, status.StatusCode) {
	if userFromDB(login) != nil || userFromDB(email) != nil {
		return 0, status.UserAlreadyExists
	}
	id := incrementor.Next("users")
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

func UserLogin(userID int) string {
	for _, user := range users {
		if user.ID == userID {
			return user.Login
		}
	}
	return ""
}
