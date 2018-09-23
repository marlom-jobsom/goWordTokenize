package util

import (
	"bufio"
	"net"
	"net/rpc"
	"os"
	"word-tokenize-in1118/internal/constant"
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
	connection, _ := net.Dial(constant.TCP, constant.PORT)
	return connection
}

// DialUDPConnection dials to a UDP connection
func DialUDPConnection() net.Conn {
	connection, _ := net.Dial(constant.UDP, constant.PORT)
	return connection
}

// DialRPCTCPClient dials to remote procedure call over TCP
func DialRPCTCPClient() *rpc.Client {
	rpcClient, _ := rpc.Dial(constant.TCP, constant.PORT)
	return rpcClient
}

// BuildTCPListener builds a TCP listener
func BuildTCPListener() net.Listener {
	tcpAddress, _ := net.ResolveTCPAddr(constant.TCP, constant.PORT)
	listen, _ := net.ListenTCP(constant.TCP, tcpAddress)
	return listen
}

// BuildUDPListener builds a UDP listener
func BuildUDPListener() *net.UDPConn {
	udpResolver, _ := net.ResolveUDPAddr(constant.UDP, constant.PORT)
	connection, _ := net.ListenUDP(constant.UDP, udpResolver)
	return connection
}
