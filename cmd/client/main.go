package main

import (
	"github.com/marlom-jobsom/goWordTokenize/internal"
	"github.com/marlom-jobsom/goWordTokenize/internal/layers/distribution/client"
)

func main() {
	text, protocol, rpc := internal.GetClientArgs()
	clientProxy := new(client.Proxy)
	clientProxy.InvokeTextTokenize(protocol, rpc, text)
}
