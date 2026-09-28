<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Basics / Testing / table driven test

## Overview

This directory explores **GO-Basics / Testing / table driven test** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Basics / Testing / table driven test in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`table_driven.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Testing/table_driven_test/table_driven.go) | Go | Implements `Add, IsPalindrome` logic for table driven |
| [`table_driven_test.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Testing/table_driven_test/table_driven_test.go) | Go | One interesting thing ... By defination of int and int64 it would fail. but passed the test. Why ?? |

## Concepts

### 1. Table Driven (`table_driven.go`)

File: [`table_driven.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Testing/table_driven_test/table_driven.go)  
Implements `Add, IsPalindrome` logic for table driven.

Relevant code excerpt:

```go
package main

func Add(a, b int) int {
	return a + b
}

func IsPalindrome(msg string) bool {
	n := len(msg)
	for i := 0; i < n / 2; i++ {
		if msg[i] != msg[n - i - 1] {
			return false
		}
	}
	return true
}
```

**Explanation**:
- Implements function(s): `Add()`, `IsPalindrome()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Table Driven Test (`table_driven_test.go`)

File: [`table_driven_test.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Testing/table_driven_test/table_driven_test.go)  
One interesting thing ... By defination of int and int64 it would fail. but passed the test. Why ??.

Relevant code excerpt:

```go
func TestAdd(t *testing.T) {

	testCasesInt := []struct {
		a, b, sum int64
	}{
		{10, 20, 30},
		{40, 43, 83},

		// One interesting thing ...
		// By defination of int and int64 it would fail. but passed the test. Why ??
		// 32-bit architect machine : int works like int32
		// 64-bit architext machine : int works like int64
		{1000001212, 1212, 1000002424},
	}

	for i, test := range testCasesInt {
		t.Run(fmt.Sprintf("Int : Case-%d", i), func(t *testing.T) {
			expected := test.sum
			got := Add(int(test.a), int(test.b))
			if expected != int64(got) {
				t.Errorf("Add(%d, %d) = %d, want %d", test.a, test.b, got, expected)
			}
```

**Explanation**:
- Implements function(s): `TestAdd()`, `TestIsPalindrome()`.
- Defines C routine(s): `TestAdd()`, `TestIsPalindrome()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`table_driven.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Testing/table_driven_test/table_driven.go) provides or tests `table_driven`.
2. [`table_driven_test.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Testing/table_driven_test/table_driven_test.go) provides or tests `table_driven_test`.

```text
Caller / Test Runner
  ├──> [table_driven.go] (Implements `Add, IsPalindrome` logi...)
  ├──> [table_driven_test.go] (One interesting thing ... By defina...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run table_driven.go
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

1. What is the primary role of the functions/routines demonstrated in `table_driven.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `table_driven.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`table_driven.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Testing/table_driven_test/table_driven.go)
- [`table_driven_test.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Testing/table_driven_test/table_driven_test.go)
