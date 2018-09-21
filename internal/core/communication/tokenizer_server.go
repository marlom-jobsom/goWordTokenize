package communication

import (
	"encoding/json"
	"log"
	"net"
	"net/rpc"
	"word-tokenize-in1118/internal/constant"
	"word-tokenize-in1118/internal/core/nlg"
)

// TokenizerServer ...
type TokenizerServer struct{}

// BringUpRPCTCPServer runs a RCP server over TCP
func (TokenizerServer) BringUpRPCTCPServer() {
	log.Println("Bring up RPC over TCP")
	listen := buildTCPListener()

	rpcServer := rpc.NewServer()
	rpcServer.RegisterName(constant.NLG, new(nlg.NLG))

	log.Println("Address", listen.Addr())
	rpcServer.Accept(listen)
}

// BringUpTCPServer runs a server over TCP
func (TokenizerServer) BringUpTCPServer() {
	log.Println("Bring up server over TCP")
	listener := buildTCPListener()
	defer listener.Close()
	log.Println("Address", listener.Addr())

	for {
		connection, _ := listener.Accept()
		jsonEncoder := json.NewEncoder(connection)
		jsonDecoder := json.NewDecoder(connection)

		request := handleTCPRequest(connection, jsonDecoder)
		buildTCPResponse(request, jsonEncoder)

		connection.Close()
	}
}

// BringUpUDPServer runs a server over UDP
func (TokenizerServer) BringUpUDPServer() {
	log.Println("Bring up server over UDP")
	listener := buildUDPListener()
	defer listener.Close()
	log.Println("Address", listener.LocalAddr())

	for {
		request, requestAddress := handleUDPRequest(listener)
		buildUDPResponse(listener, request, requestAddress)
	}
}

// Helper: builds a TCP listener
func buildTCPListener() net.Listener {
	tcpAddress, _ := net.ResolveTCPAddr(constant.TCP, constant.PORT)
	listen, _ := net.ListenTCP(constant.TCP, tcpAddress)
	return listen
}

// Helper: handles incoming TCP requests
func handleTCPRequest(connection net.Conn, jsonDecoder *json.Decoder) Request {
	var request Request
	jsonDecoder.Decode(&request)
	log.Println("Receive:", request)
	return request
}

// Helper: build outgoing TCP responses
func buildTCPResponse(request Request, jsonEncoder *json.Encoder) {
	tokens := nlg.TextTokenize(request.Content)
	log.Println("Send tokens:", tokens)
	jsonEncoder.Encode(Response{Content: tokens})
}

// Helper: builds a UDP listener
func buildUDPListener() *net.UDPConn {
	udpResolver, _ := net.ResolveUDPAddr(constant.UDP, constant.PORT)
	connection, _ := net.ListenUDP(constant.UDP, udpResolver)
	return connection
}

// Helper: handles incoming UDP requests
func handleUDPRequest(connection *net.UDPConn) (Request, *net.UDPAddr) {
	var buffer [2048]byte
	var request Request

	cutPoint, requestAddress, _ := connection.ReadFromUDP(buffer[0:])
	json.Unmarshal(buffer[:cutPoint], &request)
	log.Println("Request:", request)

	return request, requestAddress
}

// Helper: build outcomming UDP responses
func buildUDPResponse(connection *net.UDPConn, request Request, requestAddress *net.UDPAddr) {
	tokens := nlg.TextTokenize(request.Content)
	response, _ := json.Marshal(Response{Content: tokens})
	connection.WriteToUDP(response, requestAddress)
	log.Println("Response:", tokens)
}
