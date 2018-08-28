package main

import (
	"encoding/json"
	"flag"
	"log"
	"net"
	"word-tokenize-middleware-socket/core"
)

// getCliArgs ... Get argument from CLI
func getCliArgs() (string, string) {
	var text string
	var protocol string

	flag.StringVar(&text, "text", "", "Text to be tokenize")
	flag.StringVar(&protocol, "protocol", "", "The client network prototol (tcp or udp)")
	flag.Parse()

	return text, protocol
}

func main() {
	text, protocol := getCliArgs()

	connection, _ := net.Dial(protocol, "localhost:5000")
	jsonEncoder := json.NewEncoder(connection)
	jsonDecoder := json.NewDecoder(connection)
	defer connection.Close()

	request := core.Request{Content: text}
	jsonEncoder.Encode(request)
	log.Println("Send:", request)

	var response core.Response
	jsonDecoder.Decode(&response)
	log.Println("Receive:", response)
}
