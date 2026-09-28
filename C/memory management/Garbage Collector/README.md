<!-- AUTO-GENERATED FROM SOURCE CODE -->
# C / memory management / Garbage Collector

## Overview

This directory explores **C / memory management / Garbage Collector** in **C**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of C / memory management / Garbage Collector in C.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`gc.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/gc.c) | C | I don't have any idea how to implement, need to study more and more.... / Incomplete Code |
| [`object.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/object.c) | C | include<stdio.h> include<stdlib.h> |
| [`object_final.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/object_final.c) | C | include<stdio.h> include<stdlib.h> |
| [`reference_counting_gc.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/reference_counting_gc.c) | C | include<stdio.h> include<stdlib.h> |

## Concepts

### 1. Gc (`gc.c`)

File: [`gc.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/gc.c)  
I don't have any idea how to implement, need to study more and more.... / Incomplete Code.

Relevant code excerpt:

```c
struct GarbageObject* next;
} garbaje_object_t;

typedef struct GarbageCollector {
	garbaje_object_t* root;
} gc_t;

void* gc_alloc(gc_t* gc, size_t size) {
	void* ptr = malloc(size);
	return ptr;
}

int main() {

	int* arr = (int*) gc_alloc(NULL, sizeof(int) * 10);
	memset(arr, 0, sizeof(int) * 10);

	arr[0] = 10;
	arr[1] = 30;

	for (int i = 0; i < 10; i++) {
		printf("%d ", arr[i]);
```

**Explanation**:
- Defines C routine(s): `gc_alloc()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Object (`object.c`)

File: [`object.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/object.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
void PrintObject(object_t* obj) {
	if (obj == NULL) {
		return;
	}
	switch (obj->Kind) {
		case INT:
			printf("%d\n", obj->Data.v_int);
			break;
		case FLOAT:
			printf("%0.3f\n", obj->Data.v_float);
			break;
		default:
			printf("unknown data types\n");
			break;
	}
}

int main() {

	object_t* obj1 = NewObjectInt(10);
	object_t* obj2 = NewObjectFloat(20.012);
```

**Explanation**:
- Defines C routine(s): `NewObjectInt()`, `NewObjectFloat()`, `PrintObject()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Object Final (`object_final.c`)

File: [`object_final.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/object_final.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
struct Object {
	object_kind_t Kind;	// what kind of object
	object_data_t Data;	// Actual value
};

object_t* NewObjectInt(int val) {
	object_t* obj = (object_t*) malloc(sizeof(object_t));
	if (obj == NULL) {
		return NULL;
	}
	obj->Kind = INT;
	obj->Data.v_int = val;
	return obj;
}

object_t* NewObjectFloat(float val) {
	object_t* obj = (object_t*) malloc(sizeof(object_t));
	if (obj == NULL) {
		return NULL;
	}
	obj->Kind = FLOAT;
	obj->Data.v_float = val;
```

**Explanation**:
- Defines C routine(s): `NewObjectInt()`, `NewObjectFloat()`, `NewObjectString()`, `NewObjectArray()`, `NewObjectVector3()`, `set_array()`, `get_array()`, `len()`, `add()`, `PrintObject()`, `freeObject()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Reference Counting Gc (`reference_counting_gc.c`)

File: [`reference_counting_gc.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/reference_counting_gc.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#include<stdbool.h>
#include<unistd.h>

int main() {

	return 0;
}
```

**Explanation**:
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`gc.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/gc.c) provides or tests `gc`.
2. [`object.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/object.c) provides or tests `object`.
3. [`object_final.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/object_final.c) provides or tests `object_final`.
4. [`reference_counting_gc.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/reference_counting_gc.c) provides or tests `reference_counting_gc`.

```text
Caller / Test Runner
  ├──> [gc.c] (I don't have any idea how to implem...)
  ├──> [object.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [object_final.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [reference_counting_gc.c] (include<stdio.h> include<stdlib.h>...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Compile with GCC
gcc -Wall -Wextra gc.c object.c object_final.c reference_counting_gc.c -o main
./main
```

## Important Details

- **Memory & Scope**: Notice variable allocation, pointer indirection, and lifetime across function boundaries.
- **Error & Return Handling**: Always verify return values and errors before proceeding to prevent panics or undefined behavior.
- **Resource Management**: Check that open files, network sockets, database handles, and heap allocations are properly closed or freed.
- **Object Pool & Free-List Allocation**: Pre-allocates memory blocks in contiguous slabs, reducing `malloc`/`free` system call overhead and avoiding heap fragmentation.
- **Custom Garbage Collection**: Manages object headers with reference counts or mark-and-sweep tracking over dynamically created object unions.
- **Stack vs Heap Lifetime**: Variables allocated on the stack are popped upon function return; heap allocations persist until explicitly freed or garbage-collected.

## Common Mistakes

- Assuming implicit synchronization or thread safety where none is provided.
- Ignoring error return values or missing zero-value edge cases.
- Misunderstanding pass-by-value versus pointer/reference semantics for complex types.
- Returning pointers to local stack-allocated variables, resulting in undefined behavior when the stack frame is overwritten.
- Double freeing or losing pointer references to allocated blocks, resulting in memory leaks.

## Language Notes

- **Language Features**: Written using idiomatic C paradigms.
- **C Memory Management**: Requires explicit lifetime ownership and alignment considerations (`sizeof(void*)`, `#pragma pack`).

## Issues / Risks

> File [`gc.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/gc.c) contains experimental incomplete code as noted in comments.

## Related Concepts

- Data Structures & Algorithms in C
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `gc.c`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `gc.c` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`gc.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/gc.c)
- [`object.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/object.c)
- [`object_final.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/object_final.c)
- [`reference_counting_gc.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/Garbage Collector/reference_counting_gc.c)
