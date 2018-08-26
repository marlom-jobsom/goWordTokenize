package main

import (
	"encoding/json"
	"log"
	"net"
	"word-tokenize-middleware-socket/core"
	"word-tokenize-middleware-socket/util"
)

func main() {
	connection, _ := net.Dial("udp", "localhost:5000")
	jsonEncoder := json.NewEncoder(connection)
	jsonDecoder := json.NewDecoder(connection)
	defer connection.Close()

	text := util.GetCliArgs()
	request := core.Request{Content: text}
	jsonEncoder.Encode(request)
	log.Println("Send:", request)

	var response core.Response
	jsonDecoder.Decode(&response)
	log.Println("Receive:", response)
}
