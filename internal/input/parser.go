// Package input provides input parsing and validation for haiku text.
package input

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/thornzero/haikugo/internal/haiku"
)

// Parser handles different input sources and formats for haiku text.
type Parser struct {
	autosplit bool
}

// New creates a new Parser with the specified autosplit setting.
func New(autosplit bool) *Parser {
	return &Parser{autosplit: autosplit}
}

// ParseFromFile reads and parses a haiku from the specified file path.
func (p *Parser) ParseFromFile(filename string) (*haiku.Haiku, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return p.ParseFromString(string(content))
}

// ParseFromReader reads and parses a haiku from an io.Reader.
func (p *Parser) ParseFromReader(r io.Reader) (*haiku.Haiku, error) {
	content, err := readAll(r)
	if err != nil {
		return nil, err
	}
	return p.ParseFromString(content)
}

// ParseFromString parses a haiku from a string.
func (p *Parser) ParseFromString(text string) (*haiku.Haiku, error) {
	lines, source := p.prepareLines(text)

	if len(lines) != 3 {
		return nil, fmt.Errorf("haiku must have exactly 3 lines, got %d (source=%s)", len(lines), source)
	}

	return haiku.NewHaiku(lines), nil
}

// ParseFromStdin reads and parses a haiku from standard input.
func (p *Parser) ParseFromStdin() (*haiku.Haiku, error) {
	// Check if stdin has data
	stat, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return nil, errors.New("no input provided via stdin")
	}

	return p.ParseFromReader(os.Stdin)
}

// readAll reads all content from an io.Reader as a string.
func readAll(r io.Reader) (string, error) {
	var sb strings.Builder
	scanner := bufio.NewScanner(r)
	first := true

	for scanner.Scan() {
		if !first {
			sb.WriteByte('\n')
		} else {
			first = false
		}
		sb.WriteString(scanner.Text())
	}

	return sb.String(), scanner.Err()
}

// prepareLines processes input text into exactly 3 lines for haiku analysis.
func (p *Parser) prepareLines(text string) ([]string, string) {
	trimmed := strings.TrimSpace(text)

	// If it already has multiple lines, normalize and use them
	parts := nonEmptyLines(trimmed)
	if len(parts) >= 3 {
		return parts[:3], "multiline"
	}

	if !p.autosplit {
		return parts, "raw"
	}

	// Try splitting on common inline separators
	separators := []string{" / ", "/", " | ", " |", "| ", " | ", " — ", " – "}
	for _, sep := range separators {
		if strings.Contains(trimmed, sep) {
			segments := splitAndTrim(trimmed, sep)
			if len(segments) == 3 {
				return segments, "autosplit:" + sep
			}
		}
	}

	// Last resort: split on major punctuation into 3 chunks
	candidates := regexp.MustCompile(`[.!?;:—–…]+`).Split(trimmed, -1)
	segments := filterNonEmpty(candidates)
	if len(segments) >= 3 {
		return segments[:3], "autosplit:punct"
	}

	return parts, "raw"
}

// nonEmptyLines splits text by newlines and returns only non-empty, trimmed lines.
func nonEmptyLines(text string) []string {
	lines := strings.Split(text, "\n")
	result := make([]string, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			result = append(result, line)
		}
	}

	return result
}

// splitAndTrim splits a string by separator and trims each part.
func splitAndTrim(text, separator string) []string {
	parts := strings.Split(text, separator)
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

// filterNonEmpty removes empty strings from a slice after trimming.
func filterNonEmpty(input []string) []string {
	result := make([]string, 0, len(input))
	for _, item := range input {
		if trimmed := strings.TrimSpace(item); trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
