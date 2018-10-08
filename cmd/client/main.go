package main

import (
	"word-tokenize-in1118/internal"
	"word-tokenize-in1118/internal/layers/distribution/client"
)

func main() {
	text, protocol, rpc := internal.GetClientArgs()
	clientProxy := new(client.Proxy)
	clientProxy.InvokeTextTokenize(protocol, rpc, text)
}
