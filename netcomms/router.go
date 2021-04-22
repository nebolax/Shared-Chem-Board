package netcomms

import (
	"ChemBoard/netcomms/pages/account_logic"
	"ChemBoard/netcomms/pages/board_page"
	"ChemBoard/netcomms/pages/boards_utils"
	"ChemBoard/netcomms/pages/general_pages"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(router *mux.Router) {
	router.HandleFunc("/", general_pages.LandingPage)
	router.HandleFunc("/home", general_pages.PersonalHomePage)
	router.HandleFunc("/board{id:[0-9]+}", board_page.Page)

	router.HandleFunc("/login", account_logic.LoginPage).Methods("GET")
	router.HandleFunc("/login", account_logic.ProcLogin).Methods("POST")
	router.HandleFunc("/register", account_logic.RegisterPage).Methods("GET")
	router.HandleFunc("/register", account_logic.ProcRegister).Methods("POST")
	router.HandleFunc("/logout", account_logic.Logout)

	router.HandleFunc("/account-settings", account_logic.AccSettingsPage)
	router.HandleFunc("/change-password", account_logic.ChangePasswordPage)

	router.HandleFunc("/shared-with-me", boards_utils.AvailableBoardsPage)
	router.HandleFunc("/myboards", boards_utils.MyboardsPage)

	router.HandleFunc("/newboard", boards_utils.CreateBoardPage).Methods("GET")
	router.HandleFunc("/newboard", boards_utils.ProcBoardCreation).Methods("POST")

	router.HandleFunc("/search-board", boards_utils.SearchBoardPage).Methods("GET")
	router.HandleFunc("/search-board", boards_utils.ProcBoardSearching).Methods("POST")

	router.HandleFunc("/join-board{id:[0-9]+}", boards_utils.JoinBoardPage).Methods("GET")
	router.HandleFunc("/join-board{id:[0-9]+}", boards_utils.ProcBoardJoining).Methods("POST")
}

func StartServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	setupRoutes(router)

	router.HandleFunc("/ws/board{id:[0-9]+}", board_page.HandleSockets)
	http.Handle("/", router)

	log.Println("starting http server at port 8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
