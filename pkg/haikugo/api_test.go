package haikugo

import (
	"testing"
)

func TestNewAnalyzer(t *testing.T) {
	analyzer := NewAnalyzer(1)
	if analyzer == nil {
		t.Error("NewAnalyzer returned nil")
	}

	if analyzer.GetTolerance() != 1 {
		t.Errorf("Expected tolerance 1, got %d", analyzer.GetTolerance())
	}
}

func TestParseHaiku(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
		wantLines []string
	}{
		{
			name:      "valid haiku",
			input:     "old pond\nfrog jumps\nsplash",
			wantError: false,
			wantLines: []string{"old pond", "frog jumps", "splash"},
		},
		{
			name:      "invalid line count",
			input:     "only two\nlines here",
			wantError: true,
			wantLines: nil,
		},
		{
			name:      "empty input",
			input:     "",
			wantError: true,
			wantLines: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			haiku, err := ParseHaiku(tt.input)

			if tt.wantError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if haiku == nil {
				t.Error("Expected haiku but got nil")
				return
			}

			lines := haiku.Lines()
			if len(lines) != len(tt.wantLines) {
				t.Errorf("Got %d lines, want %d", len(lines), len(tt.wantLines))
				return
			}

			for i, line := range lines {
				if line != tt.wantLines[i] {
					t.Errorf("Line %d = %q, want %q", i, line, tt.wantLines[i])
				}
			}
		})
	}
}

func TestParseHaikuWithAutosplit(t *testing.T) {
	input := "old pond / frog jumps / splash"
	haiku, err := ParseHaikuWithAutosplit(input)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	expected := []string{"old pond", "frog jumps", "splash"}
	lines := haiku.Lines()

	if len(lines) != len(expected) {
		t.Errorf("Got %d lines, want %d", len(lines), len(expected))
		return
	}

	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("Line %d = %q, want %q", i, line, expected[i])
		}
	}
}

func TestAnalyzer_Analyze(t *testing.T) {
	analyzer := NewAnalyzer(0)

	haiku, err := ParseHaiku("an old silent pond\na frog jumps into the pond\nsplash silence again")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	metrics := analyzer.Analyze(haiku)
	if metrics == nil {
		t.Error("Analyze returned nil")
		return
	}

	if !metrics.Valid575 {
		t.Error("Expected valid 5-7-5 structure")
	}

	expectedSyllables := []int{5, 7, 5}
	for i, syll := range metrics.LineSyllables {
		if syll != expectedSyllables[i] {
			t.Errorf("Line %d syllables = %d, want %d", i+1, syll, expectedSyllables[i])
		}
	}
}

func TestAnalyzer_IsValid575(t *testing.T) {
	analyzer := NewAnalyzer(0)

	validHaiku, _ := ParseHaiku("an old silent pond\na frog jumps into the pond\nsplash silence again")
	if !analyzer.IsValid575(validHaiku) {
		t.Error("Expected valid haiku to return true")
	}

	invalidHaiku, _ := ParseHaiku("too many syllables here\na frog jumps into the pond\nsplash silence again")
	if analyzer.IsValid575(invalidHaiku) {
		t.Error("Expected invalid haiku to return false")
	}
}

func TestAnalyzer_SetGetTolerance(t *testing.T) {
	analyzer := NewAnalyzer(0)

	if analyzer.GetTolerance() != 0 {
		t.Errorf("Expected initial tolerance 0, got %d", analyzer.GetTolerance())
	}

	analyzer.SetTolerance(2)
	if analyzer.GetTolerance() != 2 {
		t.Errorf("Expected tolerance 2 after setting, got %d", analyzer.GetTolerance())
	}
}

func TestHaiku_Methods(t *testing.T) {
	lines := []string{"line one", "line two", "line three"}
	haiku, err := ParseHaiku("line one\nline two\nline three")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	// Test Lines()
	resultLines := haiku.Lines()
	if len(resultLines) != len(lines) {
		t.Errorf("Lines() returned %d lines, want %d", len(resultLines), len(lines))
	}

	// Test Text()
	expectedText := "line one\nline two\nline three"
	if haiku.Text() != expectedText {
		t.Errorf("Text() = %q, want %q", haiku.Text(), expectedText)
	}

	// Test IsValid()
	if !haiku.IsValid() {
		t.Error("IsValid() should return true for 3-line haiku")
	}
}
