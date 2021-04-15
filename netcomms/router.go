package netcomms

import (
	boards "ChemBoard/netcomms/pages/board_page"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(router *mux.Router) {
	router.HandleFunc("/board{id:[0-9]+}", boards.BoardPage)
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
