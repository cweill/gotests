Want to contribute? Great! First, read this page (including the small print at the end).

### Before you contribute ###
Before we can use your code, you must sign the
[Google Individual Contributor License Agreement](https://developers.google.com/open-source/cla/individual?csw=1)
(CLA), which you can do online. The CLA is necessary mainly because you own the
copyright to your changes, even after your contribution becomes part of our
codebase, so we need your permission to use and distribute your code. We also
need to be sure of various other things—for instance that you'll tell us if you
know that your code infringes on other people's patents. You don't have to sign
the CLA until after you've submitted your code for review and a member has
approved it, but you must do it before we can put your code into our codebase.
Before you start working on a larger contribution, you should get in touch with
us first through the issue tracker with your idea so that we can help out and
possibly guide you. Coordinating up front makes it much easier to avoid
frustration later on.

### Code reviews ###
All submissions, including submissions by project members, require review. We
use Github pull requests for this purpose.

### The small print ###
Contributions made by corporations are covered by a different agreement than
the one above, the Software Grant and Corporate Contributor License Agreement.

## Development and validation

Target pull requests at `develop`. Install the minimum Go version specified in
`go.mod` or a newer stable release. CI tests that exact minimum with
`GOTOOLCHAIN=local`, so automatic toolchain downloads cannot hide incompatibility.

Before submitting, run:

```sh
bash scripts/check.sh
go test -count=1 -timeout=5m ./...
go test -race -count=1 -timeout=5m ./...
node --test .github/scripts/*.test.cjs
```

Bug fixes should include a minimal source fixture and a regression test exercising
`GenerateTests`, not only a hand-built internal model. Fixtures under `testdata/`
may be intentionally invalid or unformatted. Use `scripts/regenerate-goldens.sh`
for golden updates and inspect the resulting diff.

Use a conventional PR title: `fix:` for a bug fix, `feat:` for a feature, or
`docs:`, `test:`, `ci:`, and `chore:` for maintenance. Use `!` for breaking changes.
Squash-merge titles become the input for release notes.

### Optional local AI tests

CI and release validation use deterministic tests with mocked providers. No
live-model tests run in GitHub Actions. For optional local debugging, install
Ollama and run:

```sh
ollama serve
# In another terminal:
ollama pull qwen2.5-coder:0.5b
GOTESTS_E2E=true go test -v -count=1 -tags=e2e -timeout=15m ./internal/ai
```

The local live suite has known model-dependent failures (see #197) and is not
a release requirement.

### Maintenance automation

- `CI` is the required aggregate check: minimum/stable Go tests, race tests,
  formatting, vet, and workflow validation. Greptile is an advisory reviewer.
- Dependabot groups updates weekly. Only Go patch updates that change `go.mod`
  and/or `go.sum`, without changing Go or toolchain requirements, qualify for
  auto-merge after required CI. Actions updates, minor/major updates, and
  compatibility changes require maintainer review.
- Release Please maintains a release PR on `develop`, starting from v1.9.0.
  Review its changelog and required CI, then merge it to publish the tag and
  GitHub release. Release PRs are never auto-merged.
- Release automation uses `GITHUB_TOKEN`, with explicit workflow dispatches for
  release PR validation; no personal token is needed. Actions must be allowed to
  create pull requests in repository settings.
- If a release workflow fails, rerun it or dispatch `Release Please` on `develop`.
  If release PR CI needs rerunning, dispatch `Go` on the release PR's branch.
- Codecov is the sole coverage service. Upload failures do not block test results.
- Greptile handles automatic reviews. Repository collaborators can request
  on-demand assistance with `@claude` when its OAuth secret is configured.
