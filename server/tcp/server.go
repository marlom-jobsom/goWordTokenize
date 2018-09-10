package main

import (
	"encoding/json"
	"log"
	"net"
	"word-tokenize-socket/core"
	"word-tokenize-socket/util"
)

func buildTCPListener() net.Listener {
	listen, _ := net.Listen("tcp", ":5000")
	log.Println("TCP server address:", listen.Addr())
	return listen
}

func handleRequest(connection net.Conn, jsonDecoder *json.Decoder) core.Request {
	var request core.Request
	jsonDecoder.Decode(&request)
	log.Println("Receive:", request)
	return request
}

func buildResponse(request core.Request, jsonEncoder *json.Encoder) {
	tokens := util.TextTokenize(request)
	log.Println("Send tokens:", tokens)
	jsonEncoder.Encode(core.Response{Content: tokens})
}

func main() {
	listener := buildTCPListener()

	// Close the listener when the application closes.
	defer listener.Close()

	for {
		connection, _ := listener.Accept()
		jsonEncoder := json.NewEncoder(connection)
		jsonDecoder := json.NewDecoder(connection)

		request := handleRequest(connection, jsonDecoder)
		buildResponse(request, jsonEncoder)

		// Close the connection when finish to handle it
		connection.Close()
	}
}
