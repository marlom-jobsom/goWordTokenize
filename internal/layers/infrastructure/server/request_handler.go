package server

import (
	"encoding/json"
	"github.com/marlom-jobsom/goWordTokenize/internal/constant"
	"github.com/marlom-jobsom/goWordTokenize/internal/layers/distribution/server"
	"github.com/marlom-jobsom/goWordTokenize/internal/services/nlg"
	"github.com/marlom-jobsom/goWordTokenize/internal/util"
	"log"
	"net/rpc"
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
