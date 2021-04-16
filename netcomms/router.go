package netcomms

import (
	"ChemBoard/netcomms/pages/available_boards"
	boards "ChemBoard/netcomms/pages/board_page"
	"ChemBoard/netcomms/pages/reglogin"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(router *mux.Router) {
	router.HandleFunc("/board{id:[0-9]+}", boards.BoardPage)

	router.HandleFunc("/login", reglogin.GetLoginHTML).Methods("GET")
	router.HandleFunc("/register", reglogin.GetRegisterHTML).Methods("GET")
	router.HandleFunc("/login", reglogin.ProcLogin).Methods("POST")
	router.HandleFunc("/register", reglogin.ProcRegister).Methods("POST")

	router.HandleFunc("/myboards", available_boards.MyBoardsPage)
}

func StartServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./templates"))))

	setupRoutes(router)

	router.HandleFunc("/ws/board{id:[0-9]+}", boards.HandleSockets)
	http.Handle("/", router)

	log.Println("starting http server at port 8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
