<!-- AUTO-GENERATED FROM SOURCE CODE -->
# C / Struct

## Overview

This directory explores **C / Struct** in **C**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of C / Struct in C.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`slice.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/slice.c) | C | include<stdio.h> include<stdlib.h> |
| [`struct.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/struct.c) | C | include<stdio.h> include<stdlib.h> |
| [`struct_embeded_size.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/struct_embeded_size.c) | C | include<stdio.h> include<stdlib.h> |
| [`struct_packing.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/struct_packing.c) | C | include<stdio.h> include<stdbool.h> |
| [`struct_size.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/struct_size.c) | C | include<stdio.h> include<stdlib.h> |

## Concepts

### 1. Slice (`slice.c`)

File: [`slice.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/slice.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
void FreeSlice(slice_t s) {
	if (s.arr) free(s.arr);
}

void PrintSlice(slice_t s) {
	for (int i = 0; i < len(s); i++) {
		printf("%d ", At(s, i));
	}
	printf("\n");
}

int main() {

	printf("slice = %zu\n", sizeof(slice_t));
	slice_t arr = NewSlice(0, 0);
	if (arr.arr == NULL) {
		printf("Malloc fails\n");
		return 0;
	}
	arr = Append(arr, 10);
	arr = Append(arr, 20);
	arr = Append(arr, 30);
```

**Explanation**:
- Defines C routine(s): `NewSlice()`, `resize()`, `Append()`, `SubSlice()`, `At()`, `len()`, `FreeSlice()`, `PrintSlice()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Struct (`struct.c`)

File: [`struct.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/struct.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
union Data{
    int len;
    char *buffer;
};


int32_t main() {

    // printf("%lu\n", sizeof root);
    // printf("%lu\n", sizeof *root);

    slice x;
    x.ptr = NULL;
    x.len = 1;
    x.cap = 2;

    slice *y = &x;

    if (y->len == y->cap) {
        y->cap *= 2;
    }
```

**Explanation**:
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Struct Embeded Size (`struct_embeded_size.c`)

File: [`struct_embeded_size.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/struct_embeded_size.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
int main() {

	// size of struct
	printf("pair_t = %zu\n", sizeof(pair_t));
	printf("pool_elem_t = %zu\n\n", sizeof(pool_elem_t));

	// address and size
	uintptr_t prev = (uintptr_t) &object_pool[0];
	for (int i = 0; i < POOL_SIZE; i++) {
		uintptr_t cur = (uintptr_t) &object_pool[i];
		uintptr_t obj_addr = (uintptr_t) &object_pool[i].obj;
		printf("%lu %lu   ::::    %lu %lu\n", cur, cur - prev, obj_addr, obj_addr - cur);
		prev = cur;
	}

	// find the starting offset of field obj
	printf("\n\n");
	printf("%zu\n", offsetof(pool_elem_t, obj));

	size_t offset = offsetof(pool_elem_t, obj);
	uintptr_t obj_addr = (uintptr_t) &object_pool[1].obj;
	uintptr_t field_addr = obj_addr - offset;
```

**Explanation**:
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Struct Packing (`struct_packing.c`)

File: [`struct_packing.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/struct_packing.c)  
include<stdio.h> include<stdbool.h>.

Relevant code excerpt:

```c
struct A {
    bool a;
    int b;
    char c;
    long long d;
    char e;
};


struct B {
    bool a;
    int b;
    char c;
    long long d;
    char e;
}__attribute__((packed));


int main() {

    printf("%zu\n", sizeof(struct A));
    printf("%zu\n", sizeof(struct B));
```

**Explanation**:
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 5. Struct Size (`struct_size.c`)

File: [`struct_size.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/struct_size.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#include<stdint.h>
#include<stdbool.h>

typedef struct {
    bool alloc;
    int a;
    int b;
    long long c;
} a_t;

typedef struct {
    long long c;
    int a;
    int b;
    bool alloc;
} b_t;

int32_t main() {
```

**Explanation**:
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`slice.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/slice.c) provides or tests `slice`.
2. [`struct.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/struct.c) provides or tests `struct`.
3. [`struct_embeded_size.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/struct_embeded_size.c) provides or tests `struct_embeded_size`.
4. [`struct_packing.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/struct_packing.c) provides or tests `struct_packing`.

```text
Caller / Test Runner
  ├──> [slice.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [struct.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [struct_embeded_size.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [struct_packing.c] (include<stdio.h> include<stdbool.h>...)
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
gcc -Wall -Wextra slice.c struct.c struct_embeded_size.c struct_packing.c struct_size.c -o main
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

1. What is the primary role of the functions/routines demonstrated in `slice.c`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `slice.c` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`slice.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/slice.c)
- [`struct.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/struct.c)
- [`struct_embeded_size.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/struct_embeded_size.c)
- [`struct_packing.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/struct_packing.c)
- [`struct_size.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/05 - Struct/struct_size.c)
