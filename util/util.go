package util

import (
	"strings"
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
