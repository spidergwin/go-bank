# Go Bank — Project Progress & Roadmap

Living engineering/learning roadmap for Go Bank (Project #1).

> **Last updated:** 2026-08-21  
> **Current version:** v5 (HTTP API — uncommitted)  
> **Go module:** `Go-bank`

---

## A. Current Project State

Go Bank is an HTTP REST API built with Go's standard library (`net/http`). It evolved from a CLI banking app through 5 versions.

### File Structure

```
go-bank/
├── main.go          → Entry point, global Accounts map, startup
├── server.go        → ServeMux setup, route registration
├── handlers.go      → HTTP handler functions (7 handlers)
├── service.go       → GetAccount helper (account lookup)
├── account.go       → BankAccount struct + methods (Deposit, Withdraw, SaveTrx)
├── transaction.go   → Transaction struct, constants, ParseTrxMsg
├── structs.go       → DTOs (request/response formats, Bank metadata)
├── storage.go       → JSON file persistence (Load, Save, FileExists)
├── accounts.json    → Persisted account data
├── Go-postman-bank-test.json → Postman collection (8 requests)
├── copy.sh          → Clipboard helper script
└── go.mod           → Module definition (Go 1.26.5)
```

### Current Routes

| Method | Route | Handler | Description |
|--------|-------|---------|-------------|
| ANY | `/` | `Welcome` | Welcome message (plain text) |
| GET | `/health` | `Healthz` | Health check |
| GET | `/about` | `About` | App metadata (JSON) |
| POST | `/accounts` | `CreateAccount` | Create new account |
| GET | `/accounts/{owner}` | `GetAccountDetails` | Get account info |
| POST | `/accounts/{owner}/deposit` | `DepositHandler` | Deposit money |
| POST | `/accounts/{owner}/withdraw` | `WithdrawalHandler` | Withdraw money |
| GET | `/accounts/{owner}/transactions` | `GetTransactions` | Transaction history |

### Data Model

- **`BankAccount`** — Owner, Balance (int), Transactions (slice)
- **`Transaction`** — Type, Status, Amount, Time, Err
- **DTOs** — `CreateAccReqFmt`, `DepositWithdrawalReqFmt`, `DepositWithdrawalRespFmt`, `Bank`
- **Constants** — `DepositTransaction`, `WithdrawalTransaction`, `Pending`, `Success`, `Failed`

### Storage

- In-memory: global `map[string]*BankAccount`
- Persistence: `accounts.json` (loaded on startup, saved after mutations)

---

## B. Completed Learning Milestones

### V1 — Fundamentals
- [x] Structs
- [x] Basic functions
- [x] Methods
- [x] Pointers
- [x] Error returns
- [x] Basic input handling
- [x] Maps

### V2 — CLI Functionality
- [x] Multi-account support
- [x] Deposit / withdraw / balance operations
- [x] CLI control flow

### V3 — Code Organization
- [x] Multi-file packages (same `package main`)
- [x] Separation of concerns across files
- [x] Git tagging and versioning

### V4 — Persistence & Error Handling
- [x] `encoding/json` (Marshal, MarshalIndent, Unmarshal)
- [x] `os` file I/O (ReadFile, WriteFile, Stat)
- [x] File existence checks
- [x] JSON persistence (load/save application state)
- [x] Structured transactions (Transaction struct, SaveTrx)
- [x] `panic`, `defer`, `recover`
- [x] Constants for transaction types/statuses

### V5 — HTTP API
- [x] `net/http` server setup
- [x] `http.NewServeMux()` routing
- [x] Path parameters (`r.PathValue`)
- [x] JSON request decoding (`json.NewDecoder`)
- [x] JSON response encoding (`json.NewEncoder`)
- [x] HTTP method validation
- [x] HTTP status codes (200, 201, 400, 404, 405, 409, 500)
- [x] Content-Type headers
- [x] Request/response DTOs with JSON struct tags
- [x] Postman collection for API testing

---

## C. Current Technical Weaknesses

### Actual Bugs

1. **Race condition on `Accounts` map** — The global map is read and written by concurrent HTTP handler goroutines with no synchronization. Under concurrent requests, Go will panic with `fatal error: concurrent map read and map write`.

2. **Port parameter ignored** — `InitiateServer(port string)` receives a port but hardcodes `:3000` in `http.ListenAndServe`.

3. **`Transaction.Err` doesn't serialize** — The `error` interface marshals to `null` or `{}` in JSON. After loading from `accounts.json`, error context is permanently lost.

4. **Failed transactions not persisted** — When deposit/withdrawal fails validation, `SaveTrx()` records the failed transaction in memory but the handler returns before calling `SaveAccounts()`. Failed transactions vanish on restart.

### Design Problems

5. **Catch-all root route** — `mux.HandleFunc("/", Welcome)` matches all unregistered paths. `/anything` returns 200 instead of 404.

6. **Inconsistent service layer usage** — `GetAccount()` exists in `service.go` but only `DepositHandler` and `WithdrawalHandler` use it. Other handlers duplicate the map lookup inline.

7. **Hardcoded save path** — `LoadAccounts` takes a `source` parameter but `SaveAccounts` hardcodes `"accounts.json"`.

8. **Wrong status on decode errors** — `DepositHandler` and `WithdrawalHandler` return 500 when JSON body decoding fails. Should be 400 (bad request from the client, not a server error).

### Learning Opportunities

9. **Dead CLI code** — `PrintInfo()`, `PrintBalance()`, `ShowTransactions()` in `account.go` are leftovers from the CLI versions. Unused by the HTTP server.

10. **No graceful shutdown** — Server starts with `log.Fatalln(InitiateServer(...))` and has no signal handling for clean shutdown.

11. **`ParseTrxMsg` string manipulation** — Uses `strings.Replace` to capitalize first letter. Could use `strings.Title` or manual byte manipulation, but this is minor and works.

### Optional Improvements (Future)

- Middleware for logging, panic recovery, Content-Type
- Input validation beyond just amount > 0
- Consistent error response format (JSON instead of plain text)
- Structured logging (`log/slog`)

---

## D. Recommended Next Step

### **Fix the bugs in v5, then commit/tag it**

Before moving to new concepts, the current code has real bugs that are worth fixing as a learning exercise. This isn't busywork — each fix teaches something:

| Bug | What You Learn |
|-----|----------------|
| Port parameter ignored | Attention to function signatures, `fmt.Sprintf` for string formatting |
| Catch-all root route | How ServeMux pattern matching works in Go |
| Wrong status on decode errors | HTTP semantics (4xx vs 5xx) |
| Failed transactions not saved | Thinking about data consistency |
| Dead CLI code | Code cleanup discipline when evolving a project |

**Skip the race condition for now** — that's a concurrency concept. You'll address it properly when you learn `sync.Mutex`, which should come soon but not before understanding the problem.

**Skip the `Transaction.Err` serialization issue for now** — that's actually an interface/custom marshaling problem. It'll be a great exercise when you learn interfaces.

#### Steps
1. Fix the 4-5 simpler bugs listed above
2. Clean up dead CLI methods
3. Commit and tag as `v5.0`
4. Then we discuss what v5.1 or v6 should introduce

---

## E. Future Roadmap

```
V1  ── Language fundamentals, CLI                          ✓ tagged
V2  ── Interactive CLI, multi-account                      ✓ tagged
V3  ── Multi-file organization, Git versioning             ✓ tagged
V4  ── JSON persistence, transactions, panic/defer/recover ✓ tagged
V4.1 ── Transactions as structs                            ✓ tagged
V5  ── HTTP API with net/http                              ← current (uncommitted)
V5.x ── Bug fixes, cleanup, consistent error responses
V6  ── Service layer, interfaces, architecture cleanup
V7  ── Testing (unit tests, table-driven, httptest)
V8  ── Concurrency (mutex, race detection, safe map access)
V9  ── Database persistence (SQLite or PostgreSQL)
V10 ── Authentication / authorization
...
Future ── USSD service (reusing bank business logic)
```

*Version numbers are approximate. The roadmap evolves based on what you actually implement.*

---

## F. Version History

| Version | Tag | Commit | Summary |
|---------|-----|--------|---------|
| V1 | `v1.0` | `70e390e` | Initial bank application |
| V2 | `v2.0` | `f3ad591` | Interactive CLI |
| V3 | `v3.0` | `dc6c55c` | Modular multi-account CLI |
| V4 | `v4.0` | `4d9f792` | JSON persistence |
| V4.1 | `v4.1` | `ff00319` | Transactions as structs |
| V5 | — | uncommitted | HTTP REST API |
