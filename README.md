# HaikuGo

[![Go Version](https://img.shields.io/github/go-mod/go-version/thornzero/haikugo)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/thornzero/haikugo)](https://goreportcard.com/report/github.com/thornzero/haikugo)

A comprehensive Go library and CLI tool for analyzing and validating haiku poems. HaikuGo provides syllable counting, structure validation, and detection of traditional haiku elements like kireji (cutting words) and kigo (season words).

## Features

- **5-7-5 Structure Validation**: Validates traditional haiku syllable patterns with configurable tolerance
- **Comprehensive Metrics**: Syllable counts, word statistics, lexical density analysis
- **Literary Element Detection**:
  - Kireji (cutting words) detection with English approximations
  - Kigo (season words) identification from built-in lexicon
- **Flexible Input Methods**: stdin, files, inline text, auto-splitting
- **Multiple Output Formats**: Human-readable and JSON output
- **Library and CLI**: Use as a Go library or standalone command-line tool

## Installation

### From Source

```bash
git clone https://github.com/thornzero/haikugo.git
cd haikugo
make install
```

### Go Install

```bash
go install github.com/thornzero/haikugo/cmd/haikuctl@latest
```

## Quick Start

### Command Line Usage

```bash
# Basic validation
echo -e "an old silent pond\na frog jumps into the pond\nsplash silence again" | haikuctl

# From file
haikuctl --file my_haiku.txt

# JSON output
haikuctl --json --file haiku.txt

# Auto-split single line
haikuctl --autosplit "old pond / frog jumps in / splash"

# With syllable tolerance
haikuctl --tolerant=1 --file haiku.txt

# Exit code mode (for scripts)
haikuctl --exit-code < haiku.txt && echo "Valid haiku!"
```

### Library Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/thornzero/haikugo/pkg/haikugo"
)

func main() {
    // Parse a haiku
    haiku, err := haikugo.ParseHaiku(`an old silent pond
a frog jumps into the pond
splash silence again`)
    if err != nil {
        log.Fatal(err)
    }
    
    // Create analyzer with strict 5-7-5 validation
    analyzer := haikugo.NewAnalyzer(0)
    
    // Analyze the haiku
    metrics := analyzer.Analyze(haiku)
    
    fmt.Printf("Valid 5-7-5: %t\n", metrics.Valid575)
    fmt.Printf("Syllables: %v\n", metrics.LineSyllables)
    fmt.Printf("Season words: %v\n", metrics.SeasonWords)
    fmt.Printf("Has kireji: %t\n", metrics.HasKireji)
}
```

## API Reference

### Core Types

```go
// Create a new analyzer with syllable tolerance
analyzer := haikugo.NewAnalyzer(tolerance int)

// Parse haiku from string
haiku, err := haikugo.ParseHaiku(text string)

// Parse with auto-line-splitting
haiku, err := haikugo.ParseHaikuWithAutosplit(text string)

// Parse from file
haiku, err := haikugo.ParseHaikuFromFile(filename string)
```

### Analysis

```go
// Comprehensive analysis
metrics := analyzer.Analyze(haiku)

// Quick validation check
isValid := analyzer.IsValid575(haiku)

// Adjust tolerance
analyzer.SetTolerance(1)
```

### Metrics Structure

```go
type Metrics struct {
    Lines          []string  // The haiku lines
    LineSyllables  []int     // Syllables per line
    LineWords      []int     // Words per line
    TotalSyllables int       // Total syllable count
    TotalWords     int       // Total word count
    UniqueWords    int       // Unique word count
    LexicalDensity float64   // Ratio of unique to total words
    AvgWordLen     float64   // Average word length
    HasKireji      bool      // Contains cutting words
    KirejiHits     []string  // Found cutting words
    SeasonWords    []string  // Found season words
    Valid575       bool      // Matches 5-7-5 pattern
    Tolerance      int       // Syllable tolerance used
}
```

## Examples

### Traditional Haiku

```bash
$ echo -e "an old silent pond\na frog jumps into the pond\nsplash silence again" | haikuctl

Haiku (3 lines):
1: an old silent pond
2: a frog jumps into the pond
3: splash silence again

Syllables per line: [5 7 5]
Words per line:     [4 6 3]
Total syllables:    17
Total words:        13 (unique 12, lexical density 0.92)
Avg word length:    4.15
Kireji-like pause:  no
Season words:       none

Structure: VALID (tolerance ±0)
```

### With Season Words

```bash
$ echo -e "cherry blossoms fall\ndancing in the spring breeze\npetals kiss the earth" | haikuctl

Haiku (3 lines):
1: cherry blossoms fall
2: dancing in the spring breeze
3: petals kiss the earth

Syllables per line: [5 7 5]
Words per line:     [3 5 4]
Total syllables:    17
Total words:        12 (unique 12, lexical density 1.00)
Avg word length:    4.42
Kireji-like pause:  no
Season words:       blossom, cherry, spring

Structure: VALID (tolerance ±0)
```

### JSON Output

```bash
$ echo -e "old pond—\nfrog jumps in\nsplash!" | haikuctl --json
{
  "lines": ["old pond—", "frog jumps in", "splash!"],
  "line_syllables": [2, 3, 1],
  "line_words": [2, 3, 1],
  "total_syllables": 6,
  "total_words": 6,
  "unique_words": 6,
  "lexical_density": 1.0,
  "avg_word_len": 4.5,
  "has_kireji_like_pause": true,
  "kireji_hits": ["!", "—"],
  "season_words": null,
  "valid_575": false,
  "tolerance": 0
}
```

## Development

### Prerequisites

- Go 1.24 or later
- Make (optional, for convenience)

### Building

```bash
# Development build
make build

# Run tests
make test

# Run with coverage
make coverage

# Full development setup
make setup

# All checks (format, test, lint, build)
make all
```

### Project Structure

```tree
haikugo/
├── cmd/haikuctl/          # CLI application
├── internal/              # Internal packages
│   ├── analyzer/          # Core analysis logic
│   ├── haiku/            # Haiku data structures
│   └── input/            # Input parsing
├── pkg/haikugo/          # Public API
├── testdata/             # Test fixtures
└── Makefile              # Build automation
```

### Testing

```bash
# Run all tests
go test ./...

# Run specific package tests
go test ./internal/analyzer

# Run with verbose output
go test -v ./...

# Generate coverage report
make coverage
```

## Algorithm Notes

### Syllable Counting

HaikuGo uses a heuristic approach for English syllable counting:

1. **Exception Dictionary**: Common irregular words are handled explicitly
2. **Vowel Group Counting**: Consecutive vowels count as one syllable
3. **Silent 'e' Handling**: Terminal 'e' is often silent unless preceded by 'l'
4. **Consonant + 'le'**: Words ending in consonant + 'le' get an extra syllable

**Important**: English syllable counting is inherently heuristic. Results should be treated as estimates. For perfect accuracy, use a comprehensive phonetic dictionary.

### Literary Elements

- **Kireji Detection**: Searches for punctuation and Japanese particles that create pauses
- **Kigo Detection**: Built-in lexicon of seasonal references organized by season
- Both systems are extensible and can be customized

## Configuration

### CLI Flags

- `--file`: Read from file instead of stdin
- `--json`: Output JSON instead of human-readable format
- `--tolerant`: Allow syllable deviation (e.g., 1 allows 4-6, 6-8, 4-6)
- `--exit-code`: Use exit codes (0=valid, 1=error, 2=invalid)
- `--autosplit`: Try to split single-line input into 3 lines

### Library Configuration

```go
// Syllable tolerance
analyzer.SetTolerance(1)

// Add custom season words
// (See internal packages for advanced customization)
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes and add tests
4. Run the full test suite (`make all`)
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

### Development Guidelines

- Write tests for new functionality
- Follow Go naming conventions
- Update documentation for public APIs
- Run `make all` before submitting PRs

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by the traditional Japanese art of haiku poetry
- Syllable counting algorithms adapted from various linguistic research
- Built with love for both programming and poetry

## Roadmap

- [ ] Japanese syllable/mora counting
- [ ] Web interface
- [ ] Additional output formats (YAML, CSV)
- [ ] Batch processing capabilities
- [ ] Advanced linguistic analysis
- [ ] Integration with popular text editors

---

*"In the old pond... a frog leaps in... the sound of water" - Matsuo Bashō*
