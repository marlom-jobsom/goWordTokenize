package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
	"word-tokenize-in1118/cmd/experiments/util"
	"word-tokenize-in1118/internal/communication"
	"word-tokenize-in1118/internal/constant"
	"word-tokenize-in1118/internal/services/nlg"
)

// TestRPC tests the RPC communication timing for the client_request_handler side
func main() {
	filePath := "no_infrastructure_test_tcp_%s.txt"
	var responses []communication.Response

	go bringUpTCPServer()
	time.Sleep(constant.SecondsSleep)

	for i := 0; i < constant.MaxTries; i++ {
		request := communication.Request{Content: constant.TextTest}
		connection, _ := net.Dial(constant.TCP, constant.PORT)

		log.Println(fmt.Sprintf(constant.SendingRequest, constant.TCP), request)
		now := time.Now()
		json.NewEncoder(connection).Encode(request)

		var response communication.Response
		json.NewDecoder(connection).Decode(&response)
		elapsed := time.Since(now)
		response.Duration = elapsed
		connection.Close()

		log.Println(fmt.Sprintf(constant.ReceivingResponse, constant.TCP), response)
		responses = append(responses, response)
	}

	log.Println(util.SumResponsesDuration(responses))
	util.WriteResponsesDuration(fmt.Sprintf(filePath, time.Now().Format(constant.TimeLayout)), responses)
}

func bringUpTCPServer() {
	log.Println("Bring up server over TCP")
	tcpAddress, _ := net.ResolveTCPAddr(constant.TCP, constant.PORT)
	listen, _ := net.ListenTCP(constant.TCP, tcpAddress)
	defer listen.Close()

	log.Println(constant.Address, listen.Addr())

	for {
		connection, _ := listen.Accept()
		jsonEncoder := json.NewEncoder(connection)
		jsonDecoder := json.NewDecoder(connection)

		var request communication.Request
		jsonDecoder.Decode(&request)
		log.Println(fmt.Sprintf(constant.ReceivingRequest, constant.TCP), request)

		tokens := nlg.TextTokenize(request.Content)

		// Sending back
		log.Println(fmt.Sprintf(constant.SendingResponse, constant.TCP), tokens)
		jsonEncoder.Encode(communication.Response{Content: tokens})

		connection.Close()
	}
}
