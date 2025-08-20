// Package analyzer provides kigo (season word) detection functionality.
package analyzer

import (
	"sort"
	"strings"
)

// kigo contains a minimal set of season words for haiku analysis.
// In traditional Japanese haiku, kigo are essential seasonal references.
// This is a basic English approximation for educational purposes.
var kigo = []string{
	// Spring
	"spring", "blossom", "cherry", "plum", "thaw", "sprout", "bloom", "nest",
	"robin", "tulip", "daffodil", "crocus", "easter", "rain", "shower",

	// Summer
	"summer", "cicada", "firefly", "thunder", "lightning", "monsoon", "heat",
	"sweat", "beach", "vacation", "pool", "sun", "sunlight", "bright",

	// Autumn/Fall
	"autumn", "fall", "harvest", "maple", "leaf", "leaves", "cricket", "apple",
	"pumpkin", "orange", "golden", "acorn", "migration", "geese", "cool",

	// Winter
	"winter", "snow", "frost", "ice", "blizzard", "solstice", "cold", "freeze",
	"icicle", "mittens", "scarf", "fireplace", "bare", "gray", "grey",

	// Universal/Multiple seasons
	"moon", "full moon", "new moon", "crescent", "stars", "dew", "mist", "fog",
	"wind", "breeze", "cloud", "sky", "earth", "mountain", "river", "lake",
}

// DetectSeasonWords finds season-related words (kigo) in the given text.
// Returns a sorted, deduplicated list of found season words.
func DetectSeasonWords(text string) []string {
	lower := strings.ToLower(text)
	found := make(map[string]struct{})

	for _, seasonWord := range kigo {
		if strings.Contains(lower, seasonWord) {
			found[seasonWord] = struct{}{}
		}
	}

	if len(found) == 0 {
		return nil
	}

	result := make([]string, 0, len(found))
	for word := range found {
		result = append(result, word)
	}

	sort.Strings(result)
	return result
}

// AddSeasonWord adds a custom season word to the detection list.
// This allows extending the kigo list beyond the built-in words.
func AddSeasonWord(word string) {
	// Check if word already exists
	lower := strings.ToLower(word)
	for _, existing := range kigo {
		if existing == lower {
			return
		}
	}
	kigo = append(kigo, lower)
}

// GetSeasonWords returns a copy of the current season words list.
func GetSeasonWords() []string {
	result := make([]string, len(kigo))
	copy(result, kigo)
	return result
}

// HasSeasonWord checks if the text contains any season words.
func HasSeasonWord(text string) bool {
	return len(DetectSeasonWords(text)) > 0
}
