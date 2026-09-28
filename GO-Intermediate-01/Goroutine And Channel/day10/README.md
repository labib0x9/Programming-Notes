<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Intermediate-01 / Goroutine And Channel / day10

## Overview

This directory explores **GO-Intermediate-01 / Goroutine And Channel / day10** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Intermediate-01 / Goroutine And Channel / day10 in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`crawl.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawl.go) | Go | Implements `FindLinks, extractUrl` logic for crawl |
| [`crawler.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler.go) | Go | What is the problem in this code ?? go func() { |
| [`crawler2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler2.go) | Go | Need to check, acquire token mechanism Enforce a limit of 20 concurrent request at once |
| [`pipelines.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/pipelines.go) | Go | Pipelines -> one goroutines output is another goroutines input // Counter |

## Concepts

### 1. Crawl (`crawl.go`)

File: [`crawl.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawl.go)  
Implements `FindLinks, extractUrl` logic for crawl.

Relevant code excerpt:

```go
func FindLinks(url string) (urls []string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("ERROR")
		return
	}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Fatal("Error parsing body: ", url)
		return
	}

	extractUrl(&urls, doc)
```

**Explanation**:
- Implements function(s): `FindLinks()`, `extractUrl()`.
- Defines C routine(s): `extractUrl()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Crawler (`crawler.go`)

File: [`crawler.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler.go)  
What is the problem in this code ?? go func() {.

Relevant code excerpt:

```go
func Crawler1() {

	wordlist := make(chan []string)

	// go func() {
	// 	wordlist <- []string{"http://gopl.io"}
	// }()
	go Push("http://gopl.io", wordlist)

	seen := make(map[string]bool)
	for list := range wordlist {
		for _, link := range list {
			if seen[link] {
				continue
			}
			seen[link] = true
			// go func(link string) {
			// 	wordlist <- crawl(link)
			// }(link)
			go Push(link, wordlist)
		}
	}
```

**Explanation**:
- Implements function(s): `Crawler1()`, `Push()`, `crawl()`, `crawl()`.
- Defines C routine(s): `Crawler1()`, `func()`, `func()`, `Push()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Crawler2 (`crawler2.go`)

File: [`crawler2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler2.go)  
Need to check, acquire token mechanism Enforce a limit of 20 concurrent request at once.

Relevant code excerpt:

```go
func crawler2() {
	wordlist := make(chan []string)
	var n int

	n++
	tokens <- struct{}{}
	go Push2("http://gopl.io", wordlist)

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		lists := <-wordlist
		for _, link := range lists {
			if seen[link] {
				continue
			}
			seen[link] = true
			n++
			tokens <- struct{}{} // Acquire a token
			go Push2(link, wordlist)
		}
	}
}
```

**Explanation**:
- Implements function(s): `crawler2()`, `Push2()`, `crawl2()`.
- Defines C routine(s): `crawler2()`, `Push2()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Pipelines (`pipelines.go`)

File: [`pipelines.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/pipelines.go)  
Pipelines -> one goroutines output is another goroutines input // Counter.

Relevant code excerpt:

```go
func Pipeline() {

	naturals := make(chan int)
	squares := make(chan int)

	// // Counter
	// go func() {
	// 	for i := 0; i < 20; i++ {
	// 		naturals <- i
	// 	}
	// 	// Close the channel
	// 	close(naturals)
	// }()

	// // Squares
	// go func ()  {
	// 	for x := range naturals {
	// 		squares <- x * x
	// 	}
	// 	close(squares)
	// }()
```

**Explanation**:
- Implements function(s): `Pipeline()`, `Counter()`, `Square()`, `PrintNumber()`.
- Defines C routine(s): `Pipeline()`, `func()`, `func()`, `Counter()`, `Square()`, `PrintNumber()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`crawl.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawl.go) provides or tests `crawl`.
2. [`crawler.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler.go) provides or tests `crawler`.
3. [`crawler2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler2.go) provides or tests `crawler2`.
4. [`pipelines.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/pipelines.go) provides or tests `pipelines`.

```text
Caller / Test Runner
  ├──> [crawl.go] (Implements `FindLinks, extractUrl` ...)
  ├──> [crawler.go] (What is the problem in this code ??...)
  ├──> [crawler2.go] (Need to check, acquire token mechan...)
  ├──> [pipelines.go] (Pipelines -> one goroutines output ...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run crawl.go
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

1. What is the primary role of the functions/routines demonstrated in `crawl.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `crawl.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`crawl.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawl.go)
- [`crawler.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler.go)
- [`crawler2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler2.go)
- [`pipelines.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/pipelines.go)
