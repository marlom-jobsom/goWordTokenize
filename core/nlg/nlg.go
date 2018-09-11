package nlg

import (
	"strings"
)

// NLG handles operations for text over RPC calls
type NLG struct{}

// TextTokenizeRPC tokenizes a text over RPC calls
func (NLG) TextTokenizeRPC(text string, result *[]string) error {
	*result = TextTokenize(text)
	return nil
}

// TextTokenize tokenizes a text
func TextTokenize(content string) []string {
	var tokens []string
	mapTokens := make(map[string]struct{})
	words := strings.Split(content, " ")

	for _, word := range words {
		mapTokens[word] = struct{}{}
	}

	for key := range mapTokens {
		tokens = append(tokens, key)
	}

	return tokens
}
