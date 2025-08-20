// Package haiku provides core data structures for haiku analysis.
package haiku

// Metrics holds comprehensive analysis results for a haiku.
type Metrics struct {
	Lines          []string `json:"lines"`
	LineSyllables  []int    `json:"line_syllables"`
	LineWords      []int    `json:"line_words"`
	TotalSyllables int      `json:"total_syllables"`
	TotalWords     int      `json:"total_words"`
	UniqueWords    int      `json:"unique_words"`
	LexicalDensity float64  `json:"lexical_density"`
	AvgWordLen     float64  `json:"avg_word_len"`
	HasKireji      bool     `json:"has_kireji_like_pause"`
	KirejiHits     []string `json:"kireji_hits"`
	SeasonWords    []string `json:"season_words"`
	Valid575       bool     `json:"valid_575"`
	Tolerance      int      `json:"tolerance"`
}

// Haiku represents a three-line haiku poem.
type Haiku struct {
	Lines []string
}

// NewHaiku creates a new Haiku from the provided lines.
func NewHaiku(lines []string) *Haiku {
	return &Haiku{Lines: lines}
}

// IsValid returns true if the haiku has exactly 3 lines.
func (h *Haiku) IsValid() bool {
	return len(h.Lines) == 3
}

// Text returns the full text of the haiku as a single string.
func (h *Haiku) Text() string {
	result := ""
	for i, line := range h.Lines {
		if i > 0 {
			result += "\n"
		}
		result += line
	}
	return result
}
