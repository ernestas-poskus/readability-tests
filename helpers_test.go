package kincaid

import "testing"

var wordsCountTest = map[string]int{
	"creature some a": 3,
	"some a":          2,
	"a":               1,
}

func TestWordsCount(t *testing.T) {
	for word, expectedWordsCount := range wordsCountTest {
		got := WordsCount(SplitTextToWords([]byte(word)))
		if got != expectedWordsCount {
			t.Errorf("Expected %d got %d", expectedWordsCount, got)
		}
	}
}

var sentenceCountTest = map[string]int{
	"A set of words that is complete. In itself, typically containing a subject.": 2,
	"A set of words that is complete.":                                            1,
	"A set of words":                                                              1,
}

func TestSentenceCount(t *testing.T) {
	for sentence, expectedSentenceCount := range sentenceCountTest {
		got := SentenceCount([]byte(sentence))
		if got != expectedSentenceCount {
			t.Errorf("Expected %d got %d", expectedSentenceCount, got)
		}
	}
}
