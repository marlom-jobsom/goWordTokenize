package communication

import (
	"log"
	"net"
	"net/rpc"
	"word-tokenize-socket/constant"
	"word-tokenize-socket/core/nlg"
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
