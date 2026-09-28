<!-- AUTO-GENERATED FROM SOURCE CODE -->
# DSA / array

## Overview

This directory explores **DSA / array** in **C, Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of DSA / array in C, Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`check_if_array_is_sorted.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/check_if_array_is_sorted.c) | C | Merge Two Sorted Array. include<stdio.h> |
| [`deletion.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/deletion.c) | C | Array Deletion. include<stdio.h> |
| [`deletion.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/deletion.go) | Go | n elements deletion |
| [`insertion.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/insertion.c) | C | Array Insertion. include<stdio.h> |
| [`insertion.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/insertion.go) | Go | insertion, n elements, n + 1 room |
| [`merge.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/merge.c) | C | Merge Array. include<stdio.h> |
| [`merge.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/merge.go) | Go | merge two array merge :) |
| [`merge_two_sorted_array.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/merge_two_sorted_array.c) | C | Merge Two Sorted Array. include<stdio.h> |
| [`reverse.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/reverse.c) | C | Array Reverse. include<stdio.h> |
| [`reverse.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/reverse.go) | Go | n elements two pointer reverse |
| [`rotate_mod.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/rotate_mod.c) | C | include<stdio.h> Left - Roatation, k = 1 |
| [`rotate_reverse.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/rotate_reverse.c) | C | include<stdio.h> Left - Rotation |
| [`rotate_reverse.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/rotate_reverse.go) | Go | Left - Rotatioin |

## Concepts

### 1. Check If Array Is Sorted (`check_if_array_is_sorted.c`)

File: [`check_if_array_is_sorted.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/check_if_array_is_sorted.c)  
Merge Two Sorted Array. include<stdio.h>.

Relevant code excerpt:

```c
void print_array(int arr[], int len) {
    for (int i = 0; i < len; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}

// check if the array is sorted (ascending order)
bool is_sorted(int arr[], int len) {
    for (int i = 0; i + 1 < len; i++) {
        if (arr[i] > arr[i + 1]) { return false; }
    }
    return true;
}

// descending order
bool is_dsorted(int arr[], int len) {
    for (int i = 0; i + 1 < len; i++) {
        if (arr[i] < arr[i + 1]) { return false; }
    }
    return true;
}
```

**Explanation**:
- Defines C routine(s): `print_array()`, `is_sorted()`, `is_dsorted()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Deletion (`deletion.c`)

File: [`deletion.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/deletion.c)  
Array Deletion. include<stdio.h>.

Relevant code excerpt:

```c
void print_array(int arr[], int len) {
    for (int i = 0; i < len; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}

void delete(int arr[], int idx, int* len) {
    // if (*len >= capacity) { return; }
    memmove(&arr[idx], &arr[idx + 1], sizeof(int) * (*len - idx - 1));
    (*len)--;
}

int main() {

    int cap = 10, len = 10;
    int arr[cap];
    for (int i = 0; i < len; i++) arr[i] = i + 1;

    int idx = 0;
    delete(arr, idx, &len);
    print_array(arr, len);
```

**Explanation**:
- Defines C routine(s): `print_array()`, `delete()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Deletion (`deletion.go`)

File: [`deletion.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/deletion.go)  
n elements deletion.

Relevant code excerpt:

```go
package main

import "fmt"

func main() {
	// n elements
	n := 7
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	fmt.Println(arr)

	// deletion
	idx := 0
	copy(arr[idx:], arr[idx+1:]) // left shift
	arr = arr[:n-1]              // remove one room

	fmt.Println(arr)
}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Insertion (`insertion.c`)

File: [`insertion.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/insertion.c)  
Array Insertion. include<stdio.h>.

Relevant code excerpt:

```c
void print_array(int arr[], int len) {
    for (int i = 0; i < len; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}

void insert(int arr[], int idx, int x, int* len, int capacity) {
    if (*len >= capacity) { return; }
    memmove(&arr[idx + 1], &arr[idx], sizeof(int) * (*len - idx));
    (*len)++;
    arr[idx] = x;
}

int main() {

    int cap = 10, len = 5;
    int arr[cap];
    for (int i = 0; i < len; i++) arr[i] = i + 1;

    int idx = 0, x = 100;
    insert(arr, idx, x, &len, cap);
```

**Explanation**:
- Defines C routine(s): `print_array()`, `insert()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 5. Insertion (`insertion.go`)

File: [`insertion.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/insertion.go)  
insertion, n elements, n + 1 room.

Relevant code excerpt:

```go
package main

import "fmt"

func main() {
	// insertion, n elements, n + 1 room
	n := 7
	arr := make([]int, n+1)

	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	fmt.Println(arr)

	idx, x := 0, 100
	copy(arr[idx+1:], arr[idx:]) // right shift
	arr[idx] = x                 // insert

	fmt.Println(arr)
}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 6. Merge (`merge.c`)

File: [`merge.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/merge.c)  
Merge Array. include<stdio.h>.

Relevant code excerpt:

```c
void print_array(int arr[], int len) {
    for (int i = 0; i < len; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}

int main() {

    int lenA = 3, lenB = 4;
    int arr[lenA], brr[lenB];

    for (int i = 0; i < lenA; i++) arr[i] = i + 1;
    for (int i = 0; i < lenB; i++)  brr[i] = i + lenA + 1;

    print_array(arr, lenA);
    print_array(brr, lenB);

    // merge array
    // using loop
    int lenC = lenA + lenB;
    int merged_arr[lenC];
```

**Explanation**:
- Defines C routine(s): `print_array()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`check_if_array_is_sorted.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/check_if_array_is_sorted.c) provides or tests `check_if_array_is_sorted`.
2. [`deletion.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/deletion.c) provides or tests `deletion`.
3. [`deletion.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/deletion.go) provides or tests `deletion`.
4. [`insertion.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/insertion.c) provides or tests `insertion`.

```text
Caller / Test Runner
  ├──> [check_if_array_is_sorted.c] (Merge Two Sorted Array. include<std...)
  ├──> [deletion.c] (Array Deletion. include<stdio.h>...)
  ├──> [deletion.go] (n elements deletion...)
  ├──> [insertion.c] (Array Insertion. include<stdio.h>...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run check_if_array_is_sorted.c
```

```bash
# Compile with GCC
gcc -Wall -Wextra check_if_array_is_sorted.c deletion.c insertion.c merge.c merge_two_sorted_array.c reverse.c rotate_mod.c rotate_reverse.c -o main
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

- **Language Features**: Written using idiomatic C, Go paradigms.
- **Data Structures**: Focuses on memory locality, cache-friendly array traversals, and pointer-linked tree node structures.

## Related Concepts

- Data Structures & Algorithms in C, Go
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `check_if_array_is_sorted.c`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `check_if_array_is_sorted.c` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`check_if_array_is_sorted.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/check_if_array_is_sorted.c)
- [`deletion.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/deletion.c)
- [`deletion.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/deletion.go)
- [`insertion.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/insertion.c)
- [`insertion.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/insertion.go)
- [`merge.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/merge.c)
- [`merge.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/merge.go)
- [`merge_two_sorted_array.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/merge_two_sorted_array.c)
- [`reverse.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/reverse.c)
- [`reverse.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/reverse.go)
- [`rotate_mod.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/rotate_mod.c)
- [`rotate_reverse.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/rotate_reverse.c)
- [`rotate_reverse.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/rotate_reverse.go)
