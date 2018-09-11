package main

import (
	"fmt"
	"log"
	"word-tokenize-in1118/constant"
	"word-tokenize-in1118/core"
	"word-tokenize-in1118/core/communication"
)

func main() {
	text, protocol, rpc, test := core.GetCliArgs()
	client := new(communication.TokenizerClient)
	log.Printf("Test: %t", test)

	if rpc {
		switch protocol {
		case constant.TCP:
			client.TextTokenizeRPCTCP(text)
		case constant.UDP:
			client.TextTokenizeRPCUDP(text)
		default:
			log.Fatal(fmt.Errorf(constant.ErrorPattern, constant.TCP, constant.UDP))
		}
	} else {
		switch protocol {
		case constant.TCP:
			client.TextTokenizeTCP(text)
		case constant.UDP:
			client.TextTokenizeUDP(text)
		default:
			log.Fatal(fmt.Errorf(constant.ErrorPattern, constant.TCP, constant.UDP))
		}
	}
}
