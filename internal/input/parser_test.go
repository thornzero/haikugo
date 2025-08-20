package input

import (
	"strings"
	"testing"
)

func TestParser_ParseFromString(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		autosplit bool
		wantLines []string
		wantError bool
	}{
		{
			name:      "valid three lines",
			input:     "line one\nline two\nline three",
			autosplit: false,
			wantLines: []string{"line one", "line two", "line three"},
			wantError: false,
		},
		{
			name:      "too many lines",
			input:     "one\ntwo\nthree\nfour",
			autosplit: false,
			wantLines: []string{"one", "two", "three"},
			wantError: false,
		},
		{
			name:      "too few lines without autosplit",
			input:     "one line only",
			autosplit: false,
			wantLines: nil,
			wantError: true,
		},
		{
			name:      "autosplit with slash separator",
			input:     "old pond / frog jumps in / splash",
			autosplit: true,
			wantLines: []string{"old pond", "frog jumps in", "splash"},
			wantError: false,
		},
		{
			name:      "autosplit with pipe separator",
			input:     "cherry blossoms | falling petals | spring rain",
			autosplit: true,
			wantLines: []string{"cherry blossoms", "falling petals", "spring rain"},
			wantError: false,
		},
		{
			name:      "autosplit with em dash",
			input:     "morning dew — bird song — peaceful",
			autosplit: true,
			wantLines: []string{"morning dew", "bird song", "peaceful"},
			wantError: false,
		},
		{
			name:      "autosplit with punctuation",
			input:     "old tree. fallen leaves. quiet pond.",
			autosplit: true,
			wantLines: []string{"old tree", "fallen leaves", "quiet pond"},
			wantError: false,
		},
		{
			name:      "empty input",
			input:     "",
			autosplit: false,
			wantLines: nil,
			wantError: true,
		},
		{
			name:      "whitespace only",
			input:     "   \n  \n  ",
			autosplit: false,
			wantLines: nil,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := New(tt.autosplit)
			haiku, err := parser.ParseFromString(tt.input)

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

			if len(haiku.Lines) != len(tt.wantLines) {
				t.Errorf("Got %d lines, want %d", len(haiku.Lines), len(tt.wantLines))
				return
			}

			for i, line := range haiku.Lines {
				if line != tt.wantLines[i] {
					t.Errorf("Line %d = %q, want %q", i, line, tt.wantLines[i])
				}
			}
		})
	}
}

func TestParser_ParseFromReader(t *testing.T) {
	parser := New(false)
	input := "line one\nline two\nline three"
	reader := strings.NewReader(input)

	haiku, err := parser.ParseFromReader(reader)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	expected := []string{"line one", "line two", "line three"}
	if len(haiku.Lines) != len(expected) {
		t.Errorf("Got %d lines, want %d", len(haiku.Lines), len(expected))
		return
	}

	for i, line := range haiku.Lines {
		if line != expected[i] {
			t.Errorf("Line %d = %q, want %q", i, line, expected[i])
		}
	}
}

func TestNonEmptyLines(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "line1\nline2\nline3",
			expected: []string{"line1", "line2", "line3"},
		},
		{
			input:    "line1\n\nline3",
			expected: []string{"line1", "line3"},
		},
		{
			input:    "  line1  \n  \n  line3  ",
			expected: []string{"line1", "line3"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "\n\n\n",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := nonEmptyLines(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("Got %d lines, want %d", len(result), len(tt.expected))
				return
			}
			for i, line := range result {
				if line != tt.expected[i] {
					t.Errorf("Line %d = %q, want %q", i, line, tt.expected[i])
				}
			}
		})
	}
}

func TestSplitAndTrim(t *testing.T) {
	tests := []struct {
		input     string
		separator string
		expected  []string
	}{
		{
			input:     "a / b / c",
			separator: " / ",
			expected:  []string{"a", "b", "c"},
		},
		{
			input:     " a/b/c ",
			separator: "/",
			expected:  []string{"a", "b", "c"},
		},
		{
			input:     "single",
			separator: "/",
			expected:  []string{"single"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := splitAndTrim(tt.input, tt.separator)
			if len(result) != len(tt.expected) {
				t.Errorf("Got %d parts, want %d", len(result), len(tt.expected))
				return
			}
			for i, part := range result {
				if part != tt.expected[i] {
					t.Errorf("Part %d = %q, want %q", i, part, tt.expected[i])
				}
			}
		})
	}
}

func TestFilterNonEmpty(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			input:    []string{"a", "", "c"},
			expected: []string{"a", "c"},
		},
		{
			input:    []string{"  a  ", "  ", "  c  "},
			expected: []string{"a", "c"},
		},
		{
			input:    []string{"", "", ""},
			expected: []string{},
		},
		{
			input:    []string{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		result := filterNonEmpty(tt.input)
		if len(result) != len(tt.expected) {
			t.Errorf("Got %d items, want %d", len(result), len(tt.expected))
			continue
		}
		for i, item := range result {
			if item != tt.expected[i] {
				t.Errorf("Item %d = %q, want %q", i, item, tt.expected[i])
			}
		}
	}
}
