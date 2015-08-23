package readability

import (
	"github.com/ernestas-poskus/syllables"
)

// FleschReadingEase - test, higher scores indicate material that is easier to read;
// lower numbers mark passages that are more difficult to read.
// 90.0–100.0	easily understood by an average 11-year-old student
// 60.0–70.0	easily understood by 13- to 15-year-old students
// 0.0–30.0	best understood by university graduates
func FleschReadingEase(text []byte) float64 {
	words := SplitTextToWords(text)
	wordsCount := float64(WordsCount(words))
	if wordsCount == 0 {
		return 0
	}
	sentences := float64(SentenceCount(text))
	syllables := float64(syllables.CountSyllablesInText(words))

	return 206.835 - 1.015*(wordsCount/sentences) - 84.6*(syllables/wordsCount)
}

// FleschGradeLevel - readability tests are used extensively in the field of education
func FleschGradeLevel(text []byte) float64 {
	words := SplitTextToWords(text)
	wordsCount := float64(WordsCount(words))
	if wordsCount == 0 {
		return 0
	}
	sentences := float64(SentenceCount(text))
	syllables := float64(syllables.CountSyllablesInText(words))

	return 0.39*(wordsCount/sentences) + 11.8*(syllables/wordsCount) - 15.59
}
