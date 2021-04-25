package main

import (
	"ChemBoard/netcomms"
	_ "ChemBoard/utils/configs"
)

func main() {

	netcomms.StartServer()
}
