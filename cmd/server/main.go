package main

import (
	"fmt"
	"log"
	"word-tokenize-in1118/internal/constant"
	"word-tokenize-in1118/internal/core"
	"word-tokenize-in1118/internal/core/communication"
)

func main() {
	protocol, rpc := core.GetServerArgs()
	server := new(communication.TokenizerServer)

	if rpc {
		switch protocol {
		case constant.TCP:
			server.BringUpRPCTCPServer()
		default:
			// RPC over UDP is not supported by Golang
			// https://astaxie.gitbooks.io/build-web-application-with-golang/en/08.4.html
			// https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/rpc/go_rpc.html
			log.Fatal(fmt.Errorf(constant.RPCUDPPatternError, constant.TCP))
		}
	} else {
		switch protocol {
		case constant.TCP:
			server.BringUpTCPServer()
		case constant.UDP:
			server.BringUpUDPServer()
		default:
			log.Fatal(fmt.Errorf(constant.ProtolPatternError, constant.TCP, constant.UDP))
		}
	}
}
