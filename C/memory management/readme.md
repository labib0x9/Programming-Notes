<!-- AUTO-GENERATED FROM SOURCE CODE -->
# C / memory management

## Overview

This directory explores **C / memory management** in **C**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of C / memory management in C.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`freelist_object_pool.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/freelist_object_pool.c) | C | include<stdio.h> include<stdlib.h> |
| [`memory.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/memory.c) | C | include<stdio.h> include<stdlib.h> |
| [`memory_pool.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/memory_pool.c) | C | include<stdio.h> include<stdlib.h> |
| [`object_pool.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/object_pool.c) | C | include<stdio.h> include<stdlib.h> |
| [`stack_frame.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/stack_frame.c) | C | include<stdio.h> include<stdlib.h> |
| [`stack_frame_1.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/stack_frame_1.c) | C | include<stdio.h> include<stdlib.h> |
| [`stack_overflow.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/stack_overflow.c) | C | include<stdio.h> include<stdlib.h> |
| [`static.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/static.c) | C | include<stdio.h> include<stdlib.h> |

## Concepts

### 1. Freelist Object Pool (`freelist_object_pool.c`)

File: [`freelist_object_pool.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/freelist_object_pool.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
struct PoolElem* next;
} pool_elem_t;

const int POOL_SIZE = 10;
pool_elem_t object_pool[POOL_SIZE] = {{0}};
pool_elem_t* freeList = NULL;

void init_pool() {
	for (int i = 0; i + 1 < POOL_SIZE; i++) {
		object_pool[i].next = &object_pool[i + 1];
	}
	object_pool[POOL_SIZE - 1].next = NULL;
	freeList = &object_pool[0];
}

pair_t* Get() {
	if (freeList == NULL) {	// no free object
		return NULL;
	}
	pair_t* obj = &(freeList->obj);	// extract object
	freeList = freeList->next;	// move to the next
	return obj;
```

**Explanation**:
- Defines C routine(s): `init_pool()`, `Get()`, `Put()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Memory (`memory.c`)

File: [`memory.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/memory.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
#include<stdio.h>
#include<stdlib.h>
#include<string.h>

int32_t main() {

    // malloc

    // allocate memory for 1 int on heap..
    int *ptr = malloc(sizeof(int)); // garbage value
    if (ptr == NULL) {
        // malloc fails
    }
    // casting is required in c++
    // int *ptr1 = (int*) malloc(sizeof(int));
    // if (ptr1 == NULL) {
    //     // malloc fails
    // }

    memset(ptr, 0, sizeof(int)); // set to 0x00
    // memset(ptr, 1, sizeof(int)); // garbage value
    // sets bytes to 0x01 so ptr = 0x01010101 = some value (garbage value)
```

**Explanation**:
- Defines C routine(s): `main()`, `if()`, `if()`, `for()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Memory Pool (`memory_pool.c`)

File: [`memory_pool.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/memory_pool.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#include<stdbool.h>
#include<unistd.h>

// linked list
typedef struct Block {
    struct Block* next;
} block_t;

int main() {

	return 0;
}
```

**Explanation**:
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Object Pool (`object_pool.c`)

File: [`object_pool.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/object_pool.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
for (int i = 0; i < POOL_SIZE; i++) {
		if (object_pool[i].allocated) continue;
		object_pool[i].allocated = true;
		return &object_pool[i].obj;	// return the address
	}
	return NULL; // no free object
}

bool Put(pair_t* obj) {
	if (obj == NULL) return false;
	uintptr_t pool_addr = (uintptr_t) object_pool;	// address of the pool
	uintptr_t obj_addr = (uintptr_t) obj - (uintptr_t) offsetof(pool_elem_t, obj);	// object address inside pool, offset is used to find the initial memory address of where the obj is located in memory..
	size_t size = sizeof(pool_elem_t);	// size of the pool element struct with padding.
	int idx = (obj_addr - pool_addr) / size;	// index of the obj in pool
	if (idx >= POOL_SIZE || idx < 0) return false;
    if ((uintptr_t)(idx * size) + pool_addr != obj_addr) return false; // must be a multiplier..
	object_pool[idx].allocated = false;	// pool claims that memory
	return true;
}

int main() {
```

**Explanation**:
- Defines C routine(s): `Get()`, `Put()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 5. Stack Frame (`stack_frame.c`)

File: [`stack_frame.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/stack_frame.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
void call_me() {
	int x = 10;
	printf("Call Me: (top) = %p\n", __builtin_frame_address(0));	// call_me() stack frame address -> lower than the main()
	printf("Call Me: (top) = %lu\n", (uintptr_t) __builtin_frame_address(0));
	printf("Call Me: (ret) = %p\n", __builtin_return_address(0));	// .text section
	printf("Call Me: (ret) = %lu\n", (uintptr_t) __builtin_return_address(0));
}

int main() {

	printf("Main: (top) = %p\n", __builtin_frame_address(0));	// main() stack frame address
	printf("Main: (top) = %lu\n", (uintptr_t) __builtin_frame_address(0));
	call_me();

	return 0;
}
```

**Explanation**:
- Defines C routine(s): `call_me()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 6. Stack Frame 1 (`stack_frame_1.c`)

File: [`stack_frame_1.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/stack_frame_1.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
void stack_frame() {
	static void* prev_sp = NULL;
	void* current_sp = __builtin_frame_address(0);
	long diff = (char*) prev_sp - (char*) current_sp;

	if (prev_sp) {
		printf("Diff = %ld\n", diff);
	}

	if (current_sp) {
		prev_sp = current_sp;
	}
}

void A() {
	int x = 10;
	printf("A() = %p\n", __builtin_frame_address(0));
	stack_frame();
}

void B() {
	int y = 20;
```

**Explanation**:
- Defines C routine(s): `stack_frame()`, `A()`, `B()`, `C()`, `D()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`freelist_object_pool.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/freelist_object_pool.c) provides or tests `freelist_object_pool`.
2. [`memory.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/memory.c) provides or tests `memory`.
3. [`memory_pool.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/memory_pool.c) provides or tests `memory_pool`.
4. [`object_pool.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/object_pool.c) provides or tests `object_pool`.

```text
Caller / Test Runner
  ├──> [freelist_object_pool.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [memory.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [memory_pool.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [object_pool.c] (include<stdio.h> include<stdlib.h>...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Compile with GCC
gcc -Wall -Wextra freelist_object_pool.c memory.c memory_pool.c object_pool.c stack_frame.c stack_frame_1.c stack_overflow.c static.c -o main
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

## Related Concepts

- Data Structures & Algorithms in C
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `freelist_object_pool.c`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `freelist_object_pool.c` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`freelist_object_pool.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/freelist_object_pool.c)
- [`memory.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/memory.c)
- [`memory_pool.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/memory_pool.c)
- [`object_pool.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/object_pool.c)
- [`stack_frame.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/stack_frame.c)
- [`stack_frame_1.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/stack_frame_1.c)
- [`stack_overflow.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/stack_overflow.c)
- [`static.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/memory management/static.c)
