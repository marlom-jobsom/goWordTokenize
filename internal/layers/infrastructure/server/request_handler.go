package server

import (
	"encoding/json"
	"log"
	"net/rpc"
	"word-tokenize-in1118/internal/constant"
	"word-tokenize-in1118/internal/layers/distribution/server"
	"word-tokenize-in1118/internal/services/nlg"
	"word-tokenize-in1118/internal/util"
)

// RequestHandler ...
type RequestHandler struct{}

// BringUpRPCTCPServer runs a RCP server over TCP
func (RequestHandler) BringUpRPCTCPServer() {
	log.Println("Bring up RPC server over TCP")
	listen := util.BuildTCPListener()

	rpcServer := rpc.NewServer()
	rpcServer.RegisterName(constant.NLG, new(nlg.NLG))

	log.Println(constant.Address, listen.Addr())
	rpcServer.Accept(listen)
}

// BringUpTCPServer runs a server over TCP
func (RequestHandler) BringUpTCPServer() {
	log.Println("Bring up server over TCP")
	listener := util.BuildTCPListener()
	defer listener.Close()
	log.Println(constant.Address, listener.Addr())
	ivk := new(server.Invoker)

	for {
		connection, _ := listener.Accept()
		jsonEncoder := json.NewEncoder(connection)
		jsonDecoder := json.NewDecoder(connection)
		ivk.InvokeTextTokenizeTCP(jsonEncoder, jsonDecoder)
		connection.Close()
	}
}

// BringUpUDPServer runs a server over UDP
func (RequestHandler) BringUpUDPServer() {
	log.Println("Bring up server over UDP")
	listener := util.BuildUDPListener()
	defer listener.Close()
	log.Println(constant.Address, listener.LocalAddr())
	ivk := new(server.Invoker)

	for {
		ivk.InvokeTextTokenizeUDP(listener)
	}
}
