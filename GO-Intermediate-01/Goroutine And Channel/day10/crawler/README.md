<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Intermediate-01 / Goroutine And Channel / day10 / crawler

## Overview

This directory explores **GO-Intermediate-01 / Goroutine And Channel / day10 / crawler** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Intermediate-01 / Goroutine And Channel / day10 / crawler in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`contestEditorial.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/contestEditorial.go) | Go | Implements `extractTags, completeContest` logic for contesteditorial |
| [`hashFunction.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/hashFunction.go) | Go | Implements `removeNonContestLink, generateHash, getHash` logic for hashfunction |
| [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/main.go) | Go | https://books.toscrape.com/index.html |
| [`randomAgent.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/randomAgent.go) | Go | Implements `genFirefoxUA, genChromeUA, genEdgeUA` logic for randomagent |
| [`router.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/router.go) | Go | Implements `homePageHandler, abcContestHandler, arcContestHandler` logic for router |

## Concepts

### 1. Contesteditorial (`contestEditorial.go`)

File: [`contestEditorial.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/contestEditorial.go)  
Implements `extractTags, completeContest` logic for contesteditorial.

Relevant code excerpt:

```go
func extractTags(url string) (tag []string) {
	// (*http.Request, error)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// ERR = err
		log.Println("ERRERER : GET REQ")
		return
	}

	// Set User-Agent (Custom Header)
	// mu.Lock()
	req.Header.Set("User-Agent", userAgent())
	// mu.Unlock()

	// Make a client, Use pointer
	client := &http.Client{}

	// (*http.Response, error)
	resp, err := client.Do(req)
	if err != nil {
		// ERR = err
		log.Println("ERRERER : DO REQ ", url)
```

**Explanation**:
- Implements function(s): `extractTags()`, `completeContest()`.
- Defines C routine(s): `completeContest()`, `FindLinks()`, `FindLinks()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Hashfunction (`hashFunction.go`)

File: [`hashFunction.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/hashFunction.go)  
Implements `removeNonContestLink, generateHash, getHash` logic for hashfunction.

Relevant code excerpt:

```go
func removeNonContestLink(links []string) string {
	var contestIds string
	for _, link := range links {
		link, found := strings.CutPrefix(link, contestPage)
		if found {
			if len(contestIds) == 0 {
				contestIds = link
			} else {
				contestIds = contestIds + "+" + link
			}
		}
	}
	return contestIds
}

func generateHash(contestIds string) string {
	hashTemp := sha256.Sum256([]byte(contestIds))
	hash := hex.EncodeToString(hashTemp[:])
	return hash
}

func getHash() string {
```

**Explanation**:
- Implements function(s): `removeNonContestLink()`, `generateHash()`, `getHash()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Main (`main.go`)

File: [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/main.go)  
https://books.toscrape.com/index.html.

Relevant code excerpt:

```go
type ProblemId struct {
	problemUrl   string
	editorialUrl string
	tags         []string
}

type ContestId struct {
	contestType string
	contestUrl  string
	problemUrl  []ProblemId
}

var BaseUrl = "https://atcoder.jp"
var token = make(chan struct{}, 5)
var ERR error
var mu sync.Mutex
var contestPage = BaseUrl + "/contests/"
var contest []ContestId
var contestPageHash string
var contestFound map[string]struct{}

func FindLinks(url string) (links []string) {
```

**Explanation**:
- Implements function(s): `FindLinks()`, `Traverse()`, `crawl()`, `validateUrl()`, `bfs()`, `getNewContest()`, `main()`.
- Defines C routine(s): `Traverse()`, `crawl()`, `bfs()`, `validateUrl()`, `func()`, `getNewContest()`, `main()`.
- Defines Go struct(s): `ProblemId`, `ContestId`.
- Assembly procedure(s): `This, continious` with DOS interrupt interactions.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Randomagent (`randomAgent.go`)

File: [`randomAgent.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/randomAgent.go)  
Implements `genFirefoxUA, genChromeUA, genEdgeUA` logic for randomagent.

Relevant code excerpt:

```go
func genFirefoxUA() string {
	version := firefoxVersions[rand.Intn(len(firefoxVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s; rv:%.1f) Gecko/20100101 Firefox/%.1f", os, version, version)
}

func genChromeUA() string {
	version := chromeVersions[rand.Intn(len(chromeVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", os, version)
}

func genEdgeUA() string {
	version := edgeVersions[rand.Intn(len(edgeVersions))]
	chromeVersion := strings.Split(version, ",")[0]
	edgeVersion := strings.Split(version, ",")[1]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Edg/%s", os, chromeVersion, edgeVersion)
}

func genOperaUA() string {
	version := operaVersions[rand.Intn(len(operaVersions))]
```

**Explanation**:
- Implements function(s): `genFirefoxUA()`, `genChromeUA()`, `genEdgeUA()`, `genOperaUA()`, `userAgent()`.
- Demonstrates step-by-step logic and runtime behavior.

### 5. Router (`router.go`)

File: [`router.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/router.go)  
Implements `homePageHandler, abcContestHandler, arcContestHandler` logic for router.

Relevant code excerpt:

```go
package main

import "github.com/gin-gonic/gin"

func homePageHandler(ctx *gin.Context) {

}

func abcContestHandler(ctx *gin.Context) {

}

func arcContestHandler(ctx *gin.Context) {

}

func runWeb() {
	router := gin.Default()

	router.GET("/", homePageHandler)
	router.GET("/abc", abcContestHandler)
	router.GET("/arc", arcContestHandler)

	router.Run(":8080")
}
```

**Explanation**:
- Implements function(s): `homePageHandler()`, `abcContestHandler()`, `arcContestHandler()`, `runWeb()`.
- Defines C routine(s): `homePageHandler()`, `abcContestHandler()`, `arcContestHandler()`, `runWeb()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`contestEditorial.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/contestEditorial.go) provides or tests `contestEditorial`.
2. [`hashFunction.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/hashFunction.go) provides or tests `hashFunction`.
3. [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/main.go) provides or tests `main`.
4. [`randomAgent.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/randomAgent.go) provides or tests `randomAgent`.

```text
Caller / Test Runner
  ├──> [contestEditorial.go] (Implements `extractTags, completeCo...)
  ├──> [hashFunction.go] (Implements `removeNonContestLink, g...)
  ├──> [main.go] (https://books.toscrape.com/index.ht...)
  ├──> [randomAgent.go] (Implements `genFirefoxUA, genChrome...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run contestEditorial.go
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

1. What is the primary role of the functions/routines demonstrated in `contestEditorial.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `contestEditorial.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`contestEditorial.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/contestEditorial.go)
- [`hashFunction.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/hashFunction.go)
- [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/main.go)
- [`randomAgent.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/randomAgent.go)
- [`router.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Goroutine And Channel/day10/crawler/router.go)
