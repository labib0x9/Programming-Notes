<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Basics / Array And Slice

## Overview

This directory explores **GO-Basics / Array And Slice** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Basics / Array And Slice in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`array.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/array.go) | Go | Demonstrates main entry point and workflow for array |
| [`capacity.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/capacity.go) | Go | len = 8, cap = 8 |
| [`copy.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/copy.go) | Go | sl2 is a copy of sl1, they are not backed by the same array copy function -> copy(dst, src) |
| [`slice_declare.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/slice_declare.go) | Go | through variables and short declaration make(type, length, capacity) |
| [`temp.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/temp.go) | Go | // // PrintSlice(s1) // // 	PrintSlice(s2) |

## Concepts

### 1. Array (`array.go`)

File: [`array.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/array.go)  
Demonstrates main entry point and workflow for array.

Relevant code excerpt:

```go
package main

import "fmt"

func main() {
	a1 := [...]int{} // This is Array of zero size

	var a2 [3]int

	days := [...]string{
		1 : "Saturday",
		2 : "Sunday",
		3 : "Friday",
	}
	fmt.Println(days[1])

	arr := [5]int{0, 1, 2, 3, 4}

	fmt.Println(a1, a2, arr)

	var x [][]int
	fmt.Println(x)
}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Capacity (`capacity.go`)

File: [`capacity.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/capacity.go)  
len = 8, cap = 8.

Relevant code excerpt:

```go
package main

import "fmt"

func main() {
	// len = 8, cap = 8
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8}

	a := a1[3:5]
	fmt.Println(len(a), cap(a)) // 2 5

	a = a1[3:5:8]
	fmt.Println(len(a), cap(a)) // 2 5

	a = a1[3:5:5]
	fmt.Println(len(a), cap(a)) // 2 2
}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Copy (`copy.go`)

File: [`copy.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/copy.go)  
sl2 is a copy of sl1, they are not backed by the same array copy function -> copy(dst, src).

Relevant code excerpt:

```go
func PrintSlice(s []int) {
	fmt.Println(s)
}

func main() {

	/** --- **/

	sl1 := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 3)

	// sl2 is a copy of sl1, they are not backed by the same array
	copy(sl2, sl1)

	sl1[0] = 1000
	PrintSlice(sl2)

	/** --- **/

	// copy function -> copy(dst, src)
	// Remove range [2, 3] zero based index
	s1 := []int{0, 1, 2, 3, 4, 5}
```

**Explanation**:
- Implements function(s): `PrintSlice()`, `main()`.
- Defines C routine(s): `PrintSlice()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Slice Declare (`slice_declare.go`)

File: [`slice_declare.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/slice_declare.go)  
through variables and short declaration make(type, length, capacity).

Relevant code excerpt:

```go
package main

import "fmt"

func main() {
	// through variables and short declaration
	var s1 []int  // nil slice
	s2 := []int{} // empty slice

	// make(type, length, capacity)
	s3 := make([]int, 0)
	s4 := make([]int, 0, 10)

	// array slicing -> [low:high:capacity]
	arr := []int{1, 2, 4, 5, 6, 7, 8, 9, 10}
	s5 := arr[2:4]
	s6 := arr[3:5:5]
}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 5. Temp (`temp.go`)

File: [`temp.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/temp.go)  
// // PrintSlice(s1) // // 	PrintSlice(s2).

Relevant code excerpt:

```go
// // // PrintSlice(s1)
// // // 	PrintSlice(s2)

// // // 	PrintSlice(s3)
// // // 	PrintSlice(s4)

// // // 	// s[i : j] -> i to j - 1
// // // 	// 0 <= i <= j <= cap(s)
// // // 	// slice beyond cap(s) causes panic
// // // 	// slice beyond len(s) extends the slice
// // // 	// length = j - i
// // // 	arr := []int{1, 2, 4, 5, 6, 7, 8, 9, 10}
// // // 	s5 := arr[2:4]
// // // 	PrintSlice(s5)

// // // 	// cap(s) = 8
// // // 	// len(s) = 5
// // // 	s := []int{0, 1, 2, 3}
// // // 	s = append(s, 4)

// // // 	// Within cap, beyond len
// // // 	s6 := s[2:6]
```

**Explanation**:
- Implements function(s): `PrintSlice()`, `main()`, `Append()`, `AppendPointer()`, `printPointerSlice()`, `main()`, `Slice()`, `Print()`, `Print1()`, `Print2()`.
- Defines C routine(s): `PrintSlice()`, `main()`, `Append()`, `AppendPointer()`, `printPointerSlice()`, `main()`, `Slice()`, `Print()`, `Print1()`, `Print2()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`array.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/array.go) provides or tests `array`.
2. [`capacity.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/capacity.go) provides or tests `capacity`.
3. [`copy.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/copy.go) provides or tests `copy`.
4. [`slice_declare.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/slice_declare.go) provides or tests `slice_declare`.

```text
Caller / Test Runner
  ├──> [array.go] (Demonstrates main entry point and w...)
  ├──> [capacity.go] (len = 8, cap = 8...)
  ├──> [copy.go] (sl2 is a copy of sl1, they are not ...)
  ├──> [slice_declare.go] (through variables and short declara...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run array.go
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

1. What is the primary role of the functions/routines demonstrated in `array.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `array.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`array.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/array.go)
- [`capacity.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/capacity.go)
- [`copy.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/copy.go)
- [`slice_declare.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/slice_declare.go)
- [`temp.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/05 - Array And Slice/temp.go)
