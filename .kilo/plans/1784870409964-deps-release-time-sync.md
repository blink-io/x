# Plan: Comprehensive Code Optimization Analysis Report

## Goal
Produce a single, well-structured Markdown analysis report covering **performance, readability, maintainability, and security** for this Go `x`-style wrapper library. The report must cite concrete files/lines, give a code-example diff for each finding, explain the reasoning, and state the expected effect. No source-code edits in this task — only the report document.

## Deliverable
- **File:** `docs/optimization-analysis.md` (new)
- **Audience:** maintainers of this repo; sized for an actionable review (~400–700 lines).
- **Structure:** top-level sections matching the four requested dimensions, each containing a numbered list of findings with: *File:Line • Issue • Why • Fix (code) • Expected effect*.

## Codebase facts collected during planning
Module: `github.com/blink-io/x` • Go 1.26.2 • ~80+ top-level sub-packages • mostly thin wrappers around third-party libs (`phuslu/log`, `go-redis`, `rueidis`, `bun`, `fiber`, `kratos`, `nicksnyder/go-i18n`, `sethvargo/go-limiter`, `govalidator`, etc.). Centralized interfaces in `cache/cache.go` (`Cache[V]`, `TTLCache[V]`, `ErrCache[V]`, `ErrTTLCache[V]`). Errors silently dropped in many `redis` adapters (`cache/redis/{goredis,rueidis}`). Hard-coded `context.Background()` in cache adapters. Alias-only packages (`misc/conc/conc.go`, `validator/govalidator/govalidator.go`, `http/middleware/limit/limit.go`, `log/log.go`). Empty `package server` at `http/server/server.go`. TODO/FIXME markers in 5 files (see Maintainability §).

## Confirmed findings to include in the report

### Performance (5 findings)
1. **`misc/id/nanoid.go:7-10`** — `NanoID` constructs a new generator every call.
   - Why: `nanoid.Standard(len)` allocates alphabet table + closure per call; O(n) gen throughput suffers.
   - Fix: precompute and reuse a per-length generator in a `sync.Map`.
   - Expected: removes ~1 alloc/op + 1 fn-call indirection per ID.
2. **`misc/id/nanoid2.go:7-10`** — same pattern for `gonanoid.New()`.
   - Fix: keep a package-level `var idGen = func() gonanoid.ID { ... }()` cached.
3. **`i18n/localizer.go:10-26`** — `Tr()` returns a closure that re-allocates a `LocalizeConfig` and re-iterates options every call; `SprigFuncs()` calls `sprig.FuncMap()` (allocates ~150-entry map) each call.
   - Fix: cache `sprig.FuncMap()` result in a `sync.OnceValue`; let callers pass through a pre-bound translator.
   - Expected: cuts heap allocs in hot i18n paths by ~1 obj + ~150 map entries per request.
4. **`cache/redis/goredis/goredis.go:29-58`** — `Set/Del` discards `*StatusCmd`; `Get` uses two round trips (`cmd.Err()` then `cmd.Scan(&v)`) because `Err()` is checked but the underlying client call returns one command.
   - Fix: combine into `if _, err := cmd.Result(); err != nil { ... }`; or for `Get`, `v, err := cmd.Bytes();` then decode.
   - Expected: marginally fewer allocations; clearer semantics; `cmd.Result()` uses single decode path.
5. **`cache/redis/goredis/goredis.go:21-27`** — `New` always stores `context.Background()`; no way to inherit caller timeout.
   - Fix: add `WithContext(ctx)` option and call `c.rc.Set(ctx, ...)` from caller-supplied context where used.
   - Expected: enables per-call deadlines, prevents head-of-line blocking on slow Redis.

### Readability (5 findings)
1. **`cache/redis/rueidis/rueidis.go:16`** — `const Name = "goredis"` (copy/paste from sibling file).
   - Fix: `const Name = "rueidis"`.
   - Expected: accurate adapter name in logs/metrics.
