# template

GitHub template repository for [PurgeBot-net](https://github.com/PurgeBot-net) — provides consistent CI/CD, linting, and repo hygiene across all services.

## Usage

When creating a new repository under `PurgeBot-net`, select this as the template. Then:

1. Open `.github/workflows/ci.yml` and replace both `CHANGE_ME` occurrences with the service name (e.g. `interactions`).
2. For **library repos** (no Docker image): swap `build-service.yml` → `build-library.yml` and remove the `auto-tag` and `release` jobs.

That's it. The reusable workflows live in [PurgeBot-net/workflows](https://github.com/PurgeBot-net/workflows).

## What you get

- **CI/CD** — on push to `main`: vet, test, build + push Docker image to ghcr.io, auto-tag with next patch version, create a PR in the docker repo to update the version
- **Auto-labelling** — PRs are labelled by type (bug, feature, improvement, etc.)
- **`.editorconfig`** — consistent formatting across editors
- **`.gitignore`** — Go-specific, ignores `.env` but keeps `.env.example`
- **`.golangci.yml`** — linter config with sensible defaults
- **PR template** — type checklist, testing notes, and `go.mod` reminder
- **CODEOWNERS** — auto-requests review from `@PurgeBot-net/devs`
- **SECURITY.md** — responsible disclosure policy
- **LICENSE** — MIT
