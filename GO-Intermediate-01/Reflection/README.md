<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Intermediate-01 / Reflection

## Overview

This directory explores **GO-Intermediate-01 / Reflection** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Intermediate-01 / Reflection in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`AddFunction.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/AddFunction.go) | Go | Implements `Add` logic for addfunction |
| [`code1.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/code1.go) | Go | Demonstrates main entry point and workflow for code1 |
| [`code2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/code2.go) | Go | Implements `Print` logic for code2 |
| [`format_specific_print.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/format_specific_print.go) | Go | printTypeAndValue prints the type and value of the given interface{}. v is a rune that determines the format of the output. |
| [`need_to_test_setting_null_value.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/need_to_test_setting_null_value.go) | Go | Implements `encode, decode` logic for need to test setting null value |
| [`reflect_value.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/reflect_value.go) | Go | // panic: interface conversion: interface {} is main.MyInt, not int |
| [`static_vs_dynamic_type.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/static_vs_dynamic_type.go) | Go | Demonstrates main entry point and workflow for static vs dynamic type |
| [`static_vs_dynamic_value.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/static_vs_dynamic_value.go) | Go | Demonstrates main entry point and workflow for static vs dynamic value |
| [`struct_.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/struct_.go) | Go | fmt.Println(ut.NumField()) // panic: reflect: NumField of non-struct type *main.User |
| [`struct_decoder.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/struct_decoder.go) | Go | Decode decodes a map into a struct. Only supports string to string and string to MyInt conversions for simplicity. |
| [`type_assertion.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/type_assertion.go) | Go | Both file and r point to the same *os.File value What i did, i just copied value from one interface to another interface |

## Concepts

### 1. Addfunction (`AddFunction.go`)

File: [`AddFunction.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/AddFunction.go)  
Implements `Add` logic for addfunction.

Relevant code excerpt:

```go
type MyInt int
type AnotherMyInt MyInt

func Add(a, b interface{}) interface{} {
	vA := reflect.ValueOf(a)
	vB := reflect.ValueOf(b)
	if vA.Kind() != vB.Kind() {
		return fmt.Errorf("type mismatch: a=%s, b=%s", vA.Kind(), vB.Kind())
	}

	switch vA.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return vA.Int() + vB.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return vA.Uint() + vB.Uint()
	case reflect.Float32, reflect.Float64:
		return vA.Float() + vB.Float()
	case reflect.String:
		return vA.String() + vB.String()
	default:
		return fmt.Errorf("unsupported type: %s", vA.Kind())
	}
```

**Explanation**:
- Implements function(s): `Add()`, `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Code1 (`code1.go`)

File: [`code1.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/code1.go)  
Demonstrates main entry point and workflow for code1.

Relevant code excerpt:

```go
func main() {

	var x interface{}
	fmt.Println("Type:", reflect.TypeOf(x), "Value:", x) // {nil, nil}

	x = 42
	fmt.Println("Type:", reflect.TypeOf(x), "Value:", x) // {int, 42}

	x = "Hello, World!"
	fmt.Println("Type:", reflect.TypeOf(x), "Value:", x) // {string, Hello, World!}

	fmt.Println("x =", x)
}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Code2 (`code2.go`)

File: [`code2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/code2.go)  
Implements `Print` logic for code2.

Relevant code excerpt:

```go
func Print(v interface{}) {
	fmt.Println(reflect.TypeOf(v))
	fmt.Println(reflect.TypeOf(v).Name())
	fmt.Println(reflect.TypeOf(v).Kind())
	fmt.Println(reflect.TypeOf(v).Elem().Kind())

	fmt.Println("-----")
}

func main() {

	sl := make([]int, 0, 10)
	Print(sl) // []int, "", slice, int

	al := [10]int{}
	Print(al) // [10]int, "", array, int

	mp := make(map[string]int)
	Print(mp) // map[string]int, "", map, int

	ch := make(chan int)
	Print(ch) // chan int, "", chan, int
```

**Explanation**:
- Implements function(s): `Print()`, `main()`.
- Defines C routine(s): `Print()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Format Specific Print (`format_specific_print.go`)

File: [`format_specific_print.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/format_specific_print.go)  
printTypeAndValue prints the type and value of the given interface{}. v is a rune that determines the format of the output..

