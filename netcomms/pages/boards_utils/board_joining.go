package boards_utils

import (
	"ChemBoard/all_boards"
	"ChemBoard/netcomms/pages/account_logic"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func JoinBoardPage(w http.ResponseWriter, r *http.Request) {
	if !account_logic.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		vars := mux.Vars(r)
		boardID, _ := strconv.Atoi(vars["id"])
		tmpl, _ := template.ParseFiles("./static/boards_utils/board_joining.html")
		b, _ := all_boards.BoardByID(uint64(boardID))
		tmpl.Execute(w, b)
	}
}

func ProcBoardJoining(w http.ResponseWriter, r *http.Request) {
	if !account_logic.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		r.ParseForm()
		inpPwd := r.PostForm.Get("pwd")
		vars := mux.Vars(r)
		boardID, _ := strconv.Atoi(vars["id"])
		if all_boards.AddObserver(uint64(boardID), account_logic.GetUserID(r), inpPwd) {
			http.Redirect(w, r, fmt.Sprintf("/board%d", boardID), http.StatusSeeOther)
		}
	}
}
