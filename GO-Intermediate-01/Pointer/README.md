<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Intermediate-01 / Pointer

## Overview

This directory explores **GO-Intermediate-01 / Pointer** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Intermediate-01 / Pointer in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`basic.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/basic.go) | Go | Demonstrates main entry point and workflow for basic |
| [`header_parsing.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/header_parsing.go) | Go | Demonstrates main entry point and workflow for header parsing |
| [`int_to_slice.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/int_to_slice.go) | Go | Demonstrates main entry point and workflow for int to slice |
| [`make_vs_new.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/make_vs_new.go) | Go | So, both x, y is *int |
| [`pass_by_value_vs_reference.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/pass_by_value_vs_reference.go) | Go | Implements `PassByValue, PassByReference` logic for pass by value vs reference |
| [`pointer_arithmetic.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/pointer_arithmetic.go) | Go | pointer arithmetic now := uintptr(unsafe.Pointer(ptr)) // gc doesn't track, can be deallocate |
| [`reference_vs_pointer.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/reference_vs_pointer.go) | Go | Incomplete |
| [`reflect_value.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/reflect_value.go) | Go | elem -> int |
| [`slice_to_slice.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/slice_to_slice.go) | Go | Demonstrates main entry point and workflow for slice to slice |
| [`string_to_slice.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/string_to_slice.go) | Go | string -> slice |
| [`struct_fields.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/struct_fields.go) | Go | Demonstrates main entry point and workflow for struct fields |
| [`struct_to_slice.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/struct_to_slice.go) | Go | Implements `GetUser` logic for struct to slice |
| [`type_punning.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/type_punning.go) | Go | what happens here ?? type punning.. |

## Concepts

### 1. Basic (`basic.go`)

File: [`basic.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/basic.go)  
Demonstrates main entry point and workflow for basic.

Relevant code excerpt:

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x int
	p := &x
	t := reflect.TypeOf(p)

	fmt.Println(t)               // *int
	fmt.Println(t.Kind())        // ptr
	fmt.Println(t.Elem().Kind()) // int

}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Header Parsing (`header_parsing.go`)

File: [`header_parsing.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/header_parsing.go)  
Demonstrates main entry point and workflow for header parsing.

Relevant code excerpt:

```go
type Header struct {
	src uint16 // 2bytes
	dst uint16 // 2bytes
	tmp uint16 // 2bytes
}

func main() {

	data := []byte{
		0x12, 0x34, // 13330
		0x34, 0x12, // 4660
		0x52, 0x18, // -> why xByte to uint16 is 6226 ??? because of endianess
	}

	fmt.Println(len(data), unsafe.Sizeof(Header{})) // 6 6

	hdr := (*Header)(unsafe.Pointer(&data[0]))
	fmt.Println(hdr)

	/***/
	x := uint16(4660)
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Defines Go struct(s): `Header`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Int To Slice (`int_to_slice.go`)

File: [`int_to_slice.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/int_to_slice.go)  
Demonstrates main entry point and workflow for int to slice.

Relevant code excerpt:

```go
func main() {

	var x int = 256
	intSize := unsafe.Sizeof(x)

	ptr := (*byte)(unsafe.Pointer(&x)) // why type cast to *byte ??

	fmt.Println(ptr, intSize)
	xByte := unsafe.Slice(ptr, intSize) // int -> []byte

	fmt.Println(xByte)

	x = 123456 // changes xByte..

	fmt.Println(xByte)

	y := *(*int)(unsafe.Pointer(&xByte[0])) // []byte -> int

	fmt.Println(y)
}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Make Vs New (`make_vs_new.go`)

File: [`make_vs_new.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/make_vs_new.go)  
So, both x, y is *int.

Relevant code excerpt:

```go
func whatType(t reflect.Type) {
	fmt.Println(t)
	fmt.Println(t.Kind())
}

func main() {

	var x *int
	var y = new(int)

	t := reflect.TypeOf(x)
	whatType(t)

	t = reflect.TypeOf(y)
	whatType(t)

	// So, both x, y is *int

	fmt.Println(x) // nil
	fmt.Println(y) // has value

	// fmt.Println(*x) // nil -> panic in dereferencing
```

**Explanation**:
- Implements function(s): `whatType()`, `main()`.
- Defines C routine(s): `whatType()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 5. Pass By Value Vs Reference (`pass_by_value_vs_reference.go`)

File: [`pass_by_value_vs_reference.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/pass_by_value_vs_reference.go)  
Implements `PassByValue, PassByReference` logic for pass by value vs reference.

Relevant code excerpt:

```go
type A struct {
	v1 int
	v2 string
}

func PassByValue(v A) {
	ptr := unsafe.Pointer(&v)
	fmt.Println("PassByValue=", ptr)
}

func PassByReference(v *A) {
	ptr := unsafe.Pointer(v)
	fmt.Println("PassByReference=", ptr)
}

func main() {

	v := A{v1: 10, v2: "key"}

	ptr := unsafe.Pointer(&v)
	fmt.Println("main=", ptr)
```

**Explanation**:
- Implements function(s): `PassByValue()`, `PassByReference()`, `main()`.
- Defines C routine(s): `PassByValue()`, `PassByReference()`, `main()`.
- Defines Go struct(s): `A`.
- Demonstrates step-by-step logic and runtime behavior.

### 6. Pointer Arithmetic (`pointer_arithmetic.go`)

File: [`pointer_arithmetic.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/pointer_arithmetic.go)  
pointer arithmetic now := uintptr(unsafe.Pointer(ptr)) // gc doesn't track, can be deallocate.

Relevant code excerpt:

```go
func main() {

	// pointer arithmetic
	arr := [5]int{1, 2, 3, 4, 5}

	ptr := &arr[0]
	elemSize := unsafe.Sizeof(arr[0])

	fmt.Println(ptr, elemSize)

	// now := uintptr(unsafe.Pointer(ptr)) // gc doesn't track, can be deallocate

	index := 0
	arr_0 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(index)*elemSize))

	fmt.Println(*arr_0)

	index = 3
	arr_3 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(index)*elemSize))

	fmt.Println(*arr_3)
}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`basic.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/basic.go) provides or tests `basic`.
2. [`header_parsing.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/header_parsing.go) provides or tests `header_parsing`.
3. [`int_to_slice.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/int_to_slice.go) provides or tests `int_to_slice`.
4. [`make_vs_new.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/make_vs_new.go) provides or tests `make_vs_new`.

```text
Caller / Test Runner
  ├──> [basic.go] (Demonstrates main entry point and w...)
  ├──> [header_parsing.go] (Demonstrates main entry point and w...)
  ├──> [int_to_slice.go] (Demonstrates main entry point and w...)
  ├──> [make_vs_new.go] (So, both x, y is *int...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run basic.go
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

1. What is the primary role of the functions/routines demonstrated in `basic.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `basic.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`basic.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/basic.go)
- [`header_parsing.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/header_parsing.go)
- [`int_to_slice.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/int_to_slice.go)
- [`make_vs_new.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/make_vs_new.go)
- [`pass_by_value_vs_reference.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/pass_by_value_vs_reference.go)
- [`pointer_arithmetic.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/pointer_arithmetic.go)
- [`reference_vs_pointer.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/reference_vs_pointer.go)
- [`reflect_value.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/reflect_value.go)
- [`slice_to_slice.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/slice_to_slice.go)
- [`string_to_slice.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/string_to_slice.go)
- [`struct_fields.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/struct_fields.go)
- [`struct_to_slice.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/struct_to_slice.go)
- [`type_punning.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Pointer/type_punning.go)
