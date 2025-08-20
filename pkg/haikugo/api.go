// Package haikugo provides a public API for haiku analysis and validation.
//
// This package offers a simple interface for analyzing haiku poems,
// including syllable counting, structure validation, and literary element detection.
//
// Example usage:
//
//	analyzer := haikugo.NewAnalyzer(0) // 0 tolerance for strict 5-7-5
//	haiku, err := haikugo.ParseHaiku("old pond\nfrog jumps in\nsplash!")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	metrics := analyzer.Analyze(haiku)
//	fmt.Printf("Valid 5-7-5: %t\n", metrics.Valid575)
package haikugo

import (
	"github.com/thornzero/haikugo/internal/analyzer"
	"github.com/thornzero/haikugo/internal/haiku"
	"github.com/thornzero/haikugo/internal/input"
)

// Analyzer wraps the internal analyzer for public use.
type Analyzer struct {
	analyzer *analyzer.Analyzer
}

// Haiku wraps the internal haiku structure for public use.
type Haiku struct {
	haiku *haiku.Haiku
}

// Metrics wraps the internal metrics structure for public use.
type Metrics = haiku.Metrics

// NewAnalyzer creates a new haiku analyzer with the specified syllable tolerance.
// Tolerance allows for flexibility in the 5-7-5 pattern (e.g., tolerance=1 allows 4-6, 6-8, 4-6).
func NewAnalyzer(tolerance int) *Analyzer {
	return &Analyzer{
		analyzer: analyzer.New(tolerance),
	}
}

// ParseHaiku parses a haiku from a string. The string should contain exactly 3 lines.
func ParseHaiku(text string) (*Haiku, error) {
	parser := input.New(false)
	h, err := parser.ParseFromString(text)
	if err != nil {
		return nil, err
	}
	return &Haiku{haiku: h}, nil
}

// ParseHaikuWithAutosplit parses a haiku from a string with automatic line splitting.
// This attempts to split single-line input into 3 lines using common separators.
func ParseHaikuWithAutosplit(text string) (*Haiku, error) {
	parser := input.New(true)
	h, err := parser.ParseFromString(text)
	if err != nil {
		return nil, err
	}
	return &Haiku{haiku: h}, nil
}

// ParseHaikuFromFile reads and parses a haiku from a file.
func ParseHaikuFromFile(filename string) (*Haiku, error) {
	parser := input.New(false)
	h, err := parser.ParseFromFile(filename)
	if err != nil {
		return nil, err
	}
	return &Haiku{haiku: h}, nil
}

// Analyze performs comprehensive analysis of the haiku and returns detailed metrics.
func (a *Analyzer) Analyze(h *Haiku) *Metrics {
	return a.analyzer.Analyze(h.haiku)
}

// IsValid575 checks if the haiku follows the traditional 5-7-5 syllable pattern.
func (a *Analyzer) IsValid575(h *Haiku) bool {
	metrics := a.analyzer.Analyze(h.haiku)
	if metrics == nil {
		return false
	}
	return metrics.Valid575
}

// SetTolerance updates the syllable tolerance for validation.
func (a *Analyzer) SetTolerance(tolerance int) {
	a.analyzer.SetTolerance(tolerance)
}

// GetTolerance returns the current syllable tolerance.
func (a *Analyzer) GetTolerance() int {
	return a.analyzer.GetTolerance()
}

// Lines returns the lines of the haiku.
func (h *Haiku) Lines() []string {
	return h.haiku.Lines
}

// Text returns the full text of the haiku as a single string.
func (h *Haiku) Text() string {
	return h.haiku.Text()
}

// IsValid returns true if the haiku has exactly 3 lines.
func (h *Haiku) IsValid() bool {
	return h.haiku.IsValid()
}
