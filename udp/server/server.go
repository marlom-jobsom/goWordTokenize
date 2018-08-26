package main

import (
	"encoding/json"
	"log"
	"net"
	"word-tokenize-middleware-socket/core"
	"word-tokenize-middleware-socket/util"
)

func buildServer() *net.UDPConn {
	udpTag := "udp"
	udpResolver, _ := net.ResolveUDPAddr(udpTag, ":5000")
	connection, _ := net.ListenUDP(udpTag, udpResolver)
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
	log.Println("Starting UDP server")
	connection := buildServer()

	for {
		request, requestAddress := handleRequest(connection)
		buildResponse(connection, request, requestAddress)
	}
}
