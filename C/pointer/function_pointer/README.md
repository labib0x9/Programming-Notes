<!-- AUTO-GENERATED FROM SOURCE CODE -->
# C / pointer / function pointer

## Overview

This directory explores **C / pointer / function pointer** in **C**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of C / pointer / function pointer in C.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`function_pointer.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/function_pointer/function_pointer.c) | C | include<stdio.h> include<stdlib.h> |
| [`server_handler.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/function_pointer/server_handler.c) | C | include<stdio.h> include<stdlib.h> |
| [`task_function.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/function_pointer/task_function.c) | C | include<stdio.h> include<stdlib.h> |

## Concepts

### 1. Function Pointer (`function_pointer.c`)

File: [`function_pointer.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/function_pointer/function_pointer.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
int main() {

	int args[2] = {10, 20};
	func f = add;
	int* sum = f(args);

	printf("%d\n", *sum);
	free(sum);

	/**** ****/
	int (*exec)(int, int);
	exec = minus;
	printf("%d\n", exec(20, 10));	// minus

	exec = mult;
	printf("%d\n", exec(10, 5));	// mult

	return 0;
}
```

**Explanation**:
- Defines C routine(s): `add()`, `minus()`, `mult()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Server Handler (`server_handler.c`)

File: [`server_handler.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/function_pointer/server_handler.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
void register_route(char* method, char* route, handler_func func) {
	Route x = {method, route, func};
}

int handle_home(int client_id, char* body) {
	printf("Hello\n");
	return 0;
}

int handle_login(int client_id, char* body) {
	printf("LOGIN\n");
	return 0;
}

int main() {

	register_route("GET", "/", handle_home);
	register_route("GET", "/login", handle_login);

	return 0;
}
```

**Explanation**:
- Defines C routine(s): `register_route()`, `handle_home()`, `handle_login()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Task Function (`task_function.c`)

File: [`task_function.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/function_pointer/task_function.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
int main() {

	int* x = (int*) malloc(sizeof(int));
	*x = 10;

	task_node t = {print, x};
	t.func(t.args);

	t = (task_node){mult, x};
	int* s = t.func(t.args);
	printf("%d\n", *s);

	free(x);
	free(s);

	return 0;
}
```

**Explanation**:
- Defines C routine(s): `print()`, `mult()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`function_pointer.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/function_pointer/function_pointer.c) provides or tests `function_pointer`.
2. [`server_handler.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/function_pointer/server_handler.c) provides or tests `server_handler`.
3. [`task_function.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/function_pointer/task_function.c) provides or tests `task_function`.

```text
Caller / Test Runner
  ├──> [function_pointer.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [server_handler.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [task_function.c] (include<stdio.h> include<stdlib.h>...)
  └──> Execution Completion / Assertion
```

## Comparison: C Dynamic Slice vs Go Slice

### C Custom Slice Emulation
```c
typedef struct {
    int* data;
    size_t len;
    size_t cap;
} Slice;
```

### Go Slice Built-in
```go
s := make([]int, len, cap) // 24-byte runtime slice header: Data unsafe.Pointer, Len int, Cap int
```

## Compilation & Execution

```bash
# Compile with GCC
gcc -Wall -Wextra function_pointer.c server_handler.c task_function.c -o main
./main
```

## Important Details

- **Memory & Scope**: Notice variable allocation, pointer indirection, and lifetime across function boundaries.
- **Error & Return Handling**: Always verify return values and errors before proceeding to prevent panics or undefined behavior.
- **Resource Management**: Check that open files, network sockets, database handles, and heap allocations are properly closed or freed.
- **Struct Padding & Memory Alignment**: Compilers insert padding bytes between struct members to align them to natural CPU word boundaries unless `#pragma pack(1)` is specified.
- **Generic Swaps with `void*`**: Uses `void*` buffer pointers and `memcpy` with a dynamically sized temporary buffer to swap arbitrary data types polymorphically.
- **Dynamic Slice Emulation in C**: Implements Go-like slice semantics in C using a struct `{ void* ptr; size_t len; size_t cap; }`.

## Common Mistakes

- Assuming implicit synchronization or thread safety where none is provided.
- Ignoring error return values or missing zero-value edge cases.
- Misunderstanding pass-by-value versus pointer/reference semantics for complex types.
- Unsafe string manipulation with `strcpy`/`strcat` without buffer length limits, causing stack buffer overflow vulnerabilities.
- Pointer arithmetic mismatch: `(ptr + 1)` advances by `sizeof(*ptr)` bytes, not 1 byte, unless `ptr` is `char*` or `uint8_t*`.

## Language Notes

- **Language Features**: Written using idiomatic C paradigms.
- **C Pointers**: Untyped memory pointers (`void*`) allow low-level introspection but require explicit casting and size specifications.

## Related Concepts

- Data Structures & Algorithms in C
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `function_pointer.c`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `function_pointer.c` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`function_pointer.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/function_pointer/function_pointer.c)
- [`server_handler.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/function_pointer/server_handler.c)
- [`task_function.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/function_pointer/task_function.c)
