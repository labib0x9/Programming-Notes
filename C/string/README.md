<!-- AUTO-GENERATED FROM SOURCE CODE -->
# C / string

## Overview

This directory explores **C / string** in **C**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of C / string in C.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`basic.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/string/basic.c) | C | include<stdio.h> include<string.h> |
| [`buffer_overflow.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/string/buffer_overflow.c) | C | include<stdio.h> include<stdlib.h> |
| [`string_builder.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/string/string_builder.c) | C | include<stdio.h> include<stdlib.h> |

## Concepts

### 1. Basic (`basic.c`)

File: [`basic.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/string/basic.c)  
include<stdio.h> include<string.h>.

Relevant code excerpt:

```c
int main() {

    // string array -> stores at stack -> modifiable
    char s[] = "Hello";
    s[0] = 'X';
    printf("%s\n", s);
    printf("%lu\n", sizeof(s)); // 6 = strlen + '\0'
    printf("%lu\n", strlen(s)); // 5

    // string literal -> stores at read-only section -> non modifiable
    char *ss = "vauvau";
    // ss[0] = 'V'; // bus error -> SIGBUS/SIGSEGV
    printf("%s\n", ss); // no change
    printf("%lu\n", sizeof(ss));  // 8 = sizeof(pointer)
    printf("%lu\n", sizeof(*ss)); // 1 = s[0] -> one char
    printf("%lu\n", strlen(ss));  // 5

    ss = "but i can !"; // can assign another string..
    printf("%s\n", ss);

    ss = s;
    printf("%s\n", ss);
```

**Explanation**:
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Buffer Overflow (`buffer_overflow.c`)

File: [`buffer_overflow.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/string/buffer_overflow.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
int main() {

	/* code - 01 */

	char buffer[3];
	gets(buffer);

	printf("%s\n", buffer);

	/* code - 02 */
	char buf[1024];
	fgets(&buf, sizeof(buf), stdin);
	buffer[strcspn(buf, "\n")] = '\0';

	char newStr[3];
	strcpy(newStr, buf);

	printf("%s\n", newStr);

	return 0;
}
```

**Explanation**:
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. String Builder (`string_builder.c`)

File: [`string_builder.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/string/string_builder.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
void freeString(string_t* str) {
    if (str->data) free(str->data);
    str->data = NULL;
    str->len = 0;
}

int main() {

	return 0;
}
```

**Explanation**:
- Defines C routine(s): `NewString()`, `AppendString()`, `SubString()`, `SplitString()`, `freeString()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`basic.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/string/basic.c) provides or tests `basic`.
2. [`buffer_overflow.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/string/buffer_overflow.c) provides or tests `buffer_overflow`.
3. [`string_builder.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/string/string_builder.c) provides or tests `string_builder`.

```text
Caller / Test Runner
  ├──> [basic.c] (include<stdio.h> include<string.h>...)
  ├──> [buffer_overflow.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [string_builder.c] (include<stdio.h> include<stdlib.h>...)
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
gcc -Wall -Wextra basic.c buffer_overflow.c string_builder.c -o main
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

## Issues / Risks

> File [`buffer_overflow.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/string/buffer_overflow.c) demonstrates unsafe buffer writing (C buffer overflow vulnerability). Do not use this pattern in production code.

## Related Concepts

- Data Structures & Algorithms in C
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `basic.c`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `basic.c` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`basic.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/string/basic.c)
- [`buffer_overflow.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/string/buffer_overflow.c)
- [`string_builder.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/string/string_builder.c)
