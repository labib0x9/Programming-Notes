<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Intermediate-01 / Goroutine And Channel / unbuffer channel / send receieve issue

## Overview

This directory explores **GO-Intermediate-01 / Goroutine And Channel / unbuffer channel / send receieve issue** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Intermediate-01 / Goroutine And Channel / unbuffer channel / send receieve issue in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`sample.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/unbuffer_channel/send_receieve_issue/sample.go) | Go | Implements `Run, Run2` logic for sample |

## Concepts

### 1. Sample (`sample.go`)

File: [`sample.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/unbuffer_channel/send_receieve_issue/sample.go)  
Implements `Run, Run2` logic for sample.

Relevant code excerpt:

```go
func Run(ch chan int) {
	fmt.Println("Running !!!")
	ch <- 10          // Blocks here, waiting to send
	fmt.Println(<-ch) // this line never executes
}

func Run2(ch chan int) {
	fmt.Println("Running 2 !!!")
	go func() {
		fmt.Println(<-ch) // receiver goroutine is created and waits here
	}()
	ch <- 10 // sender sends, now that a receiver is waiting in another goroutine
}

func main() {

	ch := make(chan int)

	go Run(ch)
	// go Run2(ch)

	time.Sleep(5 * time.Second)
```

**Explanation**:
- Implements function(s): `Run()`, `Run2()`, `main()`.
- Defines C routine(s): `Run()`, `Run2()`, `func()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. Entry into [`sample.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/unbuffer_channel/send_receieve_issue/sample.go) (`main` or top-level routine).
2. Initializes internal state / variables and executes core logic.
3. Processes inputs, performs transformations or system calls, and outputs results.

```text
main() [sample.go]
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
go run sample.go
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

1. What is the primary role of the functions/routines demonstrated in `sample.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `sample.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`sample.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/unbuffer_channel/send_receieve_issue/sample.go)
