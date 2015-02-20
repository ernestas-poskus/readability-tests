package kincaid

import (
	"bytes"
	"regexp"
)

var (
	spaceByte = []byte(" ")
	dot       = regexp.MustCompile(`\.`)
)

// WordsCount - count number of words by 'space'
func WordsCount(words [][]byte) int {
	return len(words)
}

// SentenceCount - counts number of sentences by dot
func SentenceCount(text []byte) (count int) {
	count = len(dot.FindAll(text, -1))
	if count == 0 {
		count = 1 // Assume 1 sentence
	}
	return count
}

// Words - returns slice of words in bytes slice
func SplitTextToWords(text []byte) [][]byte {
	return bytes.Split(text, spaceByte)
}
