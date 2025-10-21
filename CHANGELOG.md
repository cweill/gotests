# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.7.0] - 2025-10-20

### Added
- **Recursive directory support**: You can now use patterns like `pkg/...` to generate tests for all Go files in a directory tree (fixes #186)
- Documentation for the `-named` flag in README (PR #185)
- CLAUDE.md for AI-assisted development context
- `t.Parallel()` at the top-level test function when `-parallel` flag is used (fixes #188)

### Changed
- **BREAKING**: Minimum Go version increased from 1.6 to 1.22
  - Leverages Go 1.22's per-iteration loop variable scoping
  - Removes unnecessary `tt := tt` and `name := name` shadowing in generated tests
  - Cleaner, more modern generated test code
- **Template Error Handling**: Use `t.Fatalf()` instead of `t.Errorf() + return` in subtests with return values (PR #184)
  - Better test ergonomics and clearer failure semantics
  - Prevents misleading test output when errors occur
- **Dependencies**: Replaced third-party bindata tools with stdlib `embed` package (PR #181)
  - Removed 834+ lines of generated code
  - Simplified build process (no more `go generate` needed)
  - Better maintainability and reduced dependencies
- **Installation**: Updated README to use `go install` instead of deprecated `go get` (PR #180)
- **CI/CD**: Updated GitHub Actions workflow
  - Now tests only Go 1.22.x and 1.23.x (down from 8 versions)
  - Updated actions/setup-go v2→v5
  - Updated actions/checkout v2→v4
  - Simplified coverage reporting with coverallsapp/github-action@v2

### Fixed
- **Import Path Bug**: Fixed missing import paths when receiver types and methods are defined in different files (PR #179)
  - Now correctly collects imports from all files in the package
  - Prevents compilation errors in generated test files
- **t.Parallel() Placement**: Moved `t.Parallel()` to the correct location at the top of test functions
  - Satisfies `tparallel` linter requirements
  - Ensures proper parallel test execution

### Removed
- Support for Go versions older than 1.22
- Bindata dependency (go-bindata, esc) in favor of stdlib embed
- Unnecessary loop variable shadowing in generated tests

## [1.6.0] - 2020-XX-XX

Previous release (pre-changelog).

[1.7.0]: https://github.com/cweill/gotests/compare/v1.6.0...v1.7.0
[1.6.0]: https://github.com/cweill/gotests/releases/tag/v1.6.0
