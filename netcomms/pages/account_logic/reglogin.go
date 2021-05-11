package account_logic

import (
	"ChemBoard/database"
	"ChemBoard/utils/incrementor"
	"ChemBoard/utils/status"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

//TODO user info should be taken from database, not from coockies
type DBUser struct {
	ID       uint64
	Login    string
	Email    string
	PassHash string
}

var users []DBUser

func dconv(inp []interface{}) []DBUser {
	res := []DBUser{}
	for _, el := range inp {
		res = append(res, el.(DBUser))
	}
	return res
}

func init() {
	users = dconv(database.Query(`select * from "UsersInfo"`, DBUser{}))
}

func GetUserByID(userID uint64) (DBUser, bool) {
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
	user, cs := LoginUser(inpLogmail, inpPwd)
	tmpl, _ := template.ParseFiles("./static/account_logic/login.html")
	switch cs {
	case status.OK:
		SetUserID(w, r, user.ID)
		SetUserInfo(w, r, map[interface{}]interface{}{"login": user.Login, "email": user.Email})
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	case status.IncorrectLogPass:
		tmpl.Execute(w, "Логин или пароль неверны")
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./static/account_logic/login.html")
	tmpl.Execute(w, "")
}

//RegUser is func
func RegUser(login, email, pwd string) (uint64, status.StatusCode) {
	_, exbylog := userFromDB(login)
	_, exbymemail := userFromDB(email)
	if exbylog || exbymemail {
		return 0, status.UserAlreadyExists
	}
	id := incrementor.Next("users", true)
	pb, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	passhash := string(pb)
	if err != nil {
		panic(err)
	}
	user := DBUser{id, login, email, passhash}
	users = append(users, user)
	database.Query(`insert into "UsersInfo" values($1, $2, $3, $4)`, 0, id, login, email, passhash)
	return id, status.OK
}

//LoginUser is func
func LoginUser(logmail, inpPwd string) (DBUser, status.StatusCode) {
	if duser, ok := userFromDB(logmail); ok {
		return DBUser{}, status.IncorrectLogPass
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(duser.PassHash), []byte(inpPwd)); err != nil {
			return DBUser{}, status.IncorrectLogPass
		} else {
			return duser, status.OK
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

func UserLogin(userID uint64) string {
	for _, user := range users {
		if user.ID == userID {
			return user.Login
		}
	}
	return ""
}