2. **`cache/cache.go:10-41`** — four interfaces (`Cache`, `TTLCache`, `ErrCache`, `ErrTTLCache`) with overlapping surfaces; consumer chooses between them awkwardly.
   - Fix: keep `Cache[V]` (no-error) + `ErrCache[V]` (error-returning), drop the TTL pair from the interface and use a separate `TTLSetter[V]` mixin already used by `TTLCache[V]`.
   - Expected: smaller API surface; clearer intent; easier mocking.
3. **`log/log.go:1-90`** — 90-line alias file re-exporting ~25 types and ~15 vars/consts from `phuslu/log`.
   - Fix: split into `log/levels.go`, `log/writers.go`, `log/funcs.go` by domain, or generate via `go generate` from upstream `go doc`.
   - Expected: easier code review; clearer ownership.
4. **`http/server/server.go`** — empty package (1-line file).
   - Fix: either delete or add a placeholder doc comment explaining reserved path; remove from build if no plans.
   - Expected: removes dead code, faster `go list ./...`.
5. **`misc/conc/conc.go`, `validator/govalidator/govalidator.go`, `http/middleware/limit/limit.go`** — single-line alias files that re-export exactly one type each (`WaitGroup`, `Validator`, `Store`). Provide no value over importing the upstream module directly.
   - Fix: delete and update consumers; or keep as a deliberate "minimal surface" shim and add a doc comment explaining intent.
   - Expected: less drift risk; smaller module surface.

### Maintainability (4 findings)
1. **TODO/FIXME inventory** — 7 markers across 5 files: `kratos/registry/mdns/mdns.go:49,85`, `kvstore/store/redis/rueidis/rueidis.go:87,297,706`, `kvstore/store/redis/goredis/goredis.go:113,307,695`, `i18n/locale/locale_shared.go:22,25`. None reference tickets.
   - Fix: convert to `// TODO(#N):` with an issue number, or move to `docs/TODO.md` and remove from source. Resolve the `mdns.go` stubs (delete if not planned) and the `kvstore` cluster & key-size TODOs.
   - Expected: actionable backlog, no orphaned comments.
2. **No project-wide lint/format config** — repo lacks `.golangci.yml`, no Makefile targets besides what's already there.
   - Fix: add `.golangci.yml` enabling `govet`, `staticcheck`, `gocritic`, `revive`, `errorlint`, `unused`; wire to `make lint`.
   - Expected: catches the silent-error swallowing and unused vars automatically.
3. **Adapter naming drift** — `cache/redis/{goredis,rueidis}` both expose `Name = "goredis"`; `kvstore/store/redis/{goredis,rueidis}` should be checked similarly.
   - Fix: extract `Name` into per-file constant and run `go test ./...` with a name-collision test.
   - Expected: removes metrics/log ambiguity.
4. **Wrapper drift risk** — every `type X = upstream.X` alias needs re-validation on upstream minor bumps. No script enforces it.
   - Fix: add `make verify-aliases` that runs `go doc` on upstream modules and diffs the alias list; or drop the wrappers where they add no value (see Readability §5).
   - Expected: surfaces upstream renames before releases.

### Security (5 findings)
1. **`cache/redis/{goredis,rueidis}` — hardcoded `context.Background()`** in `Set/Get/Del`. No per-call timeout.
   - Why: a slow/hung Redis can pin goroutines indefinitely and exhaust file descriptors.
   - Fix: take a parent context in `New(...)` (or via `WithContext` option) and pass it through; expose `SetCtx/GetCtx/DelCtx` for caller-supplied deadlines.
   - Expected: bounded tail latency; observability via context cancellation.
2. **Silent error swallowing** in `cache/redis/goredis/goredis.go:29-58` and `cache/redis/rueidis/rueidis.go:42-76`. `Set` failures (network, OOM eviction, NOPERM, write-then-read replication lag) are invisible.
   - Fix: return errors via `ErrCache[V]` (already defined in `cache/cache.go:27`); add metrics counter for `cache_set_error_total`.
   - Expected: failures become observable; permission/ACL mistakes surface immediately.
