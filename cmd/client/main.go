package main

import (
	"fmt"
	"log"
	"word-tokenize-in1118/internal/constant"
	"word-tokenize-in1118/internal/core"
	"word-tokenize-in1118/internal/core/communication"
)

func main() {
	text, protocol, rpc, test := core.GetClientArgs()
	client := new(communication.TokenizerClient)
	log.Printf("Test: %t", test)

	if rpc {
		switch protocol {
		case constant.TCP:
			client.TextTokenizeRPCTCP(text)
		default:
			// RPC over UDP is not supported by Golang
			// https://astaxie.gitbooks.io/build-web-application-with-golang/en/08.4.html
			// https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/rpc/go_rpc.html
			log.Fatal(fmt.Errorf(constant.RPCUDPPatternError, constant.TCP))
		}
	} else {
		switch protocol {
		case constant.TCP:
			client.TextTokenizeTCP(text)
		case constant.UDP:
			client.TextTokenizeUDP(text)
		default:
			log.Fatal(fmt.Errorf(constant.ProtolPatternError, constant.TCP, constant.UDP))
		}
	}
}
