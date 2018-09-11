package communication

import (
	"encoding/json"
	"log"
	"net"
	"net/rpc"
	"time"
	"word-tokenize-socket/constant"
	"word-tokenize-socket/util"
)

// TokenizerClient ...
type TokenizerClient struct{}

// TextTokenizeTCP performs a TCP request to text tokenize
func (tc *TokenizerClient) TextTokenizeTCP(text string) []string {
	connection := util.DialTCPConnection()
	defer connection.Close()
	log.Println("TextTokenizeTCP: Request", text)
	response := tokenizeRequest(text, connection)
	log.Println("TextTokenizeTCP: Response", text)
	return response.Content
}

// TextTokenizeRPCTCP performs a remote procedure call over TCP to text tokenize
func (tc *TokenizerClient) TextTokenizeRPCTCP(text string) []string {
	client := util.DialRPCTCPClient()
	log.Println("TextTokenizeRPCTCP: Request", text)
	response := tokenizeRequestRPC(text, client)
	log.Println("TextTokenizeRPCTCP: Response", response)
	return response.Content
}

// TextTokenizeUDP performs a UDP request to text tokenize
func (tc *TokenizerClient) TextTokenizeUDP(text string) []string {
	connection := util.DialUDPConnection()
	defer connection.Close()
	log.Println("TextTokenizeUDP: Request", text)
	response := tokenizeRequest(text, connection)
	log.Println("TextTokenizeUDP: Response", text)
	return response.Content
}

// TextTokenizeRPCUDP performs a remote procedure call over UDP to text tokenize
func (tc *TokenizerClient) TextTokenizeRPCUDP(text string) []string {
	var tokens []string
	client := util.DialRPCUDPClient()
	log.Println("TextTokenizeRPCUDP: Request", text)
	client.Call(constant.NLGTextTokenizeRPC, text, &tokens)
	log.Println("TextTokenizeRPCUDP: Response", tokens)
	return tokens
}

// Helper: trigger the text tokenize RPC request
func tokenizeRequestRPC(text string, client *rpc.Client) Response {
	var response Response
	var tokens []string
	now := time.Now()
	client.Call(constant.NLGTextTokenizeRPC, text, &tokens)
	elapsed := time.Since(now)
	response.Duration = elapsed
	response.Content = tokens
	return response
}

// Helper: trigger the text tokenize request
func tokenizeRequest(text string, connection net.Conn) Response {
	var response Response
	now := time.Now()
	encodeRequest(text, json.NewEncoder(connection))
	response = decodeResponse(json.NewDecoder(connection))
	elapsed := time.Since(now)
	response.Duration = elapsed
	return response
}

// Helper: Encodes a request
func encodeRequest(text string, jsonEncoder *json.Encoder) {
	request := Request{Content: text}
	jsonEncoder.Encode(request)
}

// Helper: Decodes a response
func decodeResponse(jsonDecoder *json.Decoder) Response {
	var response Response
	jsonDecoder.Decode(&response)
	return response
}
