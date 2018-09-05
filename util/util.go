package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"word-tokenize-middleware-socket/core"
)

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

// TimeTrack ... Timing function calls
func TimeTrack(start time.Time, name string, filePath string) time.Duration {
	elapsed := time.Since(start)
	log.Printf("TimeTrack: %s took %s", name, elapsed)

	elapsedNano := elapsed.Nanoseconds()
	AppendContentFile(filePath, fmt.Sprint(elapsedNano))
	return elapsed
}

// AppendContentFile ... Append content to a existing file
func AppendContentFile(filePath string, content string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Create(filePath)
	}

	file, _ := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(content + "\n")
	writer.Flush()
}
