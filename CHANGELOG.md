# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.10.1](https://github.com/cweill/gotests/compare/v1.10.0...v1.10.1) (2026-11-27)


### Bug Fixes

* ignore directories ending in .go during source discovery ([#204](https://github.com/cweill/gotests/issues/204)) ([5f33eb0](https://github.com/cweill/gotests/commit/5f33eb0a89799f438566a037e99db04b8ce261be))

## [1.10.0](https://github.com/cweill/gotests/compare/v1.9.0...v1.10.0) (2026-11-26)


### Features

* generate black-box tests in a _test package (-package_test) ([#206](https://github.com/cweill/gotests/issues/206)) ([3325c5e](https://github.com/cweill/gotests/commit/3325c5e678aa1a29c4702fabf68b7f37e6800ced)), closes [#33](https://github.com/cweill/gotests/issues/33). Based on [#199](https://github.com/cweill/gotests/pull/199) by @IlyasYOY.
  * The new `-package_test` flag (`Options.PackageTest`) generates tests in `package <pkg>_test` that import the package under test. It implies `-exported`.
  * Calls and types from the package under test are qualified, including inside composite types, receivers and generic instantiations. Works with the default and `testify` templates; custom templates can use the `PackageTest` and `PackageName` fields.
  * Functions whose signatures use unexported types, and methods on unexported types, are skipped and reported (`Skipped <func>: <reason>`, or the new `Options.OnSkip` callback). Unexported struct fields are left out of receiver setup.

### Bug Fixes

* skip benchmark functions when generating tests ([#205](https://github.com/cweill/gotests/issues/205)) ([22dadae](https://github.com/cweill/gotests/commit/22dadae1a5bfcdbd5850b69e1dea688df0449a66)), closes [#93](https://github.com/cweill/gotests/issues/93). Functions taking `*testing.B` no longer produce a test that calls the benchmark.

### Behavior changes

* **Existing `<pkg>_test` test files are now extended in black-box mode automatically**, even without `-package_test` ([#206](https://github.com/cweill/gotests/issues/206)). Previously gotests added unqualified calls and types to these files, which did not compile. Output for such files changes: calls and types are qualified, only exported functions get tests, and the black-box rules above apply. An existing import of the package, including its alias, is reused.
* Adding `-package_test` tests to an existing internal (`package <pkg>`) test file is refused with an error, since it would break that file's tests.

## [1.9.0](https://github.com/cweill/gotests/compare/v1.8.0...v1.9.0) (2025-10-28)

### Features

* **AI-powered test case generation** using local LLMs via Ollama ([#194](https://github.com/cweill/gotests/issues/194))
  * New `-ai` flag fills generated tests with realistic test cases, edge cases and error conditions derived from the function's implementation. Code is sent only to the configured Ollama endpoint.
  * New flags: `-ai-model` (default `qwen2.5-coder:0.5b`), `-ai-endpoint` (default `http://localhost:11434`), `-ai-min-cases` and `-ai-max-cases`

### Other

* Godoc comments for all exported functions ([#195](https://github.com/cweill/gotests/issues/195))
* Test coverage raised from 28.1% to 84.6% ([#196](https://github.com/cweill/gotests/issues/196))

## [1.8.0] - 2025-10-21

### Added
- **Full Go Generics Support** (fixes #165)
  - Generate tests for generic functions with type parameters (`func Foo[T any](val T) T`)
  - Support for all constraint types: `any`, `comparable`, union types (`int64 | float64`), approximation constraints (`~int`)
  - Generate tests for methods on generic types (`func (s *Set[T]) Add(v T)`)
  - Support for multiple type parameters (`func Pair[T, U any](first T, second U)`)
  - Smart type parameter substitution throughout generated tests
  - Intelligent constraint-to-concrete-type mapping:
    - `any` → `int` (simple value type)
    - `comparable` → `string` (commonly comparable type)
    - Union types → first option (e.g., `int64 | float64` → `int64`)
    - Approximation constraints → underlying type (e.g., `~int` → `int`)
  - Support for generic receiver types with proper type argument instantiation
  - Test names automatically strip type parameters (e.g., `Set[T].Add` → `TestSet_Add`)

### Changed
- Improved type constraint analysis in parser
  - Extracts type parameters from both function signatures and type declarations
  - Resolves type parameters for methods on generic types
- Enhanced template rendering with type substitution helpers
  - `TypeArgs` - generates concrete type arguments for generic function calls
  - `FieldType` - substitutes type parameters in field type declarations
  - `ReceiverType` - substitutes type parameters in receiver instantiations

### Technical Details
- Parser enhancements in `internal/goparser/goparser.go`:
  - New `parseTypeDecls()` extracts type parameters from type declarations
  - New `parseTypeParams()` parses AST field lists for type parameters
  - New `extractBaseTypeName()` strips type parameters from receiver types
- Model updates in `internal/models/models.go`:
  - New `TypeParam` struct with Name and Constraint fields
  - Added `TypeParams` field to `Function` struct
  - New helper methods: `IsGeneric()`, `HasGenericReceiver()`, `TypeParamMapping()`
- Render improvements in `internal/render/`:
  - New template functions for type substitution
  - Smart constraint-to-type mapping with support for numeric constraints
  - Proper handling of generic receiver instantiation

### Test Coverage
- 97.5% main package coverage (all tests passing)
- 83.5% overall project coverage
- 100% coverage on all new parser functions
- Comprehensive test suite with 8 generic patterns:
  1. Generic functions with `any` constraint
  2. Generic functions with `comparable` constraint
  3. Union type constraints
  4. Multiple type parameters
  5. Mixed generic/non-generic parameters
  6. Generic functions returning errors
  7. Methods on single-parameter generic types (`Set[T]`)
  8. Methods on multi-parameter generic types (`Map[K,V]`)

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

[1.8.0]: https://github.com/cweill/gotests/compare/v1.7.0...v1.8.0
[1.7.0]: https://github.com/cweill/gotests/compare/v1.6.0...v1.7.0
[1.6.0]: https://github.com/cweill/gotests/releases/tag/v1.6.0
