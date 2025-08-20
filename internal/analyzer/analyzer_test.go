package analyzer

import (
	"testing"

	"github.com/thornzero/haikugo/internal/haiku"
)

func TestAnalyzer_Analyze(t *testing.T) {
	tests := []struct {
		name      string
		lines     []string
		tolerance int
		wantValid bool
		wantSyll  []int
	}{
		{
			name:      "valid traditional haiku",
			lines:     []string{"an old silent pond", "a frog jumps into the pond", "splash silence again"},
			tolerance: 0,
			wantValid: true,
			wantSyll:  []int{5, 7, 5},
		},
		{
			name:      "invalid syllable count",
			lines:     []string{"too many syllables here", "seven syllables in this line", "five more to end"},
			tolerance: 0,
			wantValid: false,
			wantSyll:  []int{7, 8, 4},
		},
		{
			name:      "valid with tolerance",
			lines:     []string{"six syllables here instead", "seven syllables in this line", "four to end"},
			tolerance: 1,
			wantValid: false, // still invalid even with tolerance=1
			wantSyll:  []int{7, 8, 3},
		},
		{
			name:      "haiku with kireji",
			lines:     []string{"old pond—", "frog jumps in", "splash!"},
			tolerance: 0,
			wantValid: false, // 2-3-1 is not 5-7-5
			wantSyll:  []int{2, 3, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := New(tt.tolerance)
			h := haiku.NewHaiku(tt.lines)

			metrics := analyzer.Analyze(h)
			if metrics == nil {
				t.Fatal("Analyze returned nil")
			}

			if metrics.Valid575 != tt.wantValid {
				t.Errorf("Valid575 = %t, want %t", metrics.Valid575, tt.wantValid)
			}

			if len(metrics.LineSyllables) != len(tt.wantSyll) {
				t.Errorf("LineSyllables length = %d, want %d", len(metrics.LineSyllables), len(tt.wantSyll))
			} else {
				for i, syll := range metrics.LineSyllables {
					if syll != tt.wantSyll[i] {
						t.Errorf("LineSyllables[%d] = %d, want %d", i, syll, tt.wantSyll[i])
					}
				}
			}

			// Test basic metrics
			if metrics.TotalWords == 0 {
				t.Error("TotalWords should be greater than 0")
			}
			if metrics.TotalSyllables == 0 {
				t.Error("TotalSyllables should be greater than 0")
			}
			if metrics.LexicalDensity < 0 || metrics.LexicalDensity > 1 {
				t.Errorf("LexicalDensity = %f, should be between 0 and 1", metrics.LexicalDensity)
			}
		})
	}
}

func TestAnalyzer_IsValid575(t *testing.T) {
	tests := []struct {
		syllables []int
		tolerance int
		expected  bool
	}{
		{[]int{5, 7, 5}, 0, true},
		{[]int{4, 7, 5}, 0, false},
		{[]int{4, 7, 5}, 1, true},
		{[]int{3, 7, 5}, 1, false},
		{[]int{5, 8, 5}, 1, true},
		{[]int{5, 9, 5}, 1, false},
		{[]int{5, 7}, 0, false}, // wrong number of lines
		{[]int{}, 0, false},     // empty
	}

	for _, tt := range tests {
		analyzer := New(tt.tolerance)
		result := analyzer.IsValid575(tt.syllables)
		if result != tt.expected {
			t.Errorf("IsValid575(%v, tolerance=%d) = %t, want %t",
				tt.syllables, tt.tolerance, result, tt.expected)
		}
	}
}

func TestAnalyzer_SetGetTolerance(t *testing.T) {
	analyzer := New(0)

	if analyzer.GetTolerance() != 0 {
		t.Errorf("Initial tolerance = %d, want 0", analyzer.GetTolerance())
	}

	analyzer.SetTolerance(2)
	if analyzer.GetTolerance() != 2 {
		t.Errorf("After SetTolerance(2), got %d, want 2", analyzer.GetTolerance())
	}
}

func TestAnalyze_WithSeasonWords(t *testing.T) {
	analyzer := New(0)
	lines := []string{"spring cherry blossoms", "dancing in the warm breeze", "petals fall like snow"}
	h := haiku.NewHaiku(lines)

	metrics := analyzer.Analyze(h)
	if metrics == nil {
		t.Fatal("Analyze returned nil")
	}

	if len(metrics.SeasonWords) == 0 {
		t.Error("Expected to find season words")
	}

	// Should find "spring", "cherry", "blossom", "snow"
	expectedWords := map[string]bool{
		"spring": false, "cherry": false, "blossom": false, "snow": false,
	}

	for _, word := range metrics.SeasonWords {
		if _, exists := expectedWords[word]; exists {
			expectedWords[word] = true
		}
	}

	for word, found := range expectedWords {
		if !found {
			t.Errorf("Expected to find season word '%s'", word)
		}
	}
}

func TestAnalyze_WithKireji(t *testing.T) {
	analyzer := New(0)
	lines := []string{"old pond—", "frog jumps in", "splash!"}
	h := haiku.NewHaiku(lines)

	metrics := analyzer.Analyze(h)
	if metrics == nil {
		t.Fatal("Analyze returned nil")
	}

	if !metrics.HasKireji {
		t.Error("Expected to find kireji")
	}

	if len(metrics.KirejiHits) == 0 {
		t.Error("Expected kireji hits")
	}
}
