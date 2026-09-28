<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Intermediate-01 / Goroutine And Channel / day09 / bufferCh

## Overview

This directory explores **GO-Intermediate-01 / Goroutine And Channel / day09 / bufferCh** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Intermediate-01 / Goroutine And Channel / day09 / bufferCh in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`1.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day09/bufferCh/1.go) | Go | "fmt" "sync" |
| [`2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day09/bufferCh/2.go) | Go | "fmt" "time" |
| [`3.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day09/bufferCh/3.go) | Go | "fmt" "time" |
| [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day09/bufferCh/main.go) | Go | "fmt" "net/http" |

## Concepts

### 1. 1 (`1.go`)

File: [`1.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day09/bufferCh/1.go)  
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

// func increment(ch chan bool) {
// 	for i := 0; i < 100; i++ {
// 		mu.Lock()
// 		counter++
// 		mu.Unlock()
// 	}

// 	ch <-true
// }

// func main() {
```

**Explanation**:
- Implements function(s): `increment()`, `main()`.
- Defines C routine(s): `increment()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. 2 (`2.go`)

File: [`2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day09/bufferCh/2.go)  
"fmt" "time".

Relevant code excerpt:

```go
// package main

// import (
// 	"fmt"
// 	"time"
// )

// // func(ch chan<- int) // Sender
// // func(ch <-chan int) // Receiver
// // func(ch chan int) // Both

// func producer(ch chan<- int) {
// 	for i := 0; i <= 5; i++ {
// 		ch <- i
// 		fmt.Println("Produed: ", i)
// 		// time.Sleep(time.Millisecond * 500)
// 	}
// 	close(ch)
// }

// func consumer(ch <-chan int) {
// 	for data := range ch {
```

**Explanation**:
- Implements function(s): `producer()`, `consumer()`, `main()`.
- Defines C routine(s): `producer()`, `consumer()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. 3 (`3.go`)

File: [`3.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day09/bufferCh/3.go)  
"fmt" "time".

Relevant code excerpt:

```go
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func fakeAPI(response chan string) {
// 	// time.Sleep(3 * time.Second)
// 	time.Sleep(1 * time.Second)
// 	response <- "Received"
// }

// func fakeAPI2(response2 chan string) {
// 	time.Sleep(3 * time.Second)
// 	response2 <- "Received2"
// }

// func main() {

// 	response := make(chan string)
// 	response2 := make(chan string)
```

**Explanation**:
- Implements function(s): `fakeAPI()`, `fakeAPI2()`, `main()`.
- Defines C routine(s): `fakeAPI()`, `fakeAPI2()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Main (`main.go`)

File: [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day09/bufferCh/main.go)  
"fmt" "net/http".

Relevant code excerpt:

```go
func producer(ch chan<- int) {

	for i := 1; i <= 5; i++ {
		fmt.Println("Producing : ", i)
		ch <- i
		fmt.Println("Produced")
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Println("Producing done")

	// Run without close() and with close()
	close(ch)
}

// only receive
func consumer(ch <-chan int) {
	time.Sleep(time.Second * 5)
	for val := range ch {
		fmt.Println("Consumed: ", val)
		time.Sleep(time.Second * 5)
	}
```

**Explanation**:
- Implements function(s): `checkLink()`, `main()`, `producer()`, `consumer()`, `main()`.
- Defines C routine(s): `checkLink()`, `main()`, `producer()`, `consumer()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`1.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day09/bufferCh/1.go) provides or tests `1`.
2. [`2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day09/bufferCh/2.go) provides or tests `2`.
3. [`3.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day09/bufferCh/3.go) provides or tests `3`.
4. [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day09/bufferCh/main.go) provides or tests `main`.

```text
Caller / Test Runner
  ├──> [1.go] ("fmt" "sync"...)
  ├──> [2.go] ("fmt" "time"...)
  ├──> [3.go] ("fmt" "time"...)
  ├──> [main.go] ("fmt" "net/http"...)
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

- [`1.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day09/bufferCh/1.go)
- [`2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day09/bufferCh/2.go)
- [`3.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day09/bufferCh/3.go)
- [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day09/bufferCh/main.go)
