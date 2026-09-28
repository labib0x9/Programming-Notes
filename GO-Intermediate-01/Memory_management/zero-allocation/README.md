<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Intermediate-01 / Memory management / zero-allocation

## Overview

This directory explores **GO-Intermediate-01 / Memory management / zero-allocation** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Intermediate-01 / Memory management / zero-allocation in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`allocation.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/allocation.go) | Go | Demonstrates main entry point and workflow for allocation |
| [`allocation_count.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/allocation_count.go) | Go | Demonstrates main entry point and workflow for allocation count |
| [`slice_reuse.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/slice_reuse.go) | Go | Implements `check` logic for slice reuse |
| [`string_allocation_at_heap.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/string_allocation_at_heap.go) | Go | heap allocation for w |
| [`sync_pool.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/sync_pool.go) | Go | b = append(*b, s...) b = (*b)[0:] |
| [`zero_allocation.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/zero_allocation.go) | Go | Slice reuse sync.Pool |

## Concepts

### 1. Allocation (`allocation.go`)

File: [`allocation.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/allocation.go)  
Demonstrates main entry point and workflow for allocation.

Relevant code excerpt:

```go
func main() {

	b := []byte("Hello World")

	ptr := unsafe.Pointer(&b)
	fmt.Println(ptr)

	s := string(b)
	ptr = unsafe.Pointer(&s)
	fmt.Println(ptr)

	/** --- **/

	sl := []int{1, 2, 3}

	ptr = unsafe.Pointer(&sl)
	fmt.Println(ptr)

	sll := sl[:0]
	ptr = unsafe.Pointer(&sll)
	fmt.Println(ptr)
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Allocation Count (`allocation_count.go`)

File: [`allocation_count.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/allocation_count.go)  
Demonstrates main entry point and workflow for allocation count.

Relevant code excerpt:

```go
func main() {

	words := []string{"Hello", "world", "\n", "I", "am", "learning", "goooooooooooooooooooooooo"}

	allocCount := testing.AllocsPerRun(1000, func() {
		out := ""
		for _, w := range words {
			out += w
		}
	})
	fmt.Println("Allocations using += :", allocCount)

	allocCount = testing.AllocsPerRun(1000, func() {
		var sb strings.Builder
		sb.Grow(100) // preallocate enough capacity
		for _, w := range words {
			sb.WriteString(w)
		}
		out := sb.String()
		_ = out
	})
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Slice Reuse (`slice_reuse.go`)

File: [`slice_reuse.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/slice_reuse.go)  
Implements `check` logic for slice reuse.

Relevant code excerpt:

```go
func check(x []int) {
	fmt.Println("len:", len(x), "cap:", cap(x))
}

func main() {

	sl := make([]int, 0)

	ptr := unsafe.Pointer(&sl)
	fmt.Println("Allocates at:", ptr)

	check(sl) // {0, 0}

	for i := 0; i < 20; i++ {
		sl = append(sl, i)
	}

	check(sl) // {20, 32}

	/**
		length decreases, but the capacity remains the same..
		No allocation
```

**Explanation**:
- Implements function(s): `check()`, `main()`.
- Defines C routine(s): `check()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. String Allocation At Heap (`string_allocation_at_heap.go`)

File: [`string_allocation_at_heap.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/string_allocation_at_heap.go)  
heap allocation for w.

Relevant code excerpt:

```go
func main() {

	words := []string{"Hello", "world", "\n", "I", "am", "learning", "goooooooooooooooooooooooo"}

	out := ""

	ptr := unsafe.Pointer(&out)
	fmt.Println("Allocation at:", ptr)

	// heap allocation for w
	for _, w := range words {
		out += w
		ptr = unsafe.Pointer(&w)
		fmt.Println("Allocation at:", ptr)
	}

	out = ""
	for i := range words {
		out += words[i]
		fmt.Println(cap([]byte(out)), len(out)) // cap == len
	}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 5. Sync Pool (`sync_pool.go`)

File: [`sync_pool.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/sync_pool.go)  
b = append(*b, s...) b = (*b)[0:].

Relevant code excerpt:

```go
type Buffer []byte

func (b *Buffer) Append(s string) {
	*b = append(*b, s...)
}

func (b *Buffer) Truncate() {
	*b = (*b)[0:]
}

var bufPool = sync.Pool{
	New: func() any {
		fmt.Println("Allocating ...")
		return new(Buffer)
	},
}

func main() {

	// buf := bufPool.Get() -> Append() method doesn't work, without type assertion buf's type is any, we need to specify the type is Buffer.

	buf := bufPool.Get().(*Buffer) // pool is empty, calls New
```

**Explanation**:
- Implements function(s): `Append()`, `Truncate()`, `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 6. Zero Allocation (`zero_allocation.go`)

File: [`zero_allocation.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/zero_allocation.go)  
Slice reuse sync.Pool.

Relevant code excerpt:

```go
package main

import (
	"strings"
)

func main() {

	// Slice reuse
	s := make([]int, 10, 10000)
	s = s[:0]

	// sync.Pool

	// strings.Builder
	var x strings.Builder
	x.Grow(10000)
	x.WriteString("Hello")
}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`allocation.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/allocation.go) provides or tests `allocation`.
2. [`allocation_count.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/allocation_count.go) provides or tests `allocation_count`.
3. [`slice_reuse.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/slice_reuse.go) provides or tests `slice_reuse`.
4. [`string_allocation_at_heap.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/string_allocation_at_heap.go) provides or tests `string_allocation_at_heap`.

```text
Caller / Test Runner
  ├──> [allocation.go] (Demonstrates main entry point and w...)
  ├──> [allocation_count.go] (Demonstrates main entry point and w...)
  ├──> [slice_reuse.go] (Implements `check` logic for slice ...)
  ├──> [string_allocation_at_heap.go] (heap allocation for w...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run allocation.go
```

## Important Details

- **Memory & Scope**: Notice variable allocation, pointer indirection, and lifetime across function boundaries.
- **Error & Return Handling**: Always verify return values and errors before proceeding to prevent panics or undefined behavior.
- **Resource Management**: Check that open files, network sockets, database handles, and heap allocations are properly closed or freed.
- **Escape Analysis (`-gcflags='-m'`)**: The Go compiler determines whether a variable can be safely stack-allocated or must escape to the heap.
- **Zero-Allocation Optimization**: Reusing slice capacity `s = s[:0]`, utilizing `sync.Pool` to recycle temporary objects, and avoiding byte-to-string copy allocations.
- **`testing.AllocsPerRun`**: Measures exact heap allocations per iteration to benchmark and verify zero-allocation code paths.

## Common Mistakes

- Assuming implicit synchronization or thread safety where none is provided.
- Ignoring error return values or missing zero-value edge cases.
- Misunderstanding pass-by-value versus pointer/reference semantics for complex types.
- Passing concrete values into `fmt.Println` or interface parameters, causing them to escape to the heap due to boxing.
- Storing pointers to short-lived objects in `sync.Pool` and assuming they will survive indefinitely across GC cycles.

## Language Notes

- **Language Features**: Written using idiomatic Go paradigms.
- **Go GC**: Concurrent tri-color mark-and-sweep collector with low-latency pause targets.

## Related Concepts

- Data Structures & Algorithms in Go
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `allocation.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `allocation.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`allocation.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/allocation.go)
- [`allocation_count.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/allocation_count.go)
- [`slice_reuse.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/slice_reuse.go)
- [`string_allocation_at_heap.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/string_allocation_at_heap.go)
- [`sync_pool.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/sync_pool.go)
- [`zero_allocation.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Memory_management/zero-allocation/zero_allocation.go)
