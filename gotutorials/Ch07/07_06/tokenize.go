// Code from https://github.com/LinkedInLearning/go-essential-training-2446262/blob/main/Ch07/07_06/tokenize.go
package nlp

import (
	"regexp"
	"strings"
)

var (
	wordRe = regexp.MustCompile(`\w+`)
)

// Tokenize returns slice of tokens found in text (lower case)
func Tokenize(text string) []string {
	var tokens []string
	for _, tok := range wordRe.FindAllString(text, -1) {
		tok = strings.ToLower(tok)
		tokens = append(tokens, tok)
	}

	return tokens
}