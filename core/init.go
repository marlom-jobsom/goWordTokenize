package core

import (
	"flag"
	"word-tokenize-in1118/constant"
)

// GetCliArgs reads the parameters from CLI interface
func GetCliArgs() (string, string, bool, bool) {
	var text string
	var protocol string
	var rpc bool
	var test bool

	flag.StringVar(&text, "text", "Silent is is golden", "Text to be tokenize")
	flag.StringVar(&protocol, "protocol", constant.TCP, "The client network prototol (tcp or udp)")
	flag.BoolVar(&rpc, "rpc", false, "Uses RPC client")
	flag.BoolVar(&test, "test", false, "Enable client test mode")
	flag.Parse()

	return text, protocol, rpc, test
}
