<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Intermediate-01 / Error Handling / customError

## Overview

This directory explores **GO-Intermediate-01 / Error Handling / customError** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Intermediate-01 / Error Handling / customError in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`code_01.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Error Handling/customError/code_01.go) | Go | Custom Error Whay Error() ? |
| [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Error Handling/customError/main.go) | Go | implementing the error interface. error interface has Error() method Error() method's return type is string |

## Concepts

### 1. Code 01 (`code_01.go`)

File: [`code_01.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Error Handling/customError/code_01.go)  
Custom Error Whay Error() ?.

Relevant code excerpt:

```go
type NotFoundError struct {
	Resource string
}

// Whay Error() ?
// See error interface
func (n *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", n.Resource)
}

type ValidateError struct {
	Field string
	Message string
}

func (v *ValidateError) Error() string {
	return fmt.Sprintf("%s field %s", v.Field, v.Message)
}

func findUser(id int) error {

	if id > 10 {
```

**Explanation**:
- Implements function(s): `Error()`, `Error()`, `findUser()`, `Main()`.
- Defines C routine(s): `Main()`.
- Defines Go struct(s): `NotFoundError`, `ValidateError`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Main (`main.go`)

File: [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Error Handling/customError/main.go)  
implementing the error interface. error interface has Error() method Error() method's return type is string.

Relevant code excerpt:

```go
type A struct {
	msg string
}

func (a *A) Error() string {
	return fmt.Sprintf(a.msg)
}

type AA struct {
	msg string
}

func (a AA) Error() string {
	return fmt.Sprintf(a.msg)
}

func getError() error {
	errA := &A{msg: "Error Occurs from A"} // without & -> error ? why ? because A receives a pointer receiver
	errAA := AA{msg: "Error Occurs from AA"}
	return errors.Join(errA, errAA)
}
```

**Explanation**:
- Implements function(s): `Error()`, `Error()`, `getError()`, `Error()`, `checkBoundaryFor8ByteInt()`, `checkBoundaryFor16ByteInt()`, `main()`.
- Defines C routine(s): `main()`.
- Defines Go struct(s): `A`, `AA`, `B`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`code_01.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Error Handling/customError/code_01.go) provides or tests `code_01`.
2. [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Error Handling/customError/main.go) provides or tests `main`.

```text
Caller / Test Runner
  ├──> [code_01.go] (Custom Error Whay Error() ?...)
  ├──> [main.go] (implementing the error interface. e...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run code_01.go
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

1. What is the primary role of the functions/routines demonstrated in `code_01.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `code_01.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`code_01.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Error Handling/customError/code_01.go)
- [`main.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Error Handling/customError/main.go)
