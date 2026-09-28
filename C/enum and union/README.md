<!-- AUTO-GENERATED FROM SOURCE CODE -->
# C / enum and union

## Overview

This directory explores **C / enum and union** in **C**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of C / enum and union in C.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`enum.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/enum and union/enum.c) | C | include<stdio.h> include<stdlib.h> |
| [`union.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/enum and union/union.c) | C | include<stdio.h> include<stdlib.h> |

## Concepts

### 1. Enum (`enum.c`)

File: [`enum.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/enum and union/enum.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
#include<stdio.h>
#include<stdlib.h>

// you can also use custom values...

typedef enum DaysOfWeek {
    FRIDAY,     // 0
    SATURSAY,   // 1
    SUNDAY,     // 2
    MONDAY,     // 3
} days_of_week_t;

enum {
    PROTOCOL_TCP,
    PROTOCOL_UDP,
};

int32_t main() {

    days_of_week_t day = FRIDAY;

    return 0;
}
```

**Explanation**:
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Union (`union.c`)

File: [`union.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/enum and union/union.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
struct {
        uint8_t r;
        uint8_t g;
        uint8_t b;
        uint8_t a;
    } components;
    uint32_t rgba;
} color_t;

int32_t main() {

    union_type_t a = {.val = 10};

    printf("%d\n", a.val);
    printf("%lld\n", a.err);    // same memory interpret as long long

    color_t x = {.rgba = 2555555};  // hex = 0x00270F13, depending on endianess value will be interpreted.. r = 0x13 or r = 0x00
    printf("%d\n", x.components.r);
    printf("%d\n", x.components.g);
    printf("%d\n", x.components.b);
    printf("%d\n", x.components.a);
```

**Explanation**:
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`enum.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/enum and union/enum.c) provides or tests `enum`.
2. [`union.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/enum and union/union.c) provides or tests `union`.

```text
Caller / Test Runner
  ├──> [enum.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [union.c] (include<stdio.h> include<stdlib.h>...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Compile with GCC
gcc -Wall -Wextra enum.c union.c -o main
./main
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

- **Language Features**: Written using idiomatic C paradigms.

## Related Concepts

- Data Structures & Algorithms in C
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `enum.c`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `enum.c` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`enum.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/enum and union/enum.c)
- [`union.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/enum and union/union.c)
