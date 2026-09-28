<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Intermediate-01 / Goroutine And Channel / buffer channel / worker pool

## Overview

This directory explores **GO-Intermediate-01 / Goroutine And Channel / buffer channel / worker pool** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Intermediate-01 / Goroutine And Channel / buffer channel / worker pool in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`worker.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/buffer_channel/worker_pool/worker.go) | Go | Shared queue -> fixed sized worker pool NOTE: Distribution isn't fair. |

## Concepts

### 1. Worker (`worker.go`)

File: [`worker.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/buffer_channel/worker_pool/worker.go)  
Shared queue -> fixed sized worker pool NOTE: Distribution isn't fair..

Relevant code excerpt:

```go
func pool(numWorker int, task []int, result chan string, wg *sync.WaitGroup) {
	job := make(chan int, len(task)) // Shared job queue (buffered to length of tasks)

	// Create numWorker goroutines, each acting as a worker.
	for i := 0; i < numWorker; i++ {
		wg.Add(1) // Track each worker's completion
		go worker(i, job, result, wg)
	}

	// Submit all tasks into the shared job queue
	for _, j := range task {
		job <- j
	}
	close(job) // Signal to workers: no more tasks
}

// worker continuously pulls tasks from the job queue and sends results to the result channel.
func worker(curWorker int, job chan int, result chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Mark worker as done once it exits

	for ch := range job {
		result <- fmt.Sprintf("%d -> %d", curWorker, ch) // Report task handling
```

**Explanation**:
- Implements function(s): `pool()`, `worker()`, `main()`.
- Defines C routine(s): `pool()`, `worker()`, `main()`, `func()`.
- Assembly procedure(s): `might` with DOS interrupt interactions.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. Entry into [`worker.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/buffer_channel/worker_pool/worker.go) (`main` or top-level routine).
2. Initializes internal state / variables and executes core logic.
3. Processes inputs, performs transformations or system calls, and outputs results.

```text
main() [worker.go]
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
go run worker.go
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

1. What is the primary role of the functions/routines demonstrated in `worker.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `worker.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`worker.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/buffer_channel/worker_pool/worker.go)
