package kincaid

import (
	"fmt"
	"testing"
)

var fleschReadingEaseTest = map[string]string{
	"The Australian platypus is seemingly a hybrid of a mammal and reptilian creature.": "24.44"}

func TestFleschReadingEase(t *testing.T) {
	for sentence, expected := range fleschReadingEaseTest {
		got := fmt.Sprintf("%.2f", FleschReadingEase([]byte(sentence)))
		if got != expected {
			t.Errorf("Expected %f got %s", expected, got)
		}
	}
}

var fleschGradeLevelTest = map[string]string{
	"The Australian platypus is seemingly a hybrid of a mammal and reptilian creature.": "13.08"}

func TestFleschGradeLevel(t *testing.T) {
	for sentence, expected := range fleschGradeLevelTest {
		got := fmt.Sprintf("%.2f", FleschGradeLevel([]byte(sentence)))
		if got != expected {
			t.Errorf("Expected %f got %s", expected, got)
		}
	}
}
