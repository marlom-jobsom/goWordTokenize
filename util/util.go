package util

import (
	"flag"
	"net"
	"strings"
	"word-tokenize-middleware-socket/core"
)

// GetCliArgs ... Get argument from CLI
func GetCliArgs() string {
	var text string

	flag.StringVar(&text, "text", "", "Text to be tokenize")
	flag.Parse()

	return text
}

// ReceiveText ... Receive the text content from connection
func ReceiveText(connection net.Conn) string {
	buffer := make([]byte, 1024)
	cutPoint, _ := connection.Read(buffer)
	text := string(buffer[:cutPoint])
	return text
}

// TextTokenize ... Tokenize a text content
func TextTokenize(request core.Request) []string {
	var tokens []string
	mapTokens := make(map[string]struct{})
	words := strings.Split(request.Content, " ")

	for _, word := range words {
		mapTokens[word] = struct{}{}
	}

	for key := range mapTokens {
		tokens = append(tokens, key)
	}

	return tokens
}

// StringsToBytes ... Convert array of strings to array of bytes
func StringsToBytes(content []string) [][]byte {
	var matrix = make([][]byte, 0, len(content))

	for _, element := range content {
		bytesElement := []byte(element)

		if bytesElement != nil {
			matrix = append(matrix, bytesElement)
		}
	}

	return matrix
}

// BytesToStrings ... Convert array of bytes to array of strings
func BytesToStrings(content [][]byte) []string {
	var stringContent []string

	for _, element := range content {
		stringContent = append(stringContent, string(element))
	}

	return stringContent
}
