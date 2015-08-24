package readability

import "testing"

var wordsCountTest = map[string]int{
	"creature some a": 3,
	"some a":          2,
	"a":               1,
}

func TestWordsCount(t *testing.T) {
	for word, expectedWordsCount := range wordsCountTest {
		got := WordsCount(SplitTextWords([]byte(word)))
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

var characterCountTests = []struct {
	input    string
	expected int
}{
	{"aoms", 4},
	{"šeimeną", 7},
}

func TestCharacterCount(t *testing.T) {
	for _, test := range characterCountTests {
		got := CharacterCount([]byte(test.input))
		if got != test.expected {
			t.Errorf("Expected %d character count for %s, got %d", test.expected, test.input, got)
		}
	}
}

var alphaTests = []struct {
	input    string
	expected string
}{
	{"(Dale and Chall, 1948)", "Dale and Chall 1948"},
}

func TestAlpha(t *testing.T) {
	for _, test := range alphaTests {
		got := string(Alpha([]byte(test.input)))
		if got != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, got)
		}
	}
}
