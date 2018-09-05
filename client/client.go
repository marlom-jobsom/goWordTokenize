package main

import (
	"encoding/json"
	"flag"
	"log"
	"net"
	"time"
	"word-tokenize-middleware-socket/core"
	"word-tokenize-middleware-socket/util"
)

func getCliArgs() (string, string, bool) {
	var text string
	var protocol string
	var test bool

	flag.StringVar(&text, "text", "", "Text to be tokenize")
	flag.StringVar(&protocol, "protocol", "", "The client network prototol (tcp or udp)")
	flag.BoolVar(&test, "test", false, "Enable client test mode")
	flag.Parse()

	return text, protocol, test
}

func buildConnection(protocol string) (net.Conn, *json.Encoder, *json.Decoder) {
	connection, _ := net.Dial(protocol, ":5000")
	jsonEncoder := json.NewEncoder(connection)
	jsonDecoder := json.NewDecoder(connection)
	return connection, jsonEncoder, jsonDecoder
}

func buildRequest(text string, jsonEncoder *json.Encoder) {
	request := core.Request{Content: text}
	jsonEncoder.Encode(request)
	log.Println("Send:", request)
}

func handleResponse(jsonDecoder *json.Decoder) {
	var response core.Response
	jsonDecoder.Decode(&response)
	log.Println("Receive:", response)
}

func testExecution(protocol string) {
	requestNum := 50
	text := "is Silence is golden"
	defer util.TimeTrack(time.Now(), "testExecution", protocol)

	for i := 0; i < requestNum; i++ {
		connection, jsonEncoder, jsonDecoder := buildConnection(protocol)
		test(text, jsonEncoder, jsonDecoder, protocol)

		// Close the connection when finish to handle it
		connection.Close()
	}
}

func test(text string, jsonEncoder *json.Encoder, jsonDecoder *json.Decoder, protocol string) {
	defer util.TimeTrack(time.Now(), "test", protocol)
	buildRequest(text, jsonEncoder)
	handleResponse(jsonDecoder)
}

func main() {
	text, protocol, test := getCliArgs()

	if test {
		testExecution(protocol)
	} else {
		connection, jsonEncoder, jsonDecoder := buildConnection(protocol)

		// Close the connection when the application closes.
		defer connection.Close()

		buildRequest(text, jsonEncoder)
		handleResponse(jsonDecoder)
	}
}
