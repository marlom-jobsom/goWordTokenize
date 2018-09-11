package communication

import (
	"encoding/json"
	"log"
	"net"
	"net/rpc"
	"word-tokenize-in1118/constant"
	"word-tokenize-in1118/core/nlg"
)

// TokenizerServer ...
type TokenizerServer struct{}

// BringUpRPCTCPServer runs a RCP server over TCP
func (TokenizerServer) BringUpRPCTCPServer() {
	log.Println("Bring up RPC over TCP")
	tcpAddress, _ := net.ResolveTCPAddr(constant.TCP, constant.PORT)
	listen, _ := net.ListenTCP(constant.TCP, tcpAddress)

	rpcServer := rpc.NewServer()
	rpcServer.RegisterName(constant.NLG, new(nlg.NLG))

	log.Println("Address", listen.Addr())
	rpcServer.Accept(listen)
}

// BringUpTCPServer runs a server over TCP
func (TokenizerServer) BringUpTCPServer() {
	listener := buildTCPListener()
	defer listener.Close()

	for {
		connection, _ := listener.Accept()
		jsonEncoder := json.NewEncoder(connection)
		jsonDecoder := json.NewDecoder(connection)

		request := handleTCPRequest(connection, jsonDecoder)
		buildTCPResponse(request, jsonEncoder)

		connection.Close()
	}
}

// Helper: builds a TCP listener
func buildTCPListener() net.Listener {
	listen, _ := net.Listen(constant.TCP, constant.PORT)
	log.Println("TCP server address:", listen.Addr())
	return listen
}

// Helper: handles incoming TCP requests
func handleTCPRequest(connection net.Conn, jsonDecoder *json.Decoder) Request {
	var request Request
	jsonDecoder.Decode(&request)
	log.Println("Receive:", request)
	return request
}

// Helper: build outcomming TCP responses
func buildTCPResponse(request Request, jsonEncoder *json.Encoder) {
	tokens := nlg.TextTokenize(request.Content)
	log.Println("Send tokens:", tokens)
	jsonEncoder.Encode(Response{Content: tokens})
}
