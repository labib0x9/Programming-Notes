<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Intermediate-01 / Goroutine And Channel / goroutine

## Overview

This directory explores **GO-Intermediate-01 / Goroutine And Channel / goroutine** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Intermediate-01 / Goroutine And Channel / goroutine in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`ex-01.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/goroutine/ex-01.go) | Go | main goroutine = parent test() creates a goroutine from main goroutine |
| [`goroutine.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/goroutine/goroutine.go) | Go | This code will run Print() which is a goroutine because main() function is also a goroutine and main goroutine, others are child goroutine |

## Concepts

### 1. Ex-01 (`ex-01.go`)

File: [`ex-01.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/goroutine/ex-01.go)  
main goroutine = parent test() creates a goroutine from main goroutine.

Relevant code excerpt:

```go
package main

import (
	"fmt"
	"time"
)

// main goroutine = parent
// test() creates a goroutine from main goroutine
// parent dies -> spawn goroutine dies

func test() {
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("GOROUTINE")
	}()
}

func main() {
	test()
	fmt.Println("MAIN")
	time.Sleep(30 * time.Second)
}

// OUTPUT: MAIN, GOROUTINE
```

**Explanation**:
- Implements function(s): `test()`, `main()`.
- Defines C routine(s): `test()`, `func()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Goroutine (`goroutine.go`)

File: [`goroutine.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/goroutine/goroutine.go)  
This code will run Print() which is a goroutine because main() function is also a goroutine and main goroutine, others are child goroutine.

Relevant code excerpt:

```go
package main

import "fmt"

// This code will run Print() which is a goroutine
// because main() function is also a goroutine and main goroutine, others are child goroutine
// So execution of main() function will terminate other child goroutine
func Goroutine() {
	go Print("wrold")
}

func Print(msg string) {
	fmt.Println("Hello ", msg)
}

func main() {
	Goroutine()
}
```

**Explanation**:
- Implements function(s): `Goroutine()`, `Print()`, `main()`.
- Defines C routine(s): `Goroutine()`, `Print()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`ex-01.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/goroutine/ex-01.go) provides or tests `ex-01`.
2. [`goroutine.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/goroutine/goroutine.go) provides or tests `goroutine`.

```text
Caller / Test Runner
  ├──> [ex-01.go] (main goroutine = parent test() crea...)
  ├──> [goroutine.go] (This code will run Print() which is...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run ex-01.go
```

## Important Details

- **Memory & Scope**: Notice variable allocation, pointer indirection, and lifetime across function boundaries.
- **Error & Return Handling**: Always verify return values and errors before proceeding to prevent panics or undefined behavior.
- **Resource Management**: Check that open files, network sockets, database handles, and heap allocations are properly closed or freed.
- **Goroutine Scheduling (M:N Model)**: Go maps $M$ goroutines onto $N$ OS threads across $P$ processor contexts with cooperative and pre-emptive scheduling.
- **Channel Synchronization**: Unbuffered channels require both sender and receiver to be ready (rendezvous), whereas buffered channels block only when full.
- **`sync.Once` & Singleton Pattern**: Guarantees that a critical initialization function is executed strictly once across concurrent goroutines using atomic operations and mutex fallback.

## Common Mistakes

- Assuming implicit synchronization or thread safety where none is provided.
- Ignoring error return values or missing zero-value edge cases.
- Misunderstanding pass-by-value versus pointer/reference semantics for complex types.
- Goroutine leakage: spawning goroutines that block indefinitely on an unread channel or unreleased lock.
- Sending to or closing a closed channel, which triggers an immediate runtime panic.
- Data races on loop iteration variables captured inside goroutine closures.

## Language Notes

- **Language Features**: Written using idiomatic Go paradigms.
- **Go Channels**: Thread-safe FIFO queues managed by runtime `hchan` structures with wait queues (`sudog`).

## Related Concepts

- Data Structures & Algorithms in Go
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `ex-01.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `ex-01.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`ex-01.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/goroutine/ex-01.go)
- [`goroutine.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/goroutine/goroutine.go)
