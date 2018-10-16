package main

import (
	"fmt"
	"github.com/marlom-jobsom/goWordTokenize/cmd/experiments/util"
	"github.com/marlom-jobsom/goWordTokenize/internal/communication"
	"github.com/marlom-jobsom/goWordTokenize/internal/constant"
	"github.com/marlom-jobsom/goWordTokenize/internal/layers/infrastructure/client"
	"github.com/marlom-jobsom/goWordTokenize/internal/layers/infrastructure/server"
	"log"
	"time"
)

// TestRPC tests the RPC communication timing for the client_request_handler side
func main() {
	filePath := "no_layer_distribution_test_rpc_%s.txt"
	var responses []communication.Response
	var clientRequestHandler = new(client.RequestHandler)
	var serverRequestHandler = new(server.RequestHandler)

	go serverRequestHandler.BringUpRPCTCPServer()
	time.Sleep(constant.SecondsSleep)

	for i := 0; i < constant.MaxTries; i++ {
		response := clientRequestHandler.TextTokenizeRPCTCP(constant.TextTest)
		responses = append(responses, response)
	}

	log.Println(util.SumResponsesDuration(responses))
	util.WriteResponsesDuration(fmt.Sprintf(filePath, time.Now().Format(constant.TimeLayout)), responses)
}
