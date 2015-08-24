package readability

import (
	"bytes"
	"regexp"
	"unicode/utf8"
)

var (
	emptyByte = []byte("")
	spaceByte = []byte(" ")
	dot       = regexp.MustCompile(`\.`)
	alpha     = regexp.MustCompile(`[[:punct:]]`)
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

// Alpha - [a-z]
func Alpha(text []byte) []byte {
	return alpha.ReplaceAll(text, emptyByte)
}

// SplitTextWords - returns slice of words in bytes slice
func SplitTextWords(text []byte) [][]byte {
	return bytes.Split(text, spaceByte)
}

// CharacterCount - counts character in text
func CharacterCount(text []byte) int {
	return utf8.RuneCount(Alpha(text))
}
