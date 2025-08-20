# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial project structure and modular architecture
- Comprehensive haiku analysis with syllable counting
- Kireji (cutting word) detection
- Kigo (season word) detection
- Multiple input methods (stdin, file, inline, autosplit)
- JSON and human-readable output formats
- Configurable syllable tolerance
- Public Go library API
- CLI application (haikuctl)
- Comprehensive test suite
- Build automation with Makefile
- CI/CD pipeline with GitHub Actions
- GoReleaser configuration for releases

### Changed
- Refactored from monolithic single-file to modular architecture
- Improved syllable counting algorithm with exception dictionary
- Enhanced error handling and user experience

### Technical Details
- Go 1.24+ compatibility
- Modular package structure (internal/, pkg/, cmd/)
- 100% test coverage for core functionality
- Professional project layout following Go best practices

## [0.1.0] - 2024-12-19

### Added
- Initial release
- Core haiku analysis functionality
- CLI tool with all major features
- Go library for integration

---

## Version History

- **0.1.0**: Initial release with full feature set
- **Unreleased**: Development version with latest improvements

## Contributing

When adding new features or making changes, please update this changelog following the [Keep a Changelog](https://keepachangelog.com/) format.
