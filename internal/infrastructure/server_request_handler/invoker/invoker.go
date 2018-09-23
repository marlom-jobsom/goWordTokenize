package invoker

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"word-tokenize-in1118/internal/communication"
	"word-tokenize-in1118/internal/constant"
	"word-tokenize-in1118/internal/services/nlg"
)

// Invoker ...
type Invoker struct{}

// InvokeTextTokenizeTCP invokes text tokenize for TCP request
func (invoker *Invoker) InvokeTextTokenizeTCP(encoder *json.Encoder, decoder *json.Decoder) {
	request := receiveTCPRequest(decoder)
	sendTCPResponse(request, encoder)
}

// InvokeTextTokenizeUDP invokes text tokenize for UDP request
func (invoker *Invoker) InvokeTextTokenizeUDP(connection *net.UDPConn) {
	request, requestAddress := receiveUDPRequest(connection)
	sendUDPResponse(connection, request, requestAddress)
}

// Helper: receives incoming TCP requests
func receiveTCPRequest(jsonDecoder *json.Decoder) communication.Request {
	var request communication.Request
	jsonDecoder.Decode(&request)
	log.Println(fmt.Sprintf(constant.ReceivingRequest, constant.TCP), request)
	return request
}

// Helper: sends outgoing TCP responses
func sendTCPResponse(request communication.Request, jsonEncoder *json.Encoder) {
	tokens := nlg.TextTokenize(request.Content)
	log.Println(fmt.Sprintf(constant.SendingResponse, constant.TCP), tokens)
	jsonEncoder.Encode(communication.Response{Content: tokens})
}

// Helper: receives incoming UDP requests
func receiveUDPRequest(connection *net.UDPConn) (communication.Request, *net.UDPAddr) {
	var buffer [2048]byte
	var request communication.Request

	cutPoint, requestAddress, _ := connection.ReadFromUDP(buffer[0:])
	json.Unmarshal(buffer[:cutPoint], &request)
	log.Println(fmt.Sprintf(constant.ReceivingRequest, constant.UDP), request)

	return request, requestAddress
}

// Helper: sends outgoing UDP responses
func sendUDPResponse(connection *net.UDPConn, request communication.Request, requestAddress *net.UDPAddr) {
	tokens := nlg.TextTokenize(request.Content)
	response, _ := json.Marshal(communication.Response{Content: tokens})
	connection.WriteToUDP(response, requestAddress)
	log.Println(fmt.Sprintf(constant.SendingResponse, constant.UDP), tokens)
}
