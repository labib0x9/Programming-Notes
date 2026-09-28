<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Intermediate-01 / Interface-intr

## Overview

This directory explores **GO-Intermediate-01 / Interface-intr** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Intermediate-01 / Interface-intr in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`function_interception.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface-intr/function_interception.go) | Go | another type this struct implements the interface |
| [`implements_io_reader.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface-intr/implements_io_reader.go) | Go | Necessary method to implement io.Reader Extra method |
| [`io_reader_to_io_bytescanner.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface-intr/io_reader_to_io_bytescanner.go) | Go | implements io.ByteScanner |

## Concepts

### 1. Function Interception (`function_interception.go`)

File: [`function_interception.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface-intr/function_interception.go)  
another type this struct implements the interface.

Relevant code excerpt:

```go
type MyNewInterface interface {
	Task(name string) (out string, err error)
}

// another type
type FancyName string

func (f FancyName) Task(name string) (out string, err error) {
	return "Fancy Name: " + string(f) + " " + name, nil
}

// this struct implements the interface
type Struct1 struct {
	mni MyNewInterface
}

func (s Struct1) Task(name string) (out string, err error) {
	if len(name) == 0 {
		return "", errors.New("name cannot be empty")
	}
	if strings.HasSuffix(name, "Faisal") {
		fmt.Println("Err = Contains Faisal as last name")
```

**Explanation**:
- Implements function(s): `Task()`, `Task()`, `Process()`, `main()`.
- Defines C routine(s): `Process()`, `main()`.
- Defines Go struct(s): `Struct1`.
- Assembly procedure(s): `func` with DOS interrupt interactions.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Implements Io Reader (`implements_io_reader.go`)

File: [`implements_io_reader.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface-intr/implements_io_reader.go)  
Necessary method to implement io.Reader Extra method.

Relevant code excerpt:

```go
type syscallReader struct {
	fd int	// file descriptor
}

// Necessary method to implement io.Reader
func (s *syscallReader) Read(p []byte) (n int, err error) {
	n, err = syscall.Read(s.fd, p)
	return
}

// Extra method
func (s *syscallReader) Temp() {
	fmt.Println("TEMPORARY CALL")
}

func ReadTenBytes(r io.Reader) {
	b := make([]byte, 10)
	if _, err := r.Read(b); err != nil {
		return
	}
	fmt.Println(b)
}
```

**Explanation**:
- Implements function(s): `Read()`, `Temp()`, `ReadTenBytes()`, `main()`.
- Defines C routine(s): `ReadTenBytes()`, `main()`.
- Defines Go struct(s): `syscallReader`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Io Reader To Io Bytescanner (`io_reader_to_io_bytescanner.go`)

File: [`io_reader_to_io_bytescanner.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface-intr/io_reader_to_io_bytescanner.go)  
implements io.ByteScanner.

Relevant code excerpt:

```go
type byteReader struct {
	r         io.Reader
	buf       [1]byte
	unread    bool
	prevBuf   byte
	firstCall bool
}

func (b *byteReader) ReadByte() (byte, error) {
	var bb byte
	var err error
	b.firstCall = false

	if b.unread {
		bb = b.prevBuf
		b.unread = false
		return bb, err
	}

	n, err := io.ReadFull(b.r, b.buf[:]) // read 1 byte
	if n != 1 {
		return bb, errors.New("longer")
```

**Explanation**:
- Implements function(s): `ReadByte()`, `UnreadByte()`, `read()`, `main()`.
- Defines C routine(s): `main()`.
- Defines Go struct(s): `byteReader`, `reader`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`function_interception.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface-intr/function_interception.go) provides or tests `function_interception`.
2. [`implements_io_reader.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface-intr/implements_io_reader.go) provides or tests `implements_io_reader`.
3. [`io_reader_to_io_bytescanner.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface-intr/io_reader_to_io_bytescanner.go) provides or tests `io_reader_to_io_bytescanner`.

```text
Caller / Test Runner
  ├──> [function_interception.go] (another type this struct implements...)
  ├──> [implements_io_reader.go] (Necessary method to implement io.Re...)
  ├──> [io_reader_to_io_bytescanner.go] (implements io.ByteScanner...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run function_interception.go
```

## Important Details

- **Memory & Scope**: Notice variable allocation, pointer indirection, and lifetime across function boundaries.
- **Error & Return Handling**: Always verify return values and errors before proceeding to prevent panics or undefined behavior.
- **Resource Management**: Check that open files, network sockets, database handles, and heap allocations are properly closed or freed.
- **Interface Representation**: An interface value is represented internally as a two-word pair: `(itab / type_descriptor, data_pointer)`.
- **Laws of Reflection**: 1. Reflection goes from interface value to reflection object (`reflect.TypeOf`, `reflect.ValueOf`). 2. Reflection goes from reflection object to interface value (`v.Interface()`). 3. To modify a reflection object, the value must be settable (`v.CanSet()`).

## Common Mistakes

- Assuming implicit synchronization or thread safety where none is provided.
- Ignoring error return values or missing zero-value edge cases.
- Misunderstanding pass-by-value versus pointer/reference semantics for complex types.
- Calling `.Set()` on a `reflect.Value` created from a non-pointer or unexported struct field, causing a runtime panic.
- Comparing an interface holding a typed nil pointer (`(*MyStruct)(nil)`) with `nil`, which evaluates to `false` because the type component is non-nil.

## Language Notes

- **Language Features**: Written using idiomatic Go paradigms.
- **Go Interfaces**: Implicit structural typing (duck typing) verified at compile time, with dynamic dispatch at runtime.

## Related Concepts

- Data Structures & Algorithms in Go
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `function_interception.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `function_interception.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`function_interception.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface-intr/function_interception.go)
- [`implements_io_reader.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface-intr/implements_io_reader.go)
- [`io_reader_to_io_bytescanner.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface-intr/io_reader_to_io_bytescanner.go)
