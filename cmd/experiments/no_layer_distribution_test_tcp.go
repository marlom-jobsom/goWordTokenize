package main

import (
	"fmt"
	"log"
	"time"
	"word-tokenize-in1118/cmd/experiments/util"
	"word-tokenize-in1118/internal/communication"
	"word-tokenize-in1118/internal/constant"
	"word-tokenize-in1118/internal/layers/infrastructure/client"
	"word-tokenize-in1118/internal/layers/infrastructure/server"
)

// TestRPC tests the RPC communication timing for the client_request_handler side
func main() {
	filePath := "no_layer_distribution_test_tcp_%s.txt"
	var responses []communication.Response
	var clientRequestHandler = new(client.RequestHandler)
	var serverRequestHandler = new(server.RequestHandler)

	go serverRequestHandler.BringUpTCPServer()
	time.Sleep(constant.SecondsSleep)

	for i := 0; i < constant.MaxTries; i++ {
		request := communication.Request{Content: constant.TextTest}
		response := clientRequestHandler.TextTokenizeTCP(request)
		responses = append(responses, response)
	}

	log.Println(util.SumResponsesDuration(responses))
	util.WriteResponsesDuration(fmt.Sprintf(filePath, time.Now().Format(constant.TimeLayout)), responses)
}
