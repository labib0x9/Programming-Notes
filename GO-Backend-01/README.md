<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Backend-01

## Overview

This directory explores **GO-Backend-01** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Backend-01 in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`basic-cors-handling.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/basic-cors-handling.go) | Go | Implements `init, getAllTopics` logic for basic-cors-handling |
| [`dump_request.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/dump_request.go) | Go | Implements `Get` logic for dump request |
| [`global_router.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/global_router.go) | Go | Implements `init, getAllTopics, getTopics` logic for global router |
| [`handle_vs_handleFunc.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/handle_vs_handleFunc.go) | Go | Implements `init, getAllTopics, getTopics` logic for handle vs handlefunc |
| [`middleware.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/middleware.go) | Go | Implements `NewManager, Use, With` logic for middleware |

## Concepts

### 1. Basic-Cors-Handling (`basic-cors-handling.go`)

File: [`basic-cors-handling.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/basic-cors-handling.go)  
Implements `init, getAllTopics` logic for basic-cors-handling.

Relevant code excerpt:

```go
type Topic struct {
	PageUrl          string `json:"page_url"`
	IconUrl          string `json:"icon_url"`
	Name             string `json:"name"`
	ShortDescription string `json:"desp"`
	AHref            string `json:"a_href"`
}

var AllTopics []Topic

func init() {
	linux := Topic{
		PageUrl:          "/linux",
		IconUrl:          "/static/linux-plain.svg",
		Name:             "Linux Environments",
		ShortDescription: "Deploy isolated, ephemeral Linux containers directly from your browser. Practice privilege escalation, system administration, and network forensics securely.",
		AHref:            "Explore Linux Labs →",
	}

	ccode := Topic{
		PageUrl:          "/cprog",
		IconUrl:          "/static/c-original.svg",
```

**Explanation**:
- Implements function(s): `init()`, `getAllTopics()`, `main()`.
- Defines C routine(s): `init()`, `getAllTopics()`, `main()`.
- Defines Go struct(s): `Topic`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Dump Request (`dump_request.go`)

File: [`dump_request.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/dump_request.go)  
Implements `Get` logic for dump request.

Relevant code excerpt:

```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func Get(w http.ResponseWriter, r *http.Request) {
	rawReq, err := httputil.DumpRequest(r, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(rawReq))
}

func main() {
	// --- //
}
```

**Explanation**:
- Implements function(s): `Get()`, `main()`.
- Defines C routine(s): `Get()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Global Router (`global_router.go`)

File: [`global_router.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/global_router.go)  
Implements `init, getAllTopics, getTopics` logic for global router.

Relevant code excerpt:

```go
type Topic struct {
	PageUrl          string `json:"page_url"`
	IconUrl          string `json:"icon_url"`
	Name             string `json:"name"`
	ShortDescription string `json:"desp"`
	AHref            string `json:"a_href"`
}

var AllTopics []Topic

func init() {
	linux := Topic{
		PageUrl:          "/linux",
		IconUrl:          "/static/linux-plain.svg",
		Name:             "Linux Environments",
		ShortDescription: "Deploy isolated, ephemeral Linux containers directly from your browser. Practice privilege escalation, system administration, and network forensics securely.",
		AHref:            "Explore Linux Labs →",
	}

	ccode := Topic{
		PageUrl:          "/cprog",
		IconUrl:          "/static/c-original.svg",
```

**Explanation**:
- Implements function(s): `init()`, `getAllTopics()`, `getTopics()`, `GlobalRouter()`, `main()`.
- Defines C routine(s): `init()`, `getAllTopics()`, `getTopics()`, `main()`.
- Defines Go struct(s): `Topic`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Handle Vs Handlefunc (`handle_vs_handleFunc.go`)

File: [`handle_vs_handleFunc.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/handle_vs_handleFunc.go)  
Implements `init, getAllTopics, getTopics` logic for handle vs handlefunc.

Relevant code excerpt:

```go
type Topic struct {
	PageUrl          string `json:"page_url"`
	IconUrl          string `json:"icon_url"`
	Name             string `json:"name"`
	ShortDescription string `json:"desp"`
	AHref            string `json:"a_href"`
}

var AllTopics []Topic

func init() {
	linux := Topic{
		PageUrl:          "/linux",
		IconUrl:          "/static/linux-plain.svg",
		Name:             "Linux Environments",
		ShortDescription: "Deploy isolated, ephemeral Linux containers directly from your browser. Practice privilege escalation, system administration, and network forensics securely.",
		AHref:            "Explore Linux Labs →",
	}

	ccode := Topic{
		PageUrl:          "/cprog",
		IconUrl:          "/static/c-original.svg",
```

**Explanation**:
- Implements function(s): `init()`, `getAllTopics()`, `getTopics()`, `main()`.
- Defines C routine(s): `init()`, `getAllTopics()`, `getTopics()`, `main()`.
- Defines Go struct(s): `Topic`.
- Demonstrates step-by-step logic and runtime behavior.

### 5. Middleware (`middleware.go`)

File: [`middleware.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/middleware.go)  
Implements `NewManager, Use, With` logic for middleware.

Relevant code excerpt:

```go
type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (m *Manager) Use(middlewares ...Middleware) {
	m.globalMiddlewares = append(m.globalMiddlewares, middlewares...)
}

func (m *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		next = middleware(next)
	}

	for _, middleware := range m.globalMiddlewares {
```

**Explanation**:
- Implements function(s): `NewManager()`, `Use()`, `With()`, `Logger()`, `Info()`, `GetAllNotes()`, `main()`.
- Defines C routine(s): `GetAllNotes()`, `main()`.
- Defines Go struct(s): `Manager`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`basic-cors-handling.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/basic-cors-handling.go) provides or tests `basic-cors-handling`.
2. [`dump_request.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/dump_request.go) provides or tests `dump_request`.
3. [`global_router.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/global_router.go) provides or tests `global_router`.
4. [`handle_vs_handleFunc.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/handle_vs_handleFunc.go) provides or tests `handle_vs_handleFunc`.

```text
Caller / Test Runner
  ├──> [basic-cors-handling.go] (Implements `init, getAllTopics` log...)
  ├──> [dump_request.go] (Implements `Get` logic for dump req...)
  ├──> [global_router.go] (Implements `init, getAllTopics, get...)
  ├──> [handle_vs_handleFunc.go] (Implements `init, getAllTopics, get...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run basic-cors-handling.go
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

1. What is the primary role of the functions/routines demonstrated in `basic-cors-handling.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `basic-cors-handling.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`basic-cors-handling.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/basic-cors-handling.go)
- [`dump_request.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/dump_request.go)
- [`global_router.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/global_router.go)
- [`handle_vs_handleFunc.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/handle_vs_handleFunc.go)
- [`middleware.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/middleware.go)
