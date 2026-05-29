# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Run

```bash
go install .                # build and install the binary
http-spec path/to/*.htsf    # run spec files against a live HTTP endpoint
```

This is a single `package main` Go module with zero external dependencies. All source files live in the root directory.

There are no unit tests (`_test.go`). Testing is done by running `.htsf` spec files against live HTTP endpoints. The `./validate` script runs the full test suite inside Docker. The `./run-http-specs` script orchestrates example spec execution with expected pass/fail outcomes.

## Architecture

**10-stage numbered pipeline** — each file (`00-main.go` through `10-result-gatherer.go`) is one pipeline stage. Each stage is a function that takes `*context` and calls the next stage directly:

| Stage | File | Role |
|-------|------|------|
| 00 | `00-main.go` | CLI flags, context init, launch result gatherer |
| 01 | `01-spec-file-scatterer.go` | Fan out: one goroutine per `.htsf` file |
| 02 | `02-spec-file-processor.go` | Open file, create HTTP client, init substitutions |
| 03 | `03-spec-triplet-iterator.go` | Loop: read request/response pairs from file |
| 04 | `04-desired-request-substituter.go` | Apply `⧈name⧈` substitutions to request |
| 05 | `05-expected-response-match-parser.go` | Compile `⧆name⧆regexp⧆` matchers in response |
| 06 | `06-expected-response-substituter.go` | Apply substitutions to expected response |
| 07 | `07-desired-request-sender.go` | Send HTTP request with retry logic |
| 08 | `08-actual-response-translator.go` | Parse HTTP response into internal format |
| 09 | `09-response-comparator.go` | Line-by-line comparison, capture regexp groups |
| 10 | `10-result-gatherer.go` | Collect results, print summary, set exit code |

**Concurrency model**: files are processed concurrently (one goroutine each via stage 01); requests within a file are sequential. Results flow to stage 10 via an unbuffered `chan context`.

## Key Types

All in `type-*.go` files:

- **`context`** — mutable state threaded through the pipeline (HTTP client, substitutions map, error, waitgroup, result channel)
- **`line`** — single parsed HTSF line; `IOPrefix` determines type: `>` request, `<` response, `#` comment, `+` sleep
- **`message`** — HTTP message structure (first line, headers, blank line, body); shared by request and response
- **`request`** / **`response`** — extend `message` with scheme/hostname or status parsing
- **`specTriplet`** — pairs a desired request with its expected response

## HTSF Format

Spec files use HTSF (Hypertext Specification Format), modeled after `curl -v` output:

- `> ` prefix for request lines, `< ` prefix for response lines (bare `>` or `<` for the mandatory blank line)
- Blank line separates request from expected response
- Response headers **must be sorted in ASCII order**
- `# ` for comments, `+ duration` for delays between requests (e.g., `+ 1.5s`)
- **Regexp matchers**: `⧆optional-name⧆mandatory-regexp⧆` (U+29C6 SQUARED ASTERISK) — captures stored in substitutions map
- **Substitutions**: `⧈name⧈` (U+29C8 SQUARED SQUARE) — replaced with previously captured values
- **Built-in matchers**: `:date`, `:uuid`, `:b62:22`, `:iso8601:µs:z`
- **Built-in substitution**: `YYYY-MM-DD` (today's date)

## Coding Philosophy

See the Coding Philosophy section of README.md for the upstream scripture this codebase embodies — key examples:

- **Context as single vessel** — `*context` carries all state through the pipeline; no scattered arguments
- **Process-first naming** — pipeline stages use `-er` suffix (`scatterer`, `substituter`, `comparator`, `gatherer`)
- **Platform primitives only** — zero external dependencies, stdlib only
- **Communicating Sequential Processes** — goroutines communicate via channels, never shared mutable state
- **Shallow structure** — all source files flat in root directory
- **Validate at edges, trust within** — lines validated on construction (`line.validate()`), trusted downstream

## Conventions

- Errors propagate via `context.Err` field; `errorHandler()` returns bool for early exit
- HTTP client does not follow redirects (`http.ErrUseLastResponse`)
- Request-only mode (no expected response) outputs the actual response formatted as HTSF, reports as failure
- Commit messages: single line, ~50 characters, completing "When applied, this commit will ___"
