package haiku

import (
	"testing"
)

func TestNewHaiku(t *testing.T) {
	lines := []string{"line one", "line two", "line three"}
	h := NewHaiku(lines)

	if h == nil {
		t.Fatal("NewHaiku returned nil")
	}

	if len(h.Lines) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(h.Lines))
	}

	for i, line := range lines {
		if h.Lines[i] != line {
			t.Errorf("Lines[%d] = %q, want %q", i, h.Lines[i], line)
		}
	}
}

func TestHaiku_IsValid(t *testing.T) {
	tests := []struct {
		lines    []string
		expected bool
	}{
		{[]string{"one", "two", "three"}, true},
		{[]string{"one", "two"}, false},
		{[]string{"one"}, false},
		{[]string{}, false},
		{[]string{"one", "two", "three", "four"}, false},
	}

	for _, tt := range tests {
		h := NewHaiku(tt.lines)
		result := h.IsValid()
		if result != tt.expected {
			t.Errorf("IsValid() with %d lines = %t, want %t", len(tt.lines), result, tt.expected)
		}
	}
}

func TestHaiku_Text(t *testing.T) {
	tests := []struct {
		lines    []string
		expected string
	}{
		{
			lines:    []string{"line one", "line two", "line three"},
			expected: "line one\nline two\nline three",
		},
		{
			lines:    []string{"single"},
			expected: "single",
		},
		{
			lines:    []string{},
			expected: "",
		},
		{
			lines:    []string{"first", "second"},
			expected: "first\nsecond",
		},
	}

	for _, tt := range tests {
		h := NewHaiku(tt.lines)
		result := h.Text()
		if result != tt.expected {
			t.Errorf("Text() = %q, want %q", result, tt.expected)
		}
	}
}
