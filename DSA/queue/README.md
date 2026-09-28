<!-- AUTO-GENERATED FROM SOURCE CODE -->
# DSA / queue

## Overview

This directory explores **DSA / queue** in **C**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of DSA / queue in C.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`queue.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/queue/queue.c) | C | Implements `isEmpty, isFull, push` in C |
| [`queue_dynamic.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/queue/queue_dynamic.c) | C | include<stdio.h> include<stdlib.h> |
| [`queue_linked_list.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/queue/queue_linked_list.c) | C | include<stdio.h> include<stdlib.h> |

## Concepts

### 1. Queue (`queue.c`)

File: [`queue.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/queue/queue.c)  
Implements `isEmpty, isFull, push` in C.

Relevant code excerpt:

```c
void print() {
   for(int i = 0; i < count; i++) {
      printf("%d ", queue[(i + readAt) % FIXED_QUEUE_ELEM]);
   }
   printf("\n");
}
```

**Explanation**:
- Defines C routine(s): `isEmpty()`, `isFull()`, `push()`, `pop()`, `print()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Queue Dynamic (`queue_dynamic.c`)

File: [`queue_dynamic.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/queue/queue_dynamic.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
for (int i = 0; i < q->len; i++) {
      newQ[i] = q->q[(q->readAt + i) % q->cap];
   }

   // fix the logic...
   // memmove(newQ, q->q + q->readAt, sizeof(int) * (q->len - q->readAt));
   // memmove(newQ + (q->len - q->readAt), q->q, sizeof(int) * (q->readAt + 1));

   // update
   free(q->q);
   q->q = newQ;
   q->cap = newCap;
   q->insertAt = q->len;
   q->readAt = 0;

   return true;
}

bool push(queue* q, int val) {
   if (isFull(q)) {
      bool ok = resize(q);
      if (!ok) return ok;
```

**Explanation**:
- Defines C routine(s): `init_queue()`, `isFull()`, `isEmpty()`, `resize()`, `push()`, `pop()`, `free_queue()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Queue Linked List (`queue_linked_list.c`)

File: [`queue_linked_list.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/queue/queue_linked_list.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
struct node* next;
} node;

typedef struct {
   node* tail;
   node* head;
} queue;

queue* init_queue() {
   queue* q = (queue*) malloc(sizeof(queue));   
   q->head = q->tail = NULL;
   return q;
}

bool push(queue* q, int val) {
   // new node
   node* newNode = (node*) malloc(sizeof(node));
   if (newNode == NULL) {  // malloc fails
      return false;
   }
   newNode->val = val;
   newNode->next = NULL;
```

**Explanation**:
- Defines C routine(s): `init_queue()`, `push()`, `pop()`, `if()`, `peek()`, `free_queue()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`queue.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/queue/queue.c) provides or tests `queue`.
2. [`queue_dynamic.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/queue/queue_dynamic.c) provides or tests `queue_dynamic`.
3. [`queue_linked_list.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/queue/queue_linked_list.c) provides or tests `queue_linked_list`.

```text
Caller / Test Runner
  ├──> [queue.c] (Implements `isEmpty, isFull, push` ...)
  ├──> [queue_dynamic.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [queue_linked_list.c] (include<stdio.h> include<stdlib.h>...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Compile with GCC
gcc -Wall -Wextra queue.c queue_dynamic.c queue_linked_list.c -o main
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

1. What is the primary role of the functions/routines demonstrated in `queue.c`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `queue.c` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`queue.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/queue/queue.c)
- [`queue_dynamic.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/queue/queue_dynamic.c)
- [`queue_linked_list.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/queue/queue_linked_list.c)
