package analyzer

import (
	"testing"
)

func TestCountSyllables(t *testing.T) {
	tests := []struct {
		word     string
		expected int
	}{
		// Basic cases
		{"hello", 2},
		{"world", 1},
		{"beautiful", 3},
		{"cat", 1},
		{"dog", 1},
		{"elephant", 3},

		// Exception cases from dictionary
		{"the", 1},
		{"science", 2},
		{"poem", 2},
		{"fire", 1},
		{"every", 2},
		{"family", 2},
		{"again", 2},

		// Silent 'e' cases
		{"make", 1},
		{"table", 2},  // 'le' ending
		{"simple", 2}, // 'le' ending
		{"apple", 2},  // exception dictionary

		// Edge cases
		{"", 0},
		{"a", 1},
		{"I", 1},
		{"you", 1},
		{"through", 1}, // tricky pronunciation

		// Multiple vowel groups
		{"beautiful", 3}, // beau-ti-ful
		{"creation", 3},  // cre-a-tion
		{"reaction", 3},  // re-ac-tion

		// Consonant 'le' endings
		{"little", 2}, // lit-tle
		{"middle", 2}, // mid-dle
		{"pickle", 2}, // pick-le
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			result := CountSyllables(tt.word)
			if result != tt.expected {
				t.Errorf("CountSyllables(%q) = %d, want %d", tt.word, result, tt.expected)
			}
		})
	}
}

func TestExtractWords(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"hello world", []string{"hello", "world"}},
		{"hello, world!", []string{"hello", "world"}},
		{"it's a beautiful day", []string{"it's", "a", "beautiful", "day"}},
		{"", nil},
		{"123 abc 456", []string{"abc"}},
		{"don't worry, be happy!", []string{"don't", "worry", "be", "happy"}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ExtractWords(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("ExtractWords(%q) returned %d words, want %d", tt.input, len(result), len(tt.expected))
				return
			}
			for i, word := range result {
				if word != tt.expected[i] {
					t.Errorf("ExtractWords(%q)[%d] = %q, want %q", tt.input, i, word, tt.expected[i])
				}
			}
		})
	}
}

func TestCountLineSyllables(t *testing.T) {
	tests := []struct {
		line     string
		expected int
	}{
		{"an old silent pond", 5},
		{"a frog jumps into the pond", 7},
		{"splash silence again", 5},
		{"", 0},
		{"the quick brown fox", 4},
		{"beautiful morning sunshine", 7},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			result := CountLineSyllables(tt.line)
			if result != tt.expected {
				t.Errorf("CountLineSyllables(%q) = %d, want %d", tt.line, result, tt.expected)
			}
		})
	}
}
