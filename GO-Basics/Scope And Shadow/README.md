<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Basics / Scope And Shadow

## Overview

This directory explores **GO-Basics / Scope And Shadow** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Basics / Scope And Shadow in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`shadow_1.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Scope And Shadow/shadow_1.go) | Go | Demonstrates main entry point and workflow for shadow 1 |
| [`shadow_2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Scope And Shadow/shadow_2.go) | Go | Implements `global` logic for shadow 2 |
| [`shadow_3.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Scope And Shadow/shadow_3.go) | Go | Implements `foo` logic for shadow 3 |

## Concepts

### 1. Shadow 1 (`shadow_1.go`)

File: [`shadow_1.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Scope And Shadow/shadow_1.go)  
Demonstrates main entry point and workflow for shadow 1.

Relevant code excerpt:

```go
package main

import "fmt"

func main() {

	/**
	for short declaration := ,
	inner scope declares a new variables,
	and this variables lives only inside that inner scope,
	inner scope does't affect the outer scope variable
	**/

	x := 20

	if true {
		x := 30
		fmt.Println(x) // 30
	}

	fmt.Println(x) // 20

}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Shadow 2 (`shadow_2.go`)

File: [`shadow_2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Scope And Shadow/shadow_2.go)  
Implements `global` logic for shadow 2.

Relevant code excerpt:

```go
package main

import "fmt"

var x = 100

func main() {
	x := 50
	{
		x := 10
		fmt.Println(x)
	}
	fmt.Println(x)
	fmt.Println(global())
}

func global() int {
	return x
}
```

**Explanation**:
- Implements function(s): `main()`, `global()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Shadow 3 (`shadow_3.go`)

File: [`shadow_3.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Scope And Shadow/shadow_3.go)  
Implements `foo` logic for shadow 3.

Relevant code excerpt:

```go
package main

import "fmt"

func foo(x int) {
	fmt.Println(x)
	{
		x := x + 10
		fmt.Println(x)
	}
	fmt.Println(x)
}

func main() {
	foo(5)
}
```

**Explanation**:
- Implements function(s): `foo()`, `main()`.
- Defines C routine(s): `foo()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`shadow_1.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Scope And Shadow/shadow_1.go) provides or tests `shadow_1`.
2. [`shadow_2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Scope And Shadow/shadow_2.go) provides or tests `shadow_2`.
3. [`shadow_3.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Scope And Shadow/shadow_3.go) provides or tests `shadow_3`.

```text
Caller / Test Runner
  ├──> [shadow_1.go] (Demonstrates main entry point and w...)
  ├──> [shadow_2.go] (Implements `global` logic for shado...)
  ├──> [shadow_3.go] (Implements `foo` logic for shadow 3...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run shadow_1.go
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

1. What is the primary role of the functions/routines demonstrated in `shadow_1.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `shadow_1.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`shadow_1.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Scope And Shadow/shadow_1.go)
- [`shadow_2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Scope And Shadow/shadow_2.go)
- [`shadow_3.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Scope And Shadow/shadow_3.go)
