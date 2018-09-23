package requestor

import (
	"word-tokenize-in1118/internal/communication"
	"word-tokenize-in1118/internal/infrastructure/client_request_handler/requesthandler"
)

// Requestor route handles a client_request_handler request
type Requestor struct{}

// RequestTextTokenizeRPCTCP requests a remote procedure call request over TCP to text tokenize
func (tc *Requestor) RequestTextTokenizeRPCTCP(text string) communication.Response {
	requestHandler := new(requesthandler.RequestHandler)
	return requestHandler.TextTokenizeRPCTCP(text)
}

// RequestTextTokenizeTCP requests a TCP request to text tokenize
func (tc *Requestor) RequestTextTokenizeTCP(text string) communication.Response {
	request := communication.Request{Content: text}
	requestHandler := new(requesthandler.RequestHandler)
	return requestHandler.TextTokenizeTCP(request)
}

// RequestTextTokenizeUDP requests a UDP request to text tokenize
func (tc *Requestor) RequestTextTokenizeUDP(text string) communication.Response {
	request := communication.Request{Content: text}
	requestHandler := new(requesthandler.RequestHandler)
	return requestHandler.TextTokenizeUDP(request)
}
