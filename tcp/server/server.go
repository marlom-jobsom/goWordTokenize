package main

import (
	"encoding/json"
	"log"
	"net"
	"word-tokenize-middleware-socket/core"
	"word-tokenize-middleware-socket/util"
)

func main() {
	log.Println("Starting TCP server")
	listen, _ := net.Listen("tcp", ":5000")

	for {
		connection, _ := listen.Accept()
		jsonEncoder := json.NewEncoder(connection)
		jsonDecoder := json.NewDecoder(connection)
		defer connection.Close()

		var request core.Request
		jsonDecoder.Decode(&request)
		log.Println("Receive:", request)

		tokens := util.TextTokenize(request)
		log.Println("Send tokens:", tokens)
		jsonEncoder.Encode(core.Response{Content: tokens})
	}
}