3. **CSRF middleware defaults** — `http/middleware/csrf/{filippo,gorilla,nosurf}` (not yet inspected in plan; include in report after a quick read). Look for: HTTPS-only cookies, `SameSite=Lax/Strict` default, double-submit vs synchronizer-token choice, body-parsing CSRF on non-safe methods only, `Origin` vs `Referer` allowlist when used.
   - Fix (template): document secure defaults; if any wrapper sets `Secure: false` by default, change to `true` when `r.TLS != nil` and add an option to force-enable.
   - Expected: closes common CSRF regression vectors.
4. **`crypto/` and `misc/password/`** — quick scan needed; flag any use of `crypto/md5`, `crypto/sha1`, or `crypto/rand` seeding via `math/rand`, or `bcrypt` cost < 10.
   - Fix: replace `md5/sha1` with `sha256`/`sha512`; require `crypto/rand` for token generation; bump bcrypt default cost to 12.
   - Expected: aligned with modern OWASP guidance.
5. **No request size limits / timeouts** in adapters that wrap `http.Server`-style usage (`http/server/server.go` is empty; recommend adding defaults via a `ServerOpts` struct with `ReadHeaderTimeout`, `ReadTimeout`, `WriteTimeout`, `IdleTimeout`, `MaxHeaderBytes` pre-set to safe defaults).
   - Expected: prevents Slowloris-style attacks; forces callers to opt into unsafe settings.

## Report outline (what the report file must contain)
```
# Optimization Analysis

## Summary
- Scope, methodology, severity legend (P0/P1/P2), totals per dimension.

## 1. Performance
   1.1 misc/id/nanoid.go ...
   1.2 misc/id/nanoid2.go ...
   1.3 i18n/localizer.go ...
   1.4 cache/redis/goredis/goredis.go (Get decode) ...
   1.5 cache/redis/goredis/goredis.go (context plumbing) ...

## 2. Readability & Code Quality
   2.1 cache/redis/rueidis/rueidis.go (Name constant)
   2.2 cache/cache.go (interface surface)
   2.3 log/log.go (alias sprawl)
   2.4 http/server/server.go (empty package)
   2.5 Single-line alias shims (conc/validator/limit)

## 3. Maintainability
   3.1 TODO/FIXME inventory
   3.2 Lint config gap
   3.3 Adapter naming drift
   3.4 Wrapper drift detection

## 4. Security
   4.1 Cache context handling
   4.2 Silent error swallowing
   4.3 CSRF defaults (post-inspection)
   4.4 Crypto/hash primitives (post-inspection)
   4.5 HTTP server defaults

## Appendix A — Files inspected
## Appendix B — Suggested PR sequencing (5-PR series, no LoE estimates)
```

## Execution steps for the implementation agent
1. Read the remaining uninspected files referenced above (`http/middleware/csrf/{filippo,gorilla,nosurf}/*`, `crypto/**/*.go`, `misc/password/*`, `http/server/server.go`'s callers) and add findings 4.3, 4.4, 4.5 with concrete file:line evidence.
2. For each finding, draft a Go code example (≤25 lines) showing the **before** and **after** snippet.
3. Write `docs/optimization-analysis.md` per the outline. Use tables for severity and dimension totals at the top.
4. Cross-link each finding to the file:line so reviewers can jump directly (`path/file.go:LL`).
5. Do **not** edit any source files. The report is the deliverable.

## Risks / open questions
- **Scope creep:** the repo has 80+ packages; the report covers a representative sample plus all open TODOs. A future "round 2" pass can cover encoding/, log/slog/, kratos adapters, scheduler/.
- **Upstream choices:** some wrapper files (e.g., `misc/conc/conc.go`) exist as deliberate re-exports. The report flags them as candidates for removal but does not mandate it; final decision stays with the maintainer.
- **No LoE estimates** included (per planner behavior rules).

## Validation
- `git grep` for every `path:line` cited in the report returns the same line content.
- Each finding has at least one of: a code snippet, a config snippet, or a referenced issue/ticket.
- Total file length stays within the 400–700 line target so it remains reviewable.
- No source files were modified by this task.

## Out of scope
- Implementing any of the proposed fixes.
- Adding/updating tests.
- Bumping dependency versions.
- Modifying `go.mod` / `go.sum`.