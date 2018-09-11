package main

import (
	"word-tokenize-socket/core/communication"
)

func main() {
	server := new(communication.TokenizerServer)
	server.BringUpRPCTCPServer()
}
