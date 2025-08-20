// Package analyzer provides kireji (cutting word) detection functionality.
package analyzer

import (
	"sort"
	"strings"
)

// kirejiMarkers contains English approximations of Japanese kireji (cutting words).
// These create pauses or emphasis in haiku, similar to punctuation.
var kirejiMarkers = []string{
	"—", "–", "-", "...", "…", ":", ";", "!", "?",
	"ya", "kana", "keri", // Japanese particles sometimes used in English haiku
}

// DetectKireji searches for kireji-like markers in the text.
// Returns whether any were found and a deduplicated, sorted list of matches.
func DetectKireji(text string) (bool, []string) {
	hits := []string{}
	lower := strings.ToLower(text)

	for _, marker := range kirejiMarkers {
		if strings.Contains(lower, marker) {
			hits = append(hits, marker)
		}
	}

	// Deduplicate and sort for stable output
	if len(hits) > 0 {
		seen := make(map[string]struct{})
		result := hits[:0]

		for _, hit := range hits {
			if _, exists := seen[hit]; !exists {
				seen[hit] = struct{}{}
				result = append(result, hit)
			}
		}

		sort.Strings(result)
		return true, result
	}

	return false, nil
}

// AddKirejiMarker adds a custom kireji marker to the detection list.
// This allows extending the detection beyond the built-in markers.
func AddKirejiMarker(marker string) {
	// Check if marker already exists
	for _, existing := range kirejiMarkers {
		if existing == marker {
			return
		}
	}
	kirejiMarkers = append(kirejiMarkers, marker)
}

// GetKirejiMarkers returns a copy of the current kireji markers list.
func GetKirejiMarkers() []string {
	result := make([]string, len(kirejiMarkers))
	copy(result, kirejiMarkers)
	return result
}
