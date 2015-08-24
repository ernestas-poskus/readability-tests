package readability

// AutomatedReadabilityIndex - the automated readability index (ARI) is a
// readability test designed to gauge the understandability of a text.
func AutomatedReadabilityIndex(text []byte) float64 {
	words := SplitTextWords(text)
	wordsCount := float64(WordsCount(words))
	if wordsCount == 0 {
		return 0
	}
	sentencesCount := float64(SentenceCount(text))
	charactersCount := 0
	for _, word := range words {
		charactersCount += CharacterCount(word)
	}
	return 4.71*(float64(charactersCount)/wordsCount) + 0.5*(wordsCount/sentencesCount) - 21.43
}
