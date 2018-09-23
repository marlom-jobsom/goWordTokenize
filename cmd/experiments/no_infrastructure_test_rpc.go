package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
	"word-tokenize-in1118/cmd/experiments/util"
	"word-tokenize-in1118/internal/communication"
	"word-tokenize-in1118/internal/constant"
	"word-tokenize-in1118/internal/services/nlg"
)

// TestRPC tests the RPC communication timing for the client_request_handler side
func main() {
	filePath := "no_infrastructure_test_rpc_%s.txt"
	var responses []communication.Response

	go bringUpRPCTCPServer()
	time.Sleep(constant.SecondsSleep)

	for i := 0; i < constant.MaxTries; i++ {
		client, _ := rpc.Dial(constant.TCP, constant.PORT)

		log.Println(fmt.Sprintf(constant.SendingRequest, constant.RPC), constant.TextTest)
		var response communication.Response
		var tokens []string

		now := time.Now()
		client.Call(constant.NLGTextTokenizeRPC, constant.TextTest, &tokens)
		elapsed := time.Since(now)

		response.Duration = elapsed
		response.Content = tokens
		client.Close()

		log.Println(fmt.Sprintf(constant.ReceivingResponse, constant.RPC), response)

		responses = append(responses, response)
	}

	log.Println(util.SumResponsesDuration(responses))
	util.WriteResponsesDuration(fmt.Sprintf(filePath, time.Now().Format(constant.TimeLayout)), responses)
}

func bringUpRPCTCPServer() {
	log.Println("Bring up RPC server over TCP")

	tcpAddress, _ := net.ResolveTCPAddr(constant.TCP, constant.PORT)
	listen, _ := net.ListenTCP(constant.TCP, tcpAddress)

	rpcServer := rpc.NewServer()
	rpcServer.RegisterName(constant.NLG, new(nlg.NLG))

	log.Println(constant.Address, listen.Addr())
	rpcServer.Accept(listen)
}
