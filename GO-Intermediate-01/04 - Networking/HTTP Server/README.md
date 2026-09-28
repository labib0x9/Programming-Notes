<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Intermediate-01 / Networking / HTTP Server

## Overview

This directory explores **GO-Intermediate-01 / Networking / HTTP Server** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Intermediate-01 / Networking / HTTP Server in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`handle_vs_handleFunc.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/04 - Networking/HTTP Server/handle_vs_handleFunc.go) | Go | what we want to do, when /api is called. Both does the same perpose |

## Concepts

### 1. Handle Vs Handlefunc (`handle_vs_handleFunc.go`)

File: [`handle_vs_handleFunc.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/04 - Networking/HTTP Server/handle_vs_handleFunc.go)  
what we want to do, when /api is called. Both does the same perpose.

Relevant code excerpt:

```go
package main

import "net/http"

// what we want to do, when /api is called.
func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\r\n"))
}

func main() {

	// Both does the same perpose
	// http.Handle registers object
	// http.HandleFunc registers function

	// handles /api path
	http.HandleFunc("/api", apiHandler)

	// handles /api/ path
	http.Handle("/api/", http.HandlerFunc(apiHandler))

	http.ListenAndServe(":8080", nil)
}
```

**Explanation**:
- Implements function(s): `apiHandler()`, `main()`.
- Defines C routine(s): `apiHandler()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. Entry into [`handle_vs_handleFunc.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/04 - Networking/HTTP Server/handle_vs_handleFunc.go) (`main` or top-level routine).
2. Initializes internal state / variables and executes core logic.
3. Processes inputs, performs transformations or system calls, and outputs results.

```text
main() [handle_vs_handleFunc.go]
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
go run handle_vs_handleFunc.go
```

## Important Details

- **Memory & Scope**: Notice variable allocation, pointer indirection, and lifetime across function boundaries.
- **Error & Return Handling**: Always verify return values and errors before proceeding to prevent panics or undefined behavior.
- **Resource Management**: Check that open files, network sockets, database handles, and heap allocations are properly closed or freed.

## Common Mistakes

- Assuming implicit synchronization or thread safety where none is provided.
- Ignoring error return values or missing zero-value edge cases.
- Misunderstanding pass-by-value versus pointer/reference semantics for complex types.

## Language Notes

- **Language Features**: Written using idiomatic Go paradigms.

## Related Concepts

- Data Structures & Algorithms in Go
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `handle_vs_handleFunc.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `handle_vs_handleFunc.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`handle_vs_handleFunc.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/04 - Networking/HTTP Server/handle_vs_handleFunc.go)
