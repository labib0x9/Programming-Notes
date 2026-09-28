<!-- AUTO-GENERATED FROM SOURCE CODE -->
# DSA / stack

## Overview

This directory explores **DSA / stack** in **C**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of DSA / stack in C.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`generic_stack.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/stack/generic_stack.c) | C | include<stdio.h> include<stdlib.h> |
| [`stack.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/stack/stack.c) | C | include<stdio.h> include<stdlib.h> |

## Concepts

### 1. Generic Stack (`generic_stack.c`)

File: [`generic_stack.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/stack/generic_stack.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
void printStack(stack* st) {
	for (int i = 0; i < st->len; i++) {
		printf("%d ", *(int*) st->data[i]);
	}
	printf("\n");
}

void freeStack(stack* st) {
	if (!st) {
		return;
	}

	if (st->data) {
		// for (int i = 0; i < st->len; i++) {
		// 	if (st->heapTrack[i]) free(st->data[i]);
		// }
		free(st->data);
	}
	// if (st->heapTrack) free(st->heapTrack);
	free(st);

	// for (int i = 0; i < track_ptr; i++) {
```

**Explanation**:
- Defines C routine(s): `newStack()`, `isEmpty()`, `isFull()`, `resize()`, `push()`, `push()`, `pop()`, `printStack()`, `freeStack()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Stack (`stack.c`)

File: [`stack.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/stack/stack.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
void printStack(stack* st) {
	for (int i = 0; i < st->len; i++) {
		printf("%d ", st->data[i]);
	}
	printf("\n");
}

void freeStack(stack* st) {
	if (!st) {
		return;
	}
	if (st->data) {
		free(st->data);
	}
	free(st);
}

int main() {

	stack* st = newStack(0);
	if (st == NULL) return 0;
```

**Explanation**:
- Defines C routine(s): `newStack()`, `isEmpty()`, `isFull()`, `resize()`, `push()`, `pop()`, `printStack()`, `freeStack()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`generic_stack.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/stack/generic_stack.c) provides or tests `generic_stack`.
2. [`stack.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/stack/stack.c) provides or tests `stack`.

```text
Caller / Test Runner
  ├──> [generic_stack.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [stack.c] (include<stdio.h> include<stdlib.h>...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Compile with GCC
gcc -Wall -Wextra generic_stack.c stack.c -o main
./main
```

## Important Details

- **Memory & Scope**: Notice variable allocation, pointer indirection, and lifetime across function boundaries.
- **Error & Return Handling**: Always verify return values and errors before proceeding to prevent panics or undefined behavior.
- **Resource Management**: Check that open files, network sockets, database handles, and heap allocations are properly closed or freed.
- **Algorithm Efficiency**: Optimized for time and space complexity with deterministic memory footprints.
- **Two-Pointer & In-Place Algorithms**: Modifies data in place ($O(1)$ auxiliary space) via swap and sliding indices.

## Common Mistakes

- Assuming implicit synchronization or thread safety where none is provided.
- Ignoring error return values or missing zero-value edge cases.
- Misunderstanding pass-by-value versus pointer/reference semantics for complex types.
- Off-by-one errors in index boundaries during binary search, segment tree splitting, or array rotation.

## Language Notes

- **Language Features**: Written using idiomatic C paradigms.
- **Data Structures**: Focuses on memory locality, cache-friendly array traversals, and pointer-linked tree node structures.

## Related Concepts

- Data Structures & Algorithms in C
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `generic_stack.c`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `generic_stack.c` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`generic_stack.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/stack/generic_stack.c)
- [`stack.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/stack/stack.c)
