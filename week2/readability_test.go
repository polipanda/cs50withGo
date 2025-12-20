package week2

import (
	"fmt"
	"testing"
)

func TestDetermineReadingLevel(t *testing.T) {
	for _, scenario := range readabilityScenarios() {
		fmt.Println(scenario)
		result := determineReadingLevel(scenario.text)
		fmt.Println(result)
		if result != scenario.level {
			t.Errorf("expected %s, got %s", scenario.level, result)
		}
	}
}

type readabilityScenario struct {
	text  string
	level string
}

func readabilityScenarios() []readabilityScenario {
	return []readabilityScenario{
		{
			text:  "One fish. Two fish. Red fish. Blue fish.",
			level: BEFOREGRADE1,
		},
		{
			text:  "Would you like them here or there? I would not like them here or there. I would not like them anywhere.",
			level: GRADE2,
		},
		{
			text:  "Congratulations! Today is your day. You're off to Great Places! You're off and away!",
			level: GRADE3,
		},
		{
			text: `Harry Potter was a highly unusual boy in many ways. For one thing, he hated the summer holidays
				more than any other time of year. For another, he really wanted to do his homework, but was forced 
				to do it in secret, in the dead of the night. And he also happened to be a wizard.`,
			level: GRADE5,
		},
		{
			text:  "In my younger and more vulnerable years my father gave me some advice that I've been turning over in my mind ever since.",
			level: GRADE7,
		},
		{
			text: `Alice was beginning to get very tired of sitting by her sister on the bank, and of having nothing 
				to do: once or twice she had peeped into the book her sister was reading, but it had no pictures or 
				conversations in it, "and what is the use of a book," thought Alice "without pictures or conversation?"`,
			level: GRADE8,
		},
		{
			text: `When he was nearly thirteen, my brother Jem got his arm badly broken at the elbow. When it healed, 
				and Jem's fears of never being able to play football were assuaged, he was seldom self-conscious about 
				his injury. His left arm was somewhat shorter than his right; when he stood or walked, the back of his 
				hand was at right angles to his body, his thumb parallel to his thigh.`,
			level: GRADE8,
		},
		{
			text:  "There are more things in Heaven and Earth, Horatio, than are dreamt of in your philosophy.",
			level: GRADE9,
		},
		{
			text: `It was a bright cold day in April, and the clocks were striking thirteen. Winston Smith, his chin 
				nuzzled into his breast in an effort to escape the vile wind, slipped quickly through the glass doors 
				of Victory Mansions, though not quickly enough to prevent a swirl of gritty dust from entering along 
				with him.`,
			level: GRADE10,
		},
		{
			text: `A large class of computational problems involve the determination of properties of graphs, digraphs,
				integers, arrays of integers, finite families of finite sets, boolean formulas and elements of other 
				countable domains.`,
			level: GRADE16PLUS,
		},
	}
}

func TestCountLetters(t *testing.T) {
	text := "123asdfg43!wer word    word word"
	result := countLetters(text)
	if result != 20 {
		t.Errorf("countLetters(%s) = %d, want 20", text, result)
	}
}

func TestCountSentences(t *testing.T) {
	text := "123asdfg43!wer word    word word."
	result := countSentences(text)
	if result != 2 {
		t.Errorf("countLetters(%s) = %d, want 2", text, result)
	}
}

func TestCountWords(t *testing.T) {
	text := "123asdfg43!wer word    word word."
	result := countWords(text)
	if result != 4 {
		t.Errorf("countLetters(%s) = %d, want 4", text, result)
	}
}

func TestCalculateIndex(t *testing.T) {
	letters := 400.0
	sentences := 6.5

	index := fmt.Sprintf("%.4f", calculateIndex(letters, sentences))
	expected := "5.7960"
	if index != expected {
		t.Errorf("calculateIndex(%f, %f) = %s, want %s", letters, sentences, index, expected)
	}
}

func TestGradeToString(t *testing.T) {
	m := map[int]string{
		-1: BEFOREGRADE1,
		0:  BEFOREGRADE1,
		2:  GRADE2,
		3:  GRADE3,
		5:  GRADE5,
		7:  GRADE7,
		8:  GRADE8,
		9:  GRADE9,
		10: GRADE10,
		16: GRADE16PLUS,
		32: GRADE16PLUS,
	}

	for k, v := range m {
		if result := gradeToString(k); result != v {
			t.Errorf("expected: %s, got: %s", v, result)
		}
	}
}
