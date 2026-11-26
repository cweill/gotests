#!/usr/bin/env bash
# Fast deterministic checks; malformed/generated fixtures are deliberately excluded from gofmt.
set -euo pipefail
cd "$(dirname "$0")/.."
minimum=$(awk '$1 == "go" { print $2 }' go.mod)
grep -Fqx "__Minimum Go version:__ Go $minimum" README.md || {
  echo "README minimum Go version must match go.mod ($minimum)" >&2
  exit 1
}
failed=0
while IFS= read -r -d '' file; do
  case "$file" in testdata/*|*/testdata/*) continue ;; esac
  if [ -n "$(gofmt -l "$file")" ]; then
    echo "Run gofmt -w $file" >&2
    failed=1
  fi
done < <(git ls-files -z '*.go')
[ "$failed" -eq 0 ]
go vet ./...
git diff --exit-code -- go.mod go.sum
