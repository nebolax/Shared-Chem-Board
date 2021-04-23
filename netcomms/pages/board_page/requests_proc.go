package board_page

import (
	"ChemBoard/all_boards"
	"ChemBoard/netcomms/pages/account_logic"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func HandleSockets(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", http.StatusForbidden)
	} else {
		ws, _ := websocket.Upgrade(w, r, nil, 0, 0)
		vars := mux.Vars(r)
		boardID, _ := strconv.Atoi(vars["id"])
		if board, ok := all_boards.BoardByID(boardID); ok {
			regNewBoardObserver(ws, board.ID, account_logic.GetUserID(r))
		}
	}
}

func Page(w http.ResponseWriter, r *http.Request) {
	if !account_logic.IsUserLoggedIn(r) {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		vars := mux.Vars(r)
		boardID, _ := strconv.Atoi(vars["id"])
		if !all_boards.AvailableToUser(account_logic.GetUserID(r), boardID) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			var tmpl *template.Template
			if all_boards.IsAdmin(account_logic.GetUserID(r), boardID) {
				tmpl, _ = template.ParseFiles("./static/board_page/admin_board.html")
			} else {
				tmpl, _ = template.ParseFiles("./static/board_page/observer_board.html")
			}

			tmpl.Execute(w, nil)
		}
	}
}
