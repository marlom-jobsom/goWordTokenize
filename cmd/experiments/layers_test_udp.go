package main

import (
	"fmt"
	"github.com/marlom-jobsom/goWordTokenize/cmd/experiments/util"
	"github.com/marlom-jobsom/goWordTokenize/internal/communication"
	"github.com/marlom-jobsom/goWordTokenize/internal/constant"
	"github.com/marlom-jobsom/goWordTokenize/internal/layers/distribution/client"
	"github.com/marlom-jobsom/goWordTokenize/internal/layers/infrastructure/server"
	"log"
	"time"
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
