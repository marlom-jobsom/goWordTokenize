package main

import (
	"encoding/json"
	"log"
	"net"
	"word-tokenize-middleware-socket/core"
	"word-tokenize-middleware-socket/util"
)

func buildListener() net.Listener {
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
	listen := buildListener()

	for {
		connection, _ := listen.Accept()
		jsonEncoder := json.NewEncoder(connection)
		jsonDecoder := json.NewDecoder(connection)

		// Close the connection when the application closes.
		defer connection.Close()

		request := handleRequest(connection, jsonDecoder)
		buildResponse(request, jsonEncoder)
	}
}
