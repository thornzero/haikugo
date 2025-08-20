// Package analyzer provides syllable counting functionality.
package analyzer

import (
	"regexp"
	"strings"
	"unicode"
)

// exceptionSyllables contains common irregular words in English syllabification.
var exceptionSyllables = map[string]int{
	"the": 1, "are": 1, "hour": 1, "our": 1, "once": 1, "ones": 1,
	"fire": 1, "fired": 1, "science": 2, "poem": 2, "poems": 2,
	"every": 2, "family": 2, "camera": 3, "business": 2,
	"again": 2, "against": 2, "people": 2, "does": 1, "done": 1,
	"gone": 1, "one": 1, "two": 1, "three": 1, "four": 1, "five": 1,
	"eight": 1, "twelve": 1, "world": 1, "apple": 2, "chocolate": 3,
	"table": 2, "simple": 2, "little": 2, "middle": 2, "pickle": 2,
	"creation": 3, "reaction": 3,
}

var wordRe = regexp.MustCompile(`[A-Za-z']+`)

// ExtractWords extracts all words from a string using regex.
func ExtractWords(s string) []string {
	return wordRe.FindAllString(s, -1)
}

// CountSyllables estimates the number of syllables in a word using heuristic rules.
// Note: English syllable counting is inherently heuristic. This implementation
// provides reasonable estimates but may not be 100% accurate for all words.
func CountSyllables(word string) int {
	if word == "" {
		return 0
	}

	// Check exception dictionary first
	if v, ok := exceptionSyllables[strings.ToLower(word)]; ok {
		return v
	}

	w := []rune(word)
	// Strip leading/trailing non-letters (keep apostrophes inside)
	start, end := 0, len(w)
	for start < end && !isLetter(w[start]) {
		start++
	}
	for end > start && !isLetter(w[end-1]) {
		end--
	}
	if start >= end {
		return 0
	}

	r := w[start:end]
	lower := strings.ToLower(string(r))

	vowels := func(r rune) bool { return strings.ContainsRune("aeiouy", r) }
	count := 0
	prevV := false

	// Count vowel groups
	for _, ch := range lower {
		isV := vowels(ch)
		if isV && !prevV {
			count++
		}
		prevV = isV
	}

	// Handle terminal silent 'e'
	if strings.HasSuffix(lower, "e") && !strings.HasSuffix(lower, "le") && count > 1 {
		count--
	}

	// Handle 'le' ending like "table" - often adds a syllable if preceded by consonant
	if strings.HasSuffix(lower, "le") && len(lower) > 2 {
		pre := rune(lower[len(lower)-3])
		if !strings.ContainsRune("aeiouy", pre) {
			count++
		}
	}

	// Ensure at least one syllable
	if count < 1 {
		count = 1
	}

	return count
}

// isLetter returns true if the rune is a letter or apostrophe.
func isLetter(r rune) bool {
	return unicode.IsLetter(r) || r == '\''
}

// CountLineSyllables counts total syllables in a line of text.
func CountLineSyllables(line string) int {
	words := ExtractWords(line)
	total := 0
	for _, word := range words {
		total += CountSyllables(strings.ToLower(word))
	}
	return total
}
