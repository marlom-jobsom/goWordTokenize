package main

import (
	"log"
	"time"
	"word-tokenize-in1118/constant"
	"word-tokenize-in1118/core/communication"
	"word-tokenize-in1118/test/util"
)

// TestRPC tests the RPC communication timing for the client side
func main() {
	filePath := "test_udp.txt"
	var responses []communication.Response
	client := new(communication.TokenizerClient)
	server := new(communication.TokenizerServer)

	go server.BringUpUDPServer()
	time.Sleep(2000000000)

	for i := 0; i < constant.MaxTries; i++ {
		response := client.TextTokenizeUDP(constant.TextTest)
		responses = append(responses, response)
	}

	log.Println(util.SumResponsesDuration(responses))
	util.WriteResponsesDuration(filePath, responses)
}
