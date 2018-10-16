package internal

import (
	"flag"
	"github.com/marlom-jobsom/goWordTokenize/internal/constant"
)

// GetClientArgs reads the parameters from CLI interface for main client_request_handler
func GetClientArgs() (string, string, bool) {
	var text string
	var protocol string
	var rpc bool

	flag.StringVar(&text, "text", "Silent is is golden", "Text to be tokenize")
	flag.StringVar(&protocol, "protocol", constant.TCP, "The client_request_handler network prototol (tcp or udp)")
	flag.BoolVar(&rpc, "rpc", false, "Uses RPC client_request_handler (over TCP only)")
	flag.Parse()

	return text, protocol, rpc
}

// GetServerArgs reads the parameters from CLI interface for main server_request_handler
func GetServerArgs() (string, bool) {
	var protocol string
	var rpc bool

	flag.StringVar(&protocol, "protocol", constant.TCP, "The server_request_handler network prototol (tcp or udp)")
	flag.BoolVar(&rpc, "rpc", false, "Uses RPC server_request_handler (over TCP only)")
	flag.Parse()

	return protocol, rpc
}
