package netcomms

import (
	"ChemBoard/netcomms/pages/account_settings"
	"ChemBoard/netcomms/pages/available_boards"
	"ChemBoard/netcomms/pages/board_creation"
	"ChemBoard/netcomms/pages/board_joining"
	"ChemBoard/netcomms/pages/change_password"
	"ChemBoard/netcomms/pages/drawing_board"
	"ChemBoard/netcomms/pages/landing"
	"ChemBoard/netcomms/pages/myboards"
	"ChemBoard/netcomms/pages/personal_home"
	"ChemBoard/netcomms/pages/reglogin"
	"ChemBoard/netcomms/pages/search_board"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(router *mux.Router) {
	router.HandleFunc("/", landing.Page)
	router.HandleFunc("/board{id:[0-9]+}", drawing_board.Page)

	router.HandleFunc("/login", reglogin.LoginPage).Methods("GET")
	router.HandleFunc("/register", reglogin.RegisterPage).Methods("GET")
	router.HandleFunc("/login", reglogin.ProcLogin).Methods("POST")
	router.HandleFunc("/register", reglogin.ProcRegister).Methods("POST")
	router.HandleFunc("/logout", reglogin.Logout)

	router.HandleFunc("/shared-with-me", available_boards.Page)
	router.HandleFunc("/myboards", myboards.Page)
	router.HandleFunc("/home", personal_home.Page)
	router.HandleFunc("/account-settings", account_settings.Page)
	router.HandleFunc("/change-password", change_password.Page)

	router.HandleFunc("/newboard", board_creation.Page).Methods("GET")
	router.HandleFunc("/newboard", board_creation.ProcCreation).Methods("POST")

	router.HandleFunc("/search-board", search_board.Page).Methods("GET")
	router.HandleFunc("/search-board", search_board.ProcSearching).Methods("POST")

	router.HandleFunc("/join-board{id:[0-9]+}", board_joining.Page).Methods("GET")
	router.HandleFunc("/join-board{id:[0-9]+}", board_joining.ProcJoining).Methods("POST")
}

func StartServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./templates"))))

	setupRoutes(router)

	router.HandleFunc("/ws/board{id:[0-9]+}", drawing_board.HandleSockets)
	http.Handle("/", router)

	log.Println("starting http server at port 8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
