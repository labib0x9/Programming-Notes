<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Intermediate-01 / Goroutine And Channel / day10 / syncIssue

## Overview

This directory explores **GO-Intermediate-01 / Goroutine And Channel / day10 / syncIssue** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Intermediate-01 / Goroutine And Channel / day10 / syncIssue in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`1.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/1.go) | Go | var mu sync.Mutex mu.Lock() |
| [`2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/2.go) | Go | Process rune by rune from input and count results |
| [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/main.go) | Go | "fmt" "sync" |
| [`raceCon.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/raceCon.go) | Go | "fmt" "sync" |

## Concepts

### 1. 1 (`1.go`)

File: [`1.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/1.go)  
var mu sync.Mutex mu.Lock().

Relevant code excerpt:

```go
func main() {

	var wg sync.WaitGroup
	// var mu sync.Mutex

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(I int) {
			defer wg.Done()
			// mu.Lock()
			_, min, sec := time.Now().Clock()
			fmt.Println(min, sec , " :: ", rand.Intn(1000))
			// mu.Unlock()
		}(i)
	}

	wg.Wait()
}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`, `func()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. 2 (`2.go`)

File: [`2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/2.go)  
Process rune by rune from input and count results.

Relevant code excerpt:

```go
type Result struct {
	filename                                   string
	byteCount, lineCount, wordCount, charCount int
	err                                        error
}

// Process rune by rune from input and count results
func inputProcess(r io.Reader) (result Result) {
	reader := bufio.NewReader(r)
	inWord := false

	for {
		in, size, err := reader.ReadRune() // Reads a single rune
		if err == io.EOF {                 // Check if for End Of File
			break
		}
		if err != nil {
			log.Println(err)
			result.err = err
			break
		}
```

**Explanation**:
- Implements function(s): `inputProcess()`, `PrintOutput()`, `main()`.
- Defines C routine(s): `PrintOutput()`, `main()`, `func()`, `func()`.
- Defines Go struct(s): `Result`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Main (`main.go`)

File: [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/main.go)  
"fmt" "sync".

Relevant code excerpt:

```go
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var counter = 0
// var mu sync.Mutex

// func increment(wg *sync.WaitGroup) {
// 	defer wg.Done()		// wg = wg - 1
// 	for i := 0; i < 100; i++ {
// 		mu.Lock()
// 		counter++
// 		mu.Unlock()
// 	}
// }

// func main() {

// 	var wg sync.WaitGroup	// wg = 0
```

**Explanation**:
- Implements function(s): `increment()`, `main()`.
- Defines C routine(s): `increment()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Racecon (`raceCon.go`)

File: [`raceCon.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/raceCon.go)  
"fmt" "sync".

Relevant code excerpt:

```go
// package main

// import (
// 	"fmt"
// 	"sync"
// )

/** What is the problem here ?
race condition
**/

// var counter = 0

// func increment(wg *sync.WaitGroup) {
// 	defer wg.Done()		// wg = wg - 1
// 	for i := 0; i < 100; i++ {
// 		counter++ // Race
// 	}
// }

// func main() {
```

**Explanation**:
- Implements function(s): `increment()`, `main()`.
- Defines C routine(s): `increment()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`1.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/1.go) provides or tests `1`.
2. [`2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/2.go) provides or tests `2`.
3. [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/main.go) provides or tests `main`.
4. [`raceCon.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/raceCon.go) provides or tests `raceCon`.

```text
Caller / Test Runner
  ├──> [1.go] (var mu sync.Mutex mu.Lock()...)
  ├──> [2.go] (Process rune by rune from input and...)
  ├──> [main.go] ("fmt" "sync"...)
  ├──> [raceCon.go] ("fmt" "sync"...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run 1.go
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

## Issues / Risks

> File [`raceCon.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/raceCon.go) explicitly demonstrates a race condition, synchronization flaw, or edge-case issue for educational analysis.

## Related Concepts

- Data Structures & Algorithms in Go
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `1.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `1.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`1.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/1.go)
- [`2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/2.go)
- [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/main.go)
- [`raceCon.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/syncIssue/raceCon.go)
