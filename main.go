package main

import (
	"ChemBoard/database"
	_ "ChemBoard/database"
	"fmt"

	// "ChemBoard/netcomms"
	_ "ChemBoard/utils/configs"
)

func main() {
	// netcomms.StartServer()
	chout, chin := database.NewDbObserver()
	chout <- `select * from "TestTable"`
	rows := <-chin

	for rows.Next() {
		var id interface{}
		var login interface{}
		var email interface{}
		if err := rows.Scan(&id, &login, &email); err != nil {
			panic(err)
		}
		fmt.Printf("%v, %v, %v\n", id, login, email)
	}
	rows.Close()
}