Relevant code excerpt:

```go
func printTypeAndValue(x interface{}, verb rune) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	switch verb {
	case 'T':
		fmt.Println("Type:", t)
	case 'V':
		fmt.Println("Value:", v)
	case 'B':
		fmt.Println("Type:", t, ", Value:", v)
	default:
		fmt.Println("Invalid format specifier")
	}
}

func main() {

	printTypeAndValue(42, 'T')
	printTypeAndValue(42, 'V')
	printTypeAndValue(42, 'B')
```

**Explanation**:
- Implements function(s): `printTypeAndValue()`, `main()`.
- Defines C routine(s): `printTypeAndValue()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 5. Need To Test Setting Null Value (`need_to_test_setting_null_value.go`)

File: [`need_to_test_setting_null_value.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/need_to_test_setting_null_value.go)  
Implements `encode, decode` logic for need to test setting null value.

Relevant code excerpt:

```go
type Request struct {
	Name    *string `json:"name"`
	Status  *string `json:"status"`
	Persist *bool   `json:"persist"`
}

func encode(b []byte) io.Reader {
	return bytes.NewReader(b)
}

func decode(r io.Reader, req any) {
	var body map[string]json.RawMessage
	if err := json.NewDecoder(r).Decode(&body); err != nil {
		panic(err)
	}

	rv := reflect.ValueOf(req).Elem()
	fmt.Println("Field Number:", rv.NumField())

	for i := range rv.NumField() {
		f := rv.Type().Field(i).Tag.Get("json")
		fmt.Println("JSON FILED =", f)
```

**Explanation**:
- Implements function(s): `encode()`, `decode()`, `main()`.
- Defines C routine(s): `decode()`, `main()`.
- Defines Go struct(s): `Request`.
- Demonstrates step-by-step logic and runtime behavior.

### 6. Reflect Value (`reflect_value.go`)

File: [`reflect_value.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/reflect_value.go)  
// panic: interface conversion: interface {} is main.MyInt, not int.

Relevant code excerpt:

```go
type MyInt int

func main() {

	var x MyInt = 42
	v := reflect.ValueOf(x)

	fmt.Println(v.Kind() == reflect.Int)    // true
	fmt.Println(v.Type().Name() == "MyInt") // true

	fmt.Println(v.Int()) // 42

	fmt.Println("Can interface:", v.CanInterface()) // true

	i := v.Interface()          // i is of type interface{} and holds a MyInt
	fmt.Printf("%T %v\n", i, i) // main.MyInt 42

	// // panic: interface conversion: interface {} is main.MyInt, not int
	// f := i.(int)

	f := i.(MyInt) // Is i a MyInt?
	fmt.Println(f)
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`AddFunction.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/AddFunction.go) provides or tests `AddFunction`.
2. [`code1.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/code1.go) provides or tests `code1`.
3. [`code2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/code2.go) provides or tests `code2`.
4. [`format_specific_print.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/format_specific_print.go) provides or tests `format_specific_print`.

```text
Caller / Test Runner
  ├──> [AddFunction.go] (Implements `Add` logic for addfunct...)
  ├──> [code1.go] (Demonstrates main entry point and w...)
  ├──> [code2.go] (Implements `Print` logic for code2...)
  ├──> [format_specific_print.go] (printTypeAndValue prints the type a...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run AddFunction.go
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

1. What is the primary role of the functions/routines demonstrated in `AddFunction.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `AddFunction.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`AddFunction.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/AddFunction.go)
- [`code1.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/code1.go)
- [`code2.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/code2.go)
- [`format_specific_print.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/format_specific_print.go)
- [`need_to_test_setting_null_value.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/need_to_test_setting_null_value.go)
- [`reflect_value.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/reflect_value.go)
- [`static_vs_dynamic_type.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/static_vs_dynamic_type.go)
- [`static_vs_dynamic_value.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/static_vs_dynamic_value.go)
- [`struct_.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/struct_.go)
- [`struct_decoder.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/struct_decoder.go)
- [`type_assertion.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Reflection/type_assertion.go)
