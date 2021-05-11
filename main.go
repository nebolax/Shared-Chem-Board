package main

import (
	_ "ChemBoard/database"
	"ChemBoard/netcomms"

	_ "ChemBoard/all_boards"
	_ "ChemBoard/utils/configs"
)

func main() {
	netcomms.StartServer()
}
