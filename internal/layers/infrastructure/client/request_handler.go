package client

import (
	"encoding/json"
	"fmt"
	"github.com/marlom-jobsom/goWordTokenize/internal/communication"
	"github.com/marlom-jobsom/goWordTokenize/internal/constant"
	"github.com/marlom-jobsom/goWordTokenize/internal/util"
	"log"
	"net"
	"net/rpc"
	"time"
)

// RequestHandler ...
type RequestHandler struct{}

// TextTokenizeRPCTCP handles a remote procedure call over TCP to text tokenize
func (requestHandler *RequestHandler) TextTokenizeRPCTCP(text string) communication.Response {
	client := util.DialRPCTCPClient()
	defer client.Close()
	log.Println(fmt.Sprintf(constant.SendingRequest, constant.RPC), text)
	response := sendRPC(text, client)
	log.Println(fmt.Sprintf(constant.ReceivingResponse, constant.RPC), response)
	return response
}

// TextTokenizeTCP handles a TCP request to text tokenize
func (requestHandler *RequestHandler) TextTokenizeTCP(request communication.Request) communication.Response {
	connection := util.DialTCPConnection()
	defer connection.Close()
	log.Println(fmt.Sprintf(constant.SendingRequest, constant.TCP), request)
	response := send(request, connection)
	log.Println(fmt.Sprintf(constant.ReceivingResponse, constant.TCP), response)
	return response
}

// TextTokenizeUDP handles a UDP request to text tokenize
func (requestHandler *RequestHandler) TextTokenizeUDP(request communication.Request) communication.Response {
	connection := util.DialUDPConnection()
	defer connection.Close()
	log.Println(fmt.Sprintf(constant.SendingRequest, constant.UDP), request)
	response := send(request, connection)
	log.Println(fmt.Sprintf(constant.ReceivingResponse, constant.UDP), response)
	return response
}

// Helper: sends the text tokenize RPC request
func sendRPC(text string, client *rpc.Client) communication.Response {
	var response communication.Response
	var tokens []string
	now := time.Now()
	client.Call(constant.NLGTextTokenizeRPC, text, &tokens)
	elapsed := time.Since(now)
	response.Duration = elapsed
	response.Content = tokens
	return response
}

// Helper: sends a request
func send(request communication.Request, connection net.Conn) communication.Response {
	var response communication.Response
	now := time.Now()
	json.NewEncoder(connection).Encode(request)
	response = receive(connection)
	elapsed := time.Since(now)
	response.Duration = elapsed
	return response
}

// Helper: receives a response
func receive(connection net.Conn) communication.Response {
	var response communication.Response
	json.NewDecoder(connection).Decode(&response)
	return response
}
