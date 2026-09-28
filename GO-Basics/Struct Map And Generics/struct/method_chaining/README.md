<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Basics / Struct Map And Generics / struct / method chaining

## Overview

This directory explores **GO-Basics / Struct Map And Generics / struct / method chaining** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Basics / Struct Map And Generics / struct / method chaining in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`method_chaining.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Struct Map And Generics/struct/method_chaining/method_chaining.go) | Go | Receiver type is pointer Return type is pointer |
| [`user_builder.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Struct Map And Generics/struct/method_chaining/user_builder.go) | Go | Implements `NewUserBuilder, SetUserName, SetEmail` logic for user builder |

## Concepts

### 1. Method Chaining (`method_chaining.go`)

File: [`method_chaining.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Struct Map And Generics/struct/method_chaining/method_chaining.go)  
Receiver type is pointer Return type is pointer.

Relevant code excerpt:

```go
type Teacher struct {
	Name string
	Age  int
	Id string
}

// Receiver type is pointer
// Return type is pointer
func (t *Teacher) setName(Name string) *Teacher {
	t.Name = Name
	return t
}

// Receiver type is value
// Return type is value
func (t Teacher) setAge(Age int) Teacher {
	t.Age = Age
	return t
}

// func (t *Teacher) setId(Id string) *Teacher {
// 	t.Id = Id
```

**Explanation**:
- Implements function(s): `setName()`, `setAge()`, `setId()`, `setId()`, `main()`.
- Defines C routine(s): `main()`.
- Defines Go struct(s): `Teacher`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. User Builder (`user_builder.go`)

File: [`user_builder.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Struct Map And Generics/struct/method_chaining/user_builder.go)  
Implements `NewUserBuilder, SetUserName, SetEmail` logic for user builder.

Relevant code excerpt:

```go
type UserBuilder struct {
	Username string
	Email    string
	Err      error
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (u *UserBuilder) SetUserName(username string) *UserBuilder {
	if username == "" {
		u.Err = errors.New("username is empty")
	}
	u.Username = username
	return u
}

func (u *UserBuilder) SetEmail(email string) *UserBuilder {
	if email == "" {
		u.Err = errors.New("email is empty")
	}
```

**Explanation**:
- Implements function(s): `NewUserBuilder()`, `SetUserName()`, `SetEmail()`, `Save()`, `Main()`.
- Defines C routine(s): `Main()`.
- Defines Go struct(s): `UserBuilder`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`method_chaining.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Struct Map And Generics/struct/method_chaining/method_chaining.go) provides or tests `method_chaining`.
2. [`user_builder.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Struct Map And Generics/struct/method_chaining/user_builder.go) provides or tests `user_builder`.

```text
Caller / Test Runner
  ├──> [method_chaining.go] (Receiver type is pointer Return typ...)
  ├──> [user_builder.go] (Implements `NewUserBuilder, SetUser...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run method_chaining.go
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

1. What is the primary role of the functions/routines demonstrated in `method_chaining.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `method_chaining.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`method_chaining.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Struct Map And Generics/struct/method_chaining/method_chaining.go)
- [`user_builder.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Basics/Struct Map And Generics/struct/method_chaining/user_builder.go)
