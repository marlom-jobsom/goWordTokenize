package server

import (
	"encoding/json"
	"log"
	"net"
	"word-tokenize-in1118/core"
	"word-tokenize-in1118/util"
)

func buildUDPListener() *net.UDPConn {
	udpResolver, _ := net.ResolveUDPAddr("udp", ":5000")
	connection, _ := net.ListenUDP("udp", udpResolver)

	log.Println("UDP server address:", connection.LocalAddr())
	return connection
}

func handleRequest(connection *net.UDPConn) (core.Request, *net.UDPAddr) {
	var buffer [2048]byte
	var request core.Request

	cutPoint, requestAddress, _ := connection.ReadFromUDP(buffer[0:])
	json.Unmarshal(buffer[:cutPoint], &request)
	log.Println("Request:", request)

	return request, requestAddress
}

func buildResponse(connection *net.UDPConn, request core.Request, requestAddress *net.UDPAddr) {
	tokens := util.TextTokenize(request)
	response, _ := json.Marshal(core.Response{Content: tokens})
	connection.WriteToUDP(response, requestAddress)
	log.Println("Response:", tokens)
}

func main() {
	listener := buildUDPListener()

	// Close the listener when the application closes.
	defer listener.Close()

	for {
		request, requestAddress := handleRequest(listener)
		buildResponse(listener, request, requestAddress)
	}
}
