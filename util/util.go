package util

import (
	"log"
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
func TimeTrack(start time.Time, name string) time.Duration {
	elapsed := time.Since(start)
	log.Printf("TimeTrack: %s took %s", name, elapsed)
	return elapsed
}
