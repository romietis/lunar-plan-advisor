# Contributing to Lunar Plan Advisor

Thanks for your interest in improving Lunar Plan Advisor! This is a small Go
web app that recommends the most profitable Lunar plan based on a balance.
Contributions of all sizes are welcome - bug fixes, new plan data, tests, or docs.

## Scope

The app currently supports **Lunar Denmark** only. Plan definitions (rates,
fees, caps) live in `advisor/` and `plans.json`. If Lunar's published rates
change, a PR updating these — with a link to the source — is one of the most
useful contributions.

## Prerequisites

- Go (see the version pinned in `go.mod`)
- Docker (for `make run` / `make up` and to match the CI/production image)

## Project layout

| Path                  | Purpose                                          |
|-----------------------|--------------------------------------------------|
| `advisor/`            | Core plan/profit calculation logic               |
| `cmd/web/`            | HTTP server entrypoint                           |
| `cmd/cli/`            | CLI entrypoint                                   |
| `internal/endpoints/` | HTTP handlers (`/plans`, `/plans/best`)          |
| `internal/bdd/`       | Godog (Cucumber) end-to-end tests + features     |
| `assets/`             | HTML template, CSS, JS for the web UI            |

## Development workflow

All common tasks are in the [Makefile](./Makefile):

```bash
make test    # unit tests (shuffle order, writes cover.out)
make e2e     # BDD end-to-end tests (go test -tags=e2e)
make up      # build image and run on http://localhost:8080
make stop    # stop and remove the container
make help    # full list
```

You can also run the server directly without Docker:

```bash
go run ./cmd/web
```

## Making changes

1. **Fork and branch** off `main` (e.g. `fix/negative-balance`, `feat/se-plans`).
2. **Write or update tests.** Logic changes belong with a unit test in the
   relevant package; behaviour changes to the HTTP API should have a matching
   scenario in `internal/bdd/features/`.
3. **Run `make test` and `make e2e` locally** before opening a PR.
4. **Format and vet:** `gofmt -w .` and `go vet ./...`. Code must stay
   `gofmt`-clean and Go Report Card–friendly.
5. **Keep dependencies minimal.** The app is intentionally built on the
   standard library; please discuss in an issue before adding a new dependency.

## Pull requests

- Keep PRs focused and small where possible.
- Use clear, imperative commit messages. The history follows Conventional
  Commit–style prefixes (`feat:`, `fix:`, `chore:`, `ci:`) — please match it.
- CI (`.github/workflows/test.yml`) runs unit + e2e tests and uploads coverage
  to Codecov on every PR. Please don't regress coverage.
- Describe *what* changed and *why*; link any related issue.

## Reporting bugs / requesting features

Open a GitHub issue with:
- what you expected vs. what happened,
- steps to reproduce (a `curl` against `/plans/best` is ideal for API bugs),
- the balance/plan input involved, if relevant.

## Privacy note

User plan edits are stored only in the browser's `localStorage` and are sent to
the server only as part of a single calculation request — the server stores no
per-user data. Please preserve this property: don't add server-side persistence
of user configuration without discussion.

## License

By contributing, you agree your contributions are licensed under the project's
[LICENSE](./LICENSE).
