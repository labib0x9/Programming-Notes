<!-- AUTO-GENERATED FROM SOURCE CODE -->
# C / pointer

## Overview

This directory explores **C / pointer** in **C**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of C / pointer in C.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`matrix.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/matrix.c) | C | include<stdio.h> include<stdlib.h> |
| [`swap_generic.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/swap_generic.c) | C | include<stdio.h> include<stdlib.h> |
| [`swap_int.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/swap_int.c) | C | include<stdio.h> include<stdlib.h> |
| [`swap_string.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/swap_string.c) | C | include<stdio.h> include<stdlib.h> |
| [`void_pointer.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/void_pointer.c) | C | include<stdio.h> include<stdlib.h> |
| [`zeroo_fill.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/zeroo_fill.c) | C | include<stdio.h> include<stdlib.h> |

## Concepts

### 1. Matrix (`matrix.c`)

File: [`matrix.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/matrix.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
for (int i = 0; i < n; i++) {
        int* elm = (int*) malloc(sizeof(int) * m);
        if (elm == NULL) {
            for (int j = 0; j < i; j++) {
                if (new.mat[j]) {
                    free(new.mat[j]);
                }
            }
            free(new.mat);
            new.mat = NULL;
            return new;
        }
        new.mat[i] = elm;
    }
    return new;
}

matrix_t MatrixAdd(matrix_t* a, matrix_t* b) {
    matrix_t tmp = NewMatrix(a->n, a->m);
    if (tmp.mat == NULL) {
        return;
    }
```

**Explanation**:
- Defines C routine(s): `NewMatrix()`, `MatrixAdd()`, `freeMatrix()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Swap Generic (`swap_generic.c`)

File: [`swap_generic.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/swap_generic.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
void swap(void* a, void* b, int size) {
	if (a == NULL || b == NULL || size == 0) return;
	void* tmp = (void*) malloc(size);
	if (tmp == NULL) return;

	// tmp = a;
	memmove(tmp, a, size);
	// a = b;
	memmove(a, b, size);
	// b = tmp;
	memmove(b, tmp, size);

	free(tmp);
}

int main() {

	/* */
	int x = 10, y = 20;
	swap(&x, &y, sizeof(int));

	printf("x = %d, y = %d\n", x, y);
```

**Explanation**:
- Defines C routine(s): `swap()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Swap Int (`swap_int.c`)

File: [`swap_int.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/swap_int.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#include<stdbool.h>

void swap(int* a, int* b) {
	
}

int main() {

	return 0;
}
```

**Explanation**:
- Defines C routine(s): `swap()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Swap String (`swap_string.c`)

File: [`swap_string.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/swap_string.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#include<stdbool.h>

void swap(char** a, char** b) {

}

int main() {

	return 0;
}
```

**Explanation**:
- Defines C routine(s): `swap()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 5. Void Pointer (`void_pointer.c`)

File: [`void_pointer.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/void_pointer.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
void test(void* a) {
	int* x = (int*) a;	// must type-cast
	printf("%d\n", *x);
}

void* return_from() {
	int x = 40;
	return (void*) (&x);	// hehehehehe, if you know, you know :)
}

void* return_from_escape_to_heap() {
	int *x = (int*) malloc(sizeof(int));
	*x = 60;
	HEAP_REF[current_at++] = x;
	return (void*) x;
}

int main() {

	int x = 10;
	// test(&x);
```

**Explanation**:
- Defines C routine(s): `test()`, `return_from()`, `return_from_escape_to_heap()`, `main()`, `for()`.
- Demonstrates step-by-step logic and runtime behavior.

### 6. Zeroo Fill (`zeroo_fill.c`)

File: [`zeroo_fill.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/zeroo_fill.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
void zero_fill(void* ptr, object_kind_t kind) {
	switch (kind) {
	case INT: {
		obj_int_t* obj = (obj_int_t*) ptr;
		obj->val = INT_ZERO_VAL;
		break;
	}
	case FLOAT: {
		obj_float_t* obj = (obj_float_t*) ptr;
		obj->val = FLOAT_ZERO_VAL;
		break;
	}
	case BOOL: {
		obj_bool_t* obj = (obj_bool_t*) ptr;
		obj->val = BOOL_ZERO_VAL;
		break;
	}
	default:
		return;
	}
}
```

**Explanation**:
- Defines C routine(s): `zero_fill()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`matrix.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/matrix.c) provides or tests `matrix`.
2. [`swap_generic.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/swap_generic.c) provides or tests `swap_generic`.
3. [`swap_int.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/swap_int.c) provides or tests `swap_int`.
4. [`swap_string.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/swap_string.c) provides or tests `swap_string`.

```text
Caller / Test Runner
  ├──> [matrix.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [swap_generic.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [swap_int.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [swap_string.c] (include<stdio.h> include<stdlib.h>...)
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
gcc -Wall -Wextra matrix.c swap_generic.c swap_int.c swap_string.c void_pointer.c zeroo_fill.c -o main
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

1. What is the primary role of the functions/routines demonstrated in `matrix.c`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `matrix.c` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`matrix.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/matrix.c)
- [`swap_generic.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/swap_generic.c)
- [`swap_int.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/swap_int.c)
- [`swap_string.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/swap_string.c)
- [`void_pointer.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/void_pointer.c)
- [`zeroo_fill.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/pointer/zeroo_fill.c)
