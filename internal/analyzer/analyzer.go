// Package analyzer provides comprehensive haiku analysis functionality.
package analyzer

import (
	"strings"

	"github.com/thornzero/haikugo/internal/haiku"
)

// Analyzer provides methods for analyzing haiku poems.
type Analyzer struct {
	tolerance int
}

// New creates a new Analyzer with the specified syllable tolerance.
func New(tolerance int) *Analyzer {
	return &Analyzer{tolerance: tolerance}
}

// Analyze performs comprehensive analysis of a haiku and returns metrics.
func (a *Analyzer) Analyze(h *haiku.Haiku) *haiku.Metrics {
	if !h.IsValid() {
		return nil
	}

	m := &haiku.Metrics{
		Lines:     h.Lines,
		Tolerance: a.tolerance,
	}

	m.LineSyllables = make([]int, 3)
	m.LineWords = make([]int, 3)

	var totalChars, totalLetters int
	uniqueWords := make(map[string]struct{})

	// Analyze each line
	for i, line := range h.Lines {
		words := ExtractWords(line)
		m.LineWords[i] = len(words)

		for _, word := range words {
			lowerWord := strings.ToLower(word)
			uniqueWords[lowerWord] = struct{}{}
			m.LineSyllables[i] += CountSyllables(lowerWord)
			totalLetters += len([]rune(lowerWord))
		}

		totalChars += len([]rune(line))
	}

	// Calculate totals and derived metrics
	m.TotalSyllables = m.LineSyllables[0] + m.LineSyllables[1] + m.LineSyllables[2]
	m.TotalWords = m.LineWords[0] + m.LineWords[1] + m.LineWords[2]
	m.UniqueWords = len(uniqueWords)

	if m.TotalWords > 0 {
		m.LexicalDensity = float64(m.UniqueWords) / float64(m.TotalWords)
		m.AvgWordLen = float64(totalLetters) / float64(m.TotalWords)
	}

	// Detect literary elements
	fullText := strings.Join(h.Lines, " ")
	m.HasKireji, m.KirejiHits = DetectKireji(fullText)
	m.SeasonWords = DetectSeasonWords(fullText)

	// Validate 5-7-5 structure
	m.Valid575 = a.IsValid575(m.LineSyllables)

	return m
}

// IsValid575 checks if the syllable pattern matches 5-7-5 within tolerance.
func (a *Analyzer) IsValid575(syllables []int) bool {
	if len(syllables) != 3 {
		return false
	}

	expected := []int{5, 7, 5}
	for i := 0; i < 3; i++ {
		if abs(syllables[i]-expected[i]) > a.tolerance {
			return false
		}
	}

	return true
}

// SetTolerance updates the syllable tolerance for validation.
func (a *Analyzer) SetTolerance(tolerance int) {
	a.tolerance = tolerance
}

// GetTolerance returns the current syllable tolerance.
func (a *Analyzer) GetTolerance() int {
	return a.tolerance
}

// abs returns the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
