package main

import (
	"word-tokenize-in1118/core/communication"
)

func main() {
	server := new(communication.TokenizerServer)
	server.BringUpTCPServer()
}
