<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Backend-01 / html css handling / cookie

## Overview

This directory explores **GO-Backend-01 / html css handling / cookie** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Backend-01 / html css handling / cookie in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/html css handling/cookie/main.go) | Go | CORS with credentials cannot use "*" for Allow-Origin — the browser rejects it. So we echo back whatever Origin the request came from. |

## Concepts

### 1. Main (`main.go`)

File: [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/html css handling/cookie/main.go)  
CORS with credentials cannot use "*" for Allow-Origin — the browser rejects it. So we echo back whatever Origin the request came from..

Relevant code excerpt:

```go
func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "testcookie",
		Value:    "hello",
```

**Explanation**:
- Implements function(s): `withCORS()`, `setCookieHandler()`, `checkCookieHandler()`, `main()`.
- Defines C routine(s): `func()`, `setCookieHandler()`, `checkCookieHandler()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. Entry into [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/html css handling/cookie/main.go) (`main` or top-level routine).
2. Initializes internal state / variables and executes core logic.
3. Processes inputs, performs transformations or system calls, and outputs results.

```text
main() [main.go]
  │
  ├── Initialize state / data
  ├── Execute core operations
  └── Output result / Terminate
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run main.go
```

## Important Details

- **Memory & Scope**: Notice variable allocation, pointer indirection, and lifetime across function boundaries.
- **Error & Return Handling**: Always verify return values and errors before proceeding to prevent panics or undefined behavior.
- **Resource Management**: Check that open files, network sockets, database handles, and heap allocations are properly closed or freed.
- **AMQP Exchange Routing**: RabbitMQ directs messages via direct, topic (wildcard `*`, `#`), or fanout exchange bindings with publisher confirms.
- **Server-Sent Events (SSE)**: Unidirectional HTTP streaming with `Content-Type: text/event-stream`, utilizing `http.Flusher` to push instant updates to browser clients.
- **PostgreSQL Transactional Queues**: Implements persistent message queuing using `SELECT ... FOR UPDATE SKIP LOCKED` or advisory locks to prevent concurrent worker collisions.

## Common Mistakes

- Assuming implicit synchronization or thread safety where none is provided.
- Ignoring error return values or missing zero-value edge cases.
- Misunderstanding pass-by-value versus pointer/reference semantics for complex types.
- Failing to cast `http.ResponseWriter` to `http.Flusher` in SSE endpoints before attempting to stream data chunks.
- Leaving database transactions open or omitting rollback on error paths, holding row locks indefinitely.

## Language Notes

- **Language Features**: Written using idiomatic Go paradigms.
- **Go Web Stack**: `net/http` handles each inbound connection in an isolated lightweight goroutine.

## Related Concepts

- Data Structures & Algorithms in Go
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `main.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `main.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/html css handling/cookie/main.go)
