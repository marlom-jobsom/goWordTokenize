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
	filePath := "no_infrastructure_test_udp_%s.txt"
	var responses []communication.Response

	go bringUpUDPServer()
	time.Sleep(constant.SecondsSleep)

	for i := 0; i < constant.MaxTries; i++ {
		request := communication.Request{Content: constant.TextTest}
		connection, _ := net.Dial(constant.UDP, constant.PORT)

		log.Println(fmt.Sprintf(constant.SendingRequest, constant.UDP), request)
		now := time.Now()
		json.NewEncoder(connection).Encode(request)

		var response communication.Response
		json.NewDecoder(connection).Decode(&response)
		elapsed := time.Since(now)
		response.Duration = elapsed
		connection.Close()

		log.Println(fmt.Sprintf(constant.ReceivingResponse, constant.UDP), response)
		responses = append(responses, response)
	}

	log.Println(util.SumResponsesDuration(responses))
	util.WriteResponsesDuration(fmt.Sprintf(filePath, time.Now().Format(constant.TimeLayout)), responses)
}

func bringUpUDPServer() {
	log.Println("Bring up server over UDP")
	udpResolver, _ := net.ResolveUDPAddr(constant.UDP, constant.PORT)
	listener, _ := net.ListenUDP(constant.UDP, udpResolver)
	defer listener.Close()

	log.Println(constant.Address, listener.LocalAddr())

	for {
		var buffer [2048]byte
		var request communication.Request

		cutPoint, requestAddress, _ := listener.ReadFromUDP(buffer[0:])
		json.Unmarshal(buffer[:cutPoint], &request)
		log.Println(fmt.Sprintf(constant.ReceivingRequest, constant.UDP), request)

		tokens := nlg.TextTokenize(request.Content)

		response, _ := json.Marshal(communication.Response{Content: tokens})
		listener.WriteToUDP(response, requestAddress)
		log.Println(fmt.Sprintf(constant.SendingResponse, constant.UDP), tokens)
	}
}
