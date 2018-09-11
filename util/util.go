package util

import (
	"bufio"
	"net"
	"net/rpc"
	"os"
	"word-tokenize-socket/constant"
)

// CreateFileIfDoestNotExists creates a file for the path given if it doesn't exists
func CreateFileIfDoestNotExists(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Create(filePath)
	}
}

// AppendContentFile appends content to a existing file
func AppendContentFile(filePath string, content string) {
	CreateFileIfDoestNotExists(filePath)

	file, _ := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(content + "\n")
	writer.Flush()
}

// DialTCPConnection dials to a TCP connection
func DialTCPConnection() net.Conn {
	return buildConnection(constant.TCP)
}

// DialUDPConnection dials to a UDP connection
func DialUDPConnection() net.Conn {
	return buildConnection(constant.UDP)
}

// DialRPCTCPClient dials to remote procedure call over TCP
func DialRPCTCPClient() *rpc.Client {
	rpcClient, _ := rpc.Dial(constant.TCP, constant.PORT)
	return rpcClient
}

// DialRPCUDPClient dials to remote procedure call over UDP
func DialRPCUDPClient() *rpc.Client {
	rpcClient, _ := rpc.Dial(constant.UDP, constant.PORT)
	return rpcClient
}

// Helper: dials to a connection under the protocol given
func buildConnection(protocol string) net.Conn {
	connection, _ := net.Dial(protocol, constant.PORT)
	return connection
}
