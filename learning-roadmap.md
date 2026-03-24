# Go Backend Development Learning Roadmap

## Goal

Build real Go skills that transfer to **production and open-source work**—not to tick every box in the language spec.

---

## 80/20 accelerated path

Most day-to-day Go (APIs, CLIs, services) reuses a **small core**: types, errors, JSON, `net/http`, tests, and reading others’ code. Everything else—channels, advanced concurrency, profiling, deployment—is easier **in context** of a real repo.

**This roadmap has two tracks:**

1. **Pre-OSS essentials** — short list to finish **before** (or in parallel with) your first serious OSS contribution.
2. **On the project** — learn when you hit it; no blocking “finish the book” phase.

---

## Foundation (completed)

### Phase 1: Go fundamentals
- [x] Getting started, modules, `go run` / `go build`
- [x] Syntax, types, control flow, functions, packages

### Phase 2: Go-specific features
- [x] Pointers, structs, slices, maps, interfaces

### Phase 3: Errors & concurrency (baseline)
- [x] **3.1** Error handling (`error`, `if err != nil`, `fmt.Errorf` / `%w`, `errors.Is` when needed)
- [x] **3.2** Goroutines (`go`, lifecycle, why not one OS thread per task)

**Deferred (learn on the project or when needed):** deep dive on channels, `select`, worker pools, `sync.Mutex` / `WaitGroup`, atomics—these matter for specific codepaths, not for “must finish before coding.”

---

## Pre-OSS essentials (do these next)

Tight list so you can move to **your open-source project** quickly. Check these off, then ship learning through real issues and reviews.

### A. JSON (`encoding/json`) — high priority
- [ ] `json.Marshal` / `json.Unmarshal` (and `Decoder` for streams if the API reads large bodies)
- [ ] Struct tags: `` `json:"field_name"` ``, `omitempty`, ignoring fields
- [ ] Nested structs and `[]` / `map` in JSON

*Why:* Almost every HTTP handler speaks JSON. Custom marshalers can wait until you need them.

### B. HTTP (`net/http`) — high priority
- [ ] `http.Handler` / `HandlerFunc`; registering routes on `ServeMux` (or the router the project uses)
- [ ] Methods: GET, POST, etc.; status codes; headers
- [ ] Read request body (`io.ReadAll` / `json.NewDecoder(r.Body)`); write response (`w.Write`, `json.NewEncoder(w)`)
- [ ] **Optional but common in OSS:** `http.Client` for outbound calls (timeouts, `context`)

### C. `context` — medium priority (small investment, big payoff)
- [ ] `context.Context` on request paths: cancellation, deadlines (`context.WithTimeout`)
- [ ] Passing `ctx` into functions that do I/O (matches real codebases)

*Why:* stdlib and most libraries thread `context`; you’ll read it everywhere in OSS.

### D. Testing — medium priority (enough to contribute safely)
- [ ] `*_test.go`, `func TestXxx(t *testing.T)`
- [ ] Table-driven tests (one pattern covers many cases)
- [ ] `httptest` for handler tests when the project is HTTP-heavy

*Skip for now:* fuzzing, benchmarks, heavy mocking frameworks—add when the repo demands it.

### E. Project hygiene (quick)
- [ ] `go fmt`, `go vet`, and running `go test ./...` before a PR
- [ ] Reading `go.mod` / `go.sum` and adding deps with `go get`

---

## On the OSS project (learn alongside)

Pick these up **when the issue touches them**—no need to front-load.

| Topic | When to learn |
|--------|----------------|
| Channels, `select`, pipelines | Goroutine coordination, streaming, worker code |
| `sync.Mutex`, `RWMutex`, `WaitGroup` | Shared state, parallel tests, worker pools |
| `database/sql`, drivers, `sqlx` / ORM | Storage layer, migrations |
| Middleware, JWT, cookies | Auth and cross-cutting HTTP behavior |
| `os`, `filepath`, config files | CLI, config loading, file watchers |
| Structured logging (`slog`, `zap`, etc.) | Observability in that repo |
| Docker, health checks, graceful shutdown | Deployment and ops issues |
| `pprof`, profiling | Performance issues |

---

## Open-source focus

After **Pre-OSS essentials (A–E)**, prioritize:

1. **Clone the repo**, build with `go build`, run tests, read `CONTRIBUTING.md`.
2. **Small PRs**: docs, tests, bugfixes—read existing patterns; match style.
3. **Use the issue tracker** as your curriculum: each ticket teaches one slice of the stack.

Your “final project” is no longer a toy API in isolation—it’s **real usage** in a codebase you care about. Optionally keep a tiny side API if you want a sandbox; it’s not a gate to OSS.

---

## Reference: fuller curriculum (optional later)

If you ever want depth beyond the 80/20 path, topics map roughly to: full JSON customization, file I/O and strings, advanced routing and middleware, DB integration, validation, auth, production logging, graceful shutdown, deployment, and profiling. Treat these as a **menu**, not a prerequisite.

---

## Learning resources

### Official
- [Go Tour](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)

### Practice
- [Exercism Go Track](https://exercism.org/tracks/go)

---

## Tips for JS/Python developers

1. **Errors are values** — check and return; wrapping with `%w` preserves cause for `errors.Is`.
2. **Interfaces are implicit** — implement methods; no `implements` keyword.
3. **Concurrency** — goroutines are cheap; coordination (channels, mutexes) is chosen per problem—don’t force channels everywhere.
4. **Formatting** — `go fmt` is the style guide; run it before commit.

---

## Progress tracker

### Foundation
- [x] Phase 1 — Fundamentals
- [x] Phase 2 — Pointers, structs, slices, maps, interfaces
- [x] Phase 3.1 — Error handling
- [x] Phase 3.2 — Goroutines

### Pre-OSS essentials
- [ ] A — JSON
- [ ] B — `net/http`
- [ ] C — `context`
- [ ] D — Testing basics (+ `httptest` if HTTP-heavy)
- [ ] E — `fmt` / `vet` / `go test ./...` / modules

### Open-source project
- [ ] First successful build + test run
- [ ] First merged contribution (docs, test, or fix)

---

## After this path

gRPC, WebSockets, Kubernetes, and advanced performance work fit naturally once you’re productive in a real repo—not as a second “degree” before you contribute.

Learning is iterative: ship small changes, read reviews, repeat.
