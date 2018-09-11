package core

import (
	"flag"
	"word-tokenize-in1118/constant"
)

// GetClientArgs reads the parameters from CLI interface for main client
func GetClientArgs() (string, string, bool, bool) {
	var text string
	var protocol string
	var rpc bool
	var test bool

	flag.StringVar(&text, "text", "Silent is is golden", "Text to be tokenize")
	flag.StringVar(&protocol, "protocol", constant.TCP, "The client network prototol (tcp or udp)")
	flag.BoolVar(&rpc, "rpc", false, "Uses RPC client (over TCP only)")
	flag.BoolVar(&test, "test", false, "Enable client test mode")
	flag.Parse()

	return text, protocol, rpc, test
}

// GetServerArgs reads the parameters from CLI interface for main server
func GetServerArgs() (string, bool) {
	var protocol string
	var rpc bool

	flag.StringVar(&protocol, "protocol", constant.TCP, "The server network prototol (tcp or udp)")
	flag.BoolVar(&rpc, "rpc", false, "Uses RPC server (over TCP only)")
	flag.Parse()

	return protocol, rpc
}
