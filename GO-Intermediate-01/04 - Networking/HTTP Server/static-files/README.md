<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Intermediate-01 / Networking / HTTP Server / static-files

## Overview

This directory explores **GO-Intermediate-01 / Networking / HTTP Server / static-files** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Intermediate-01 / Networking / HTTP Server / static-files in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`serve_static_files.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/04 - Networking/HTTP Server/static-files/serve_static_files.go) | Go | logging // // http://127.0.0.1:8080/static.go -> ./static/static.go |
| [`wrapped_fs.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/04 - Networking/HTTP Server/static-files/wrapped_fs.go) | Go | http.FileSystem is an interface, with only `Open(name string) (File, error)` method |

## Concepts

### 1. Serve Static Files (`serve_static_files.go`)

File: [`serve_static_files.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/04 - Networking/HTTP Server/static-files/serve_static_files.go)  
logging // // http://127.0.0.1:8080/static.go -> ./static/static.go.

Relevant code excerpt:

```go
func printRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // without this middleware doesn't work
		fmt.Println("After", r.Method, r.URL.Path)
	})
}

func main() {

	// // // http://127.0.0.1:8080/static.go -> ./static/static.go

	// // use /static as file loading
	root := http.Dir("./static")
	fs := http.FileServer(root) // root -> http.FileSystem

	// without -> http://host:port/ -> ./static
	// with ->    http://host:port/static/ -> ./static
	// Route: /static/* → ./static/*
	fsHandler := http.StripPrefix("/static/", fs)

	// http://127.0.0.1:8080/static/.env -> exposed
```

**Explanation**:
- Implements function(s): `printRequest()`, `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Wrapped Fs (`wrapped_fs.go`)

File: [`wrapped_fs.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/04 - Networking/HTTP Server/static-files/wrapped_fs.go)  
http.FileSystem is an interface, with only `Open(name string) (File, error)` method.

Relevant code excerpt:

```go
type wrapFS struct {
	fh http.FileSystem
}

func (w wrapFS) Open(name string) (http.File, error) {
	clean := path.Clean(name)
	fmt.Println("WrapFS=", clean)
	base := path.Base(clean)
	if strings.HasPrefix(base, ".") == true {
		return nil, errors.New("No no, you can't see.")
	}
	return w.fh.Open(clean)
}

// logging
func printRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Incoming=", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // without this middleware doesn't work
	})
}
```

**Explanation**:
- Implements function(s): `Open()`, `printRequest()`, `main()`.
- Defines C routine(s): `main()`.
- Defines Go struct(s): `wrapFS`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`serve_static_files.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/04 - Networking/HTTP Server/static-files/serve_static_files.go) provides or tests `serve_static_files`.
2. [`wrapped_fs.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/04 - Networking/HTTP Server/static-files/wrapped_fs.go) provides or tests `wrapped_fs`.

```text
Caller / Test Runner
  ├──> [serve_static_files.go] (logging // // http://127.0.0.1:8080...)
  ├──> [wrapped_fs.go] (http.FileSystem is an interface, wi...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run serve_static_files.go
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

1. What is the primary role of the functions/routines demonstrated in `serve_static_files.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `serve_static_files.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`serve_static_files.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/04 - Networking/HTTP Server/static-files/serve_static_files.go)
- [`wrapped_fs.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/04 - Networking/HTTP Server/static-files/wrapped_fs.go)
