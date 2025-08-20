package analyzer

import (
	"reflect"
	"testing"
)

func TestDetectSeasonWords(t *testing.T) {
	tests := []struct {
		text     string
		expected []string
	}{
		{
			text:     "cherry blossoms in spring",
			expected: []string{"blossom", "cherry", "spring"},
		},
		{
			text:     "summer heat and cicada song",
			expected: []string{"cicada", "heat", "summer"},
		},
		{
			text:     "autumn leaves fall down",
			expected: []string{"autumn", "fall", "leaves"}, // "leaves" is found, not "leaf"
		},
		{
			text:     "winter snow covers the ground",
			expected: []string{"snow", "winter"},
		},
		{
			text:     "the moon shines bright",
			expected: []string{"bright", "moon"}, // "bright" is also a season word
		},
		{
			text:     "no seasonal references here",
			expected: nil,
		},
		{
			text:     "",
			expected: nil,
		},
		{
			text:     "spring rain spring flowers",
			expected: []string{"rain", "spring"}, // should deduplicate
		},
	}

	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			result := DetectSeasonWords(tt.text)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("DetectSeasonWords(%q) = %v, want %v", tt.text, result, tt.expected)
			}
		})
	}
}

func TestHasSeasonWord(t *testing.T) {
	tests := []struct {
		text     string
		expected bool
	}{
		{"spring morning", true},
		{"no seasonal words", false},
		{"winter wonderland", true},
		{"", false},
		{"SUMMER fun", true}, // case insensitive
	}

	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			result := HasSeasonWord(tt.text)
			if result != tt.expected {
				t.Errorf("HasSeasonWord(%q) = %t, want %t", tt.text, result, tt.expected)
			}
		})
	}
}

func TestAddSeasonWord(t *testing.T) {
	// Get original length
	originalWords := GetSeasonWords()
	originalLen := len(originalWords)

	// Add a new season word
	testWord := "testseason"
	AddSeasonWord(testWord)

	// Check it was added
	newWords := GetSeasonWords()
	if len(newWords) != originalLen+1 {
		t.Errorf("Expected %d words after adding, got %d", originalLen+1, len(newWords))
	}

	// Add the same word again - should not duplicate
	AddSeasonWord(testWord)

	afterDupWords := GetSeasonWords()
	if len(afterDupWords) != originalLen+1 {
		t.Errorf("Expected %d words after duplicate add, got %d", originalLen+1, len(afterDupWords))
	}

	// Test detection with the new word
	text := "beautiful testseason day"
	result := DetectSeasonWords(text)
	found := false
	for _, word := range result {
		if word == testWord {
			found = true
			break
		}
	}
	if !found {
		t.Error("Should have found the custom season word")
	}
}
