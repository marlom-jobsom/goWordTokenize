package main

import (
	"fmt"
	"log"
	"time"
	"word-tokenize-in1118/cmd/experiments/util"
	"word-tokenize-in1118/internal/communication"
	"word-tokenize-in1118/internal/constant"
	"word-tokenize-in1118/internal/infrastructure/client_request_handler/proxy"
	"word-tokenize-in1118/internal/infrastructure/server_request_handler/requesthandler"
)

// TestRPC tests the RPC communication timing for the client_request_handler side
func main() {
	filePath := "infrastructure_test_tcp_%s.txt"
	var responses []communication.Response
	clientProxy := new(proxy.ClientProxy)
	requestHandler := new(requesthandler.RequestHandler)

	go requestHandler.BringUpTCPServer()
	time.Sleep(constant.SecondsSleep)

	for i := 0; i < constant.MaxTries; i++ {
		response := clientProxy.InvokeTextTokenize(constant.TCP, false, constant.TextTest)
		responses = append(responses, response)
	}

	log.Println(util.SumResponsesDuration(responses))
	util.WriteResponsesDuration(fmt.Sprintf(filePath, time.Now().Format(constant.TimeLayout)), responses)
}
