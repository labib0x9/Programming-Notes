<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Backend-01 / Protocols / websocket

## Overview

This directory explores **GO-Backend-01 / Protocols / websocket** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Backend-01 / Protocols / websocket in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`json_msg.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Protocols/websocket/json_msg.go) | Go | Implements `Increment, handleConn` logic for json msg |
| [`text_binary_frame.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Protocols/websocket/text_binary_frame.go) | Go | Implements `handleConn` logic for text binary frame |

## Concepts

### 1. Json Msg (`json_msg.go`)

File: [`json_msg.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Protocols/websocket/json_msg.go)  
Implements `Increment, handleConn` logic for json msg.

Relevant code excerpt:

```go
type A struct {
	Id int `json:"req_id"`
}

func (a *A) Increment() {
	a.Id++
}

func handleConn(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, http.Header{})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(r.Context(), 40*time.Second)
	defer cancel()

	// cacellation - Method A
	if deadline, ok := ctx.Deadline(); ok {
		_ = conn.SetReadDeadline(deadline)
```

**Explanation**:
- Implements function(s): `Increment()`, `handleConn()`, `main()`.
- Defines C routine(s): `handleConn()`, `main()`.
- Defines Go struct(s): `A`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Text Binary Frame (`text_binary_frame.go`)

File: [`text_binary_frame.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Protocols/websocket/text_binary_frame.go)  
Implements `handleConn` logic for text binary frame.

Relevant code excerpt:

```go
func handleConn(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, http.Header{})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(r.Context(), 40*time.Second)
	defer cancel()

	// cacellation - Method A
	if deadline, ok := ctx.Deadline(); ok {
		_ = conn.SetReadDeadline(deadline)
	}

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			// what to do?
			break
		}
```

**Explanation**:
- Implements function(s): `handleConn()`, `main()`.
- Defines C routine(s): `handleConn()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`json_msg.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Protocols/websocket/json_msg.go) provides or tests `json_msg`.
2. [`text_binary_frame.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Protocols/websocket/text_binary_frame.go) provides or tests `text_binary_frame`.

```text
Caller / Test Runner
  ├──> [json_msg.go] (Implements `Increment, handleConn` ...)
  ├──> [text_binary_frame.go] (Implements `handleConn` logic for t...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run json_msg.go
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

1. What is the primary role of the functions/routines demonstrated in `json_msg.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `json_msg.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`json_msg.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Protocols/websocket/json_msg.go)
- [`text_binary_frame.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Protocols/websocket/text_binary_frame.go)
