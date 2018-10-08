package main

import (
	"fmt"
	"log"
	"time"
	"word-tokenize-in1118/cmd/experiments/util"
	"word-tokenize-in1118/internal/communication"
	"word-tokenize-in1118/internal/constant"
	"word-tokenize-in1118/internal/layers/distribution/client"
	"word-tokenize-in1118/internal/layers/infrastructure/server"
)

// TestRPC tests the RPC communication timing for the client_request_handler side
func main() {
	filePath := "layers_test_udp_%s.txt"
	var responses []communication.Response
	clientProxy := new(client.Proxy)
	requestHandler := new(server.RequestHandler)

	go requestHandler.BringUpUDPServer()
	time.Sleep(constant.SecondsSleep)

	for i := 0; i < constant.MaxTries; i++ {
		response := clientProxy.InvokeTextTokenize(constant.UDP, false, constant.TextTest)
		responses = append(responses, response)
	}

	log.Println(util.SumResponsesDuration(responses))
	util.WriteResponsesDuration(fmt.Sprintf(filePath, time.Now().Format(constant.TimeLayout)), responses)
}
