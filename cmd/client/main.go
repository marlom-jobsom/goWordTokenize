package main

import (
	"word-tokenize-in1118/internal"
	"word-tokenize-in1118/internal/infrastructure/client_request_handler/proxy"
)

func main() {
	text, protocol, rpc := internal.GetClientArgs()
	clientProxy := new(proxy.ClientProxy)
	clientProxy.InvokeTextTokenize(protocol, rpc, text)
}
