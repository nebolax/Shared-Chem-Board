package available_boards

import (
	"ChemBoard/all_boards"
	"ChemBoard/netcomms/session_info"
	"fmt"
	"net/http"
	"text/template"
)

func MyBoardsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Page sent to user %d\n", session_info.GetSessionUserID(r))
	tmpl, _ := template.ParseFiles("./templates/available_boards.html")
	tmpl.Execute(w, all_boards.BoardsOfUser(session_info.GetSessionUserID(r)))
}
