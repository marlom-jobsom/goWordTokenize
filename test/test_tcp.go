package main

import (
	"log"
	"time"
	"word-tokenize-in1118/internal/constant"
	"word-tokenize-in1118/internal/core/communication"
	"word-tokenize-in1118/test/util"
)

// TestRPC tests the RPC communication timing for the client side
func main() {
	filePath := "test_tcp.txt"
	var responses []communication.Response
	client := new(communication.TokenizerClient)
	server := new(communication.TokenizerServer)

	go server.BringUpTCPServer()
	time.Sleep(2000000000)

	for i := 0; i < constant.MaxTries; i++ {
		response := client.TextTokenizeTCP(constant.TextTest)
		responses = append(responses, response)
	}

	log.Println(util.SumResponsesDuration(responses))
	util.WriteResponsesDuration(filePath, responses)
}
