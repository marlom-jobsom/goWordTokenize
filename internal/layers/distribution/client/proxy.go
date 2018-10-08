package client

import (
	"fmt"
	"log"
	"word-tokenize-in1118/internal/communication"
	"word-tokenize-in1118/internal/constant"
)

// Proxy route the invoke through the requestor
type Proxy struct{}

// InvokeTextTokenize invokes client_request_handler over the protocol and text given
func (proxy *Proxy) InvokeTextTokenize(protocol string, rpc bool, text string) communication.Response {
	var response communication.Response

	if rpc {
		switch protocol {
		case constant.TCP:
			response = proxy.invokeTextTokenizeRPCTCP(text)
		default:
			// RPC over UDP is not supported by Golang
			// https://astaxie.gitbooks.io/build-web-application-with-golang/en/08.4.html
			// https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/rpc/go_rpc.html
			log.Fatal(fmt.Errorf(constant.RPCUDPPatternError, constant.TCP))
		}
	} else {
		switch protocol {
		case constant.TCP:
			response = proxy.invokeTextTokenizeTCP(text)
		case constant.UDP:
			response = proxy.invokeTextTokenizeUDP(text)
		default:
			log.Fatal(fmt.Errorf(constant.ProtocolPatternError, constant.TCP, constant.UDP))
		}
	}

	return response
}

// invokeTextTokenizeRPCTCP invoke client_request_handler RPC over TCP
func (proxy *Proxy) invokeTextTokenizeRPCTCP(text string) communication.Response {
	req := new(Requestor)
	return req.RequestTextTokenizeRPCTCP(text)
}

// invokeTextTokenizeTCP invoke client_request_handler over TCP
func (proxy *Proxy) invokeTextTokenizeTCP(text string) communication.Response {
	req := new(Requestor)
	return req.RequestTextTokenizeTCP(text)
}

// invokeTextTokenizeUDP invoke client_request_handler over TCP
func (proxy *Proxy) invokeTextTokenizeUDP(text string) communication.Response {
	req := new(Requestor)
	return req.RequestTextTokenizeUDP(text)
}
