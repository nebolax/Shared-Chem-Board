package main

import (
	_ "ChemBoard/database"

	_ "ChemBoard/all_boards"
	"ChemBoard/netcomms"
	_ "ChemBoard/utils/configs"
)

func main() {
	netcomms.StartServer()
}
