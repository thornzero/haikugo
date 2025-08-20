package analyzer

import (
	"reflect"
	"testing"
)

func TestDetectKireji(t *testing.T) {
	tests := []struct {
		text        string
		expectFound bool
		expectHits  []string
	}{
		{
			text:        "old pond — frog jumps in",
			expectFound: true,
			expectHits:  []string{"—"},
		},
		{
			text:        "no pauses here at all",
			expectFound: false,
			expectHits:  nil,
		},
		{
			text:        "splash... silence again",
			expectFound: true,
			expectHits:  []string{"..."},
		},
		{
			text:        "hello, world! how are you?",
			expectFound: true,
			expectHits:  []string{"!", "?"},
		},
		{
			text:        "morning dew; birds singing",
			expectFound: true,
			expectHits:  []string{";"},
		},
		{
			text:        "cherry blossoms — spring rain — peaceful",
			expectFound: true,
			expectHits:  []string{"—"}, // should deduplicate
		},
		{
			text:        "autumn leaves kana",
			expectFound: true,
			expectHits:  []string{"kana"},
		},
		{
			text:        "",
			expectFound: false,
			expectHits:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			found, hits := DetectKireji(tt.text)
			if found != tt.expectFound {
				t.Errorf("DetectKireji(%q) found = %t, want %t", tt.text, found, tt.expectFound)
			}
			if !reflect.DeepEqual(hits, tt.expectHits) {
				t.Errorf("DetectKireji(%q) hits = %v, want %v", tt.text, hits, tt.expectHits)
			}
		})
	}
}

func TestAddKirejiMarker(t *testing.T) {
	// Get original length
	originalMarkers := GetKirejiMarkers()
	originalLen := len(originalMarkers)

	// Add a new marker
	testMarker := "@@test@@"
	AddKirejiMarker(testMarker)

	// Check it was added
	newMarkers := GetKirejiMarkers()
	if len(newMarkers) != originalLen+1 {
		t.Errorf("Expected %d markers after adding, got %d", originalLen+1, len(newMarkers))
	}

	// Add the same marker again - should not duplicate
	AddKirejiMarker(testMarker)

	afterDupMarkers := GetKirejiMarkers()
	if len(afterDupMarkers) != originalLen+1 {
		t.Errorf("Expected %d markers after duplicate add, got %d", originalLen+1, len(afterDupMarkers))
	}

	// Test detection with the new marker
	text := "hello @@test@@ world"
	found, hits := DetectKireji(text)
	if !found {
		t.Error("Should have found the custom kireji marker")
	}
	if len(hits) != 1 || hits[0] != testMarker {
		t.Errorf("Expected hits [%s], got %v", testMarker, hits)
	}
}
