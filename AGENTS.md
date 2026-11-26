# Working on gotests

## Scope and architecture

This Go module generates table-driven Go tests. The pipeline is input discovery
(`internal/input`), parsing (`internal/goparser`), models (`internal/models`),
rendering (`internal/render`), and output/import processing (`internal/output`).
`gotests.go` exposes the library; `gotests/` contains the CLI.

## Validation

- `go.mod` defines the minimum Go version. Keep the README installation minimum in sync.
- Run `bash scripts/check.sh` for formatting, version consistency, and `go vet`.
- Run `go test -count=1 -timeout=5m ./...` for deterministic tests.
- Run `go test -race -count=1 -timeout=5m ./...` when changing concurrent behavior.
- Run `node --test .github/scripts/*.test.cjs` when changing automation policy.
- Build with `go build -o /tmp/gotests ./gotests` (the repository already has a `gotests` directory).
- Live Ollama tests are opt-in: see CONTRIBUTING.md. They are not part of ordinary PR CI.

## Changes and review

- For generator fixes, add a regression that parses real source and inspects generated output.
- Consider import aliases, dot imports, receivers, generics, existing test files,
  and the distinction between source spelling and type identity.
- Keep generated Go compilable and preserve existing CLI behavior unless explicitly changing it.
- Fixtures under `testdata/` include intentionally malformed input and golden output.
  Do not globally format them. Use `scripts/regenerate-goldens.sh` when changing goldens,
  then inspect the diff; do not blindly accept output changes.
- Keep changes focused. Avoid unrelated refactors and new dependencies without a clear need.
- Use conventional PR titles (`fix:`, `feat:`, `ci:`, `docs:`, `chore:`); squash titles
  drive automated release notes. Mark breaking changes explicitly.
- AI review is advisory. Required deterministic CI must pass before merge.
- Never execute pull-request code with a privileged `pull_request_target` token.
