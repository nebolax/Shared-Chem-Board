package account_logic

import (
	"ChemBoard/utils/incrementor"
	"ChemBoard/utils/status"
	"fmt"
	"html/template"
	"net/http"
)

//TODO user info should be taken from database, not from coockies

var (
	passwords = make(map[int]string)
)

type DBUser struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}

var users []DBUser

func GetUserByID(userID int) (DBUser, bool) {
	for _, user := range users {
		if user.ID == userID {
			return user, true
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
	inpEmail := r.PostForm.Get("email")
	inpPwd := r.PostForm.Get("password")
	fmt.Printf("%s, %s, %s", inpLogin, inpEmail, inpPwd)
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
	inpLogmail := r.PostForm.Get("logmail")
	inpPwd := r.PostForm.Get("password")
	user, cs := LoginUser(inpLogmail, inpPwd)
	tmpl, _ := template.ParseFiles("./static/account_logic/login.html")
	switch cs {
	case status.OK:
		SetUserID(w, r, user.ID)
		SetUserInfo(w, r, map[interface{}]interface{}{"login": user.Login, "email": user.Email})
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
	_, exbylog := userFromDB(login)
	_, exbymemail := userFromDB(email)
	if exbylog || exbymemail {
		return 0, status.UserAlreadyExists
	}
	id := incrementor.Next("users", true)
	user := DBUser{id, login, email}
	passwords[id] = pwd
	users = append(users, user)
	return id, status.OK
}

//LoginUser is func
func LoginUser(logmail, inpPwd string) (DBUser, status.StatusCode) {
	if user, ok := userFromDB(logmail); !ok {
		fmt.Println("No such user!")
		return DBUser{}, status.NoSuchUser
	} else {
		if passwords[user.ID] != inpPwd {
			return DBUser{}, status.IncorrectPassword
		} else {
			return user, status.OK
		}
	}
}

func userFromDB(logmail string) (DBUser, bool) {
	for _, user := range users {
		if user.Login == logmail || user.Email == logmail {
			return user, true
		}
	}

	return DBUser{}, false
}

func UserLogin(userID int) string {
	for _, user := range users {
		if user.ID == userID {
			return user.Login
		}
	}
	return ""
}
