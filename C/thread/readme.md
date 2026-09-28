<!-- AUTO-GENERATED FROM SOURCE CODE -->
# C / thread

## Overview

This directory explores **C / thread** in **C**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of C / thread in C.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`communite_through_thread.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/communite_through_thread.c) | C | include<stdio.h> include<stdlib.h> |
| [`condition_variable.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/condition_variable.c) | C | include<stdio.h> include<stdlib.h> |
| [`create_thread.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/create_thread.c) | C | include<stdio.h> include<unistd.h> |
| [`join_thread.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/join_thread.c) | C | include<stdio.h> include<unistd.h> |
| [`mutex_lock.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/mutex_lock.c) | C | include<stdio.h> include<stdlib.h> |
| [`rece_condition.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/rece_condition.c) | C | include<stdio.h> include<stdlib.h> |
| [`thread_id.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/thread_id.c) | C | include<stdio.h> include<stdlib.h> |

## Concepts

### 1. Communite Through Thread (`communite_through_thread.c`)

File: [`communite_through_thread.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/communite_through_thread.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
for (int i = 0; i <= 10; i++) {
      sleep(1);
      printf("Alice = %d\n", i);
      (*ptr) += 2;
   }
   return ptr; // return result pointer
}

int main() {

   pthread_t alice;
   pthread_create(&alice, NULL, AliceTurn, NULL);

   int* result;
   pthread_join(alice, (void*) &result); // wait alice to finish and we should pass the pointer of resulting variable..

   printf("%d\n", *result);

   return 0;
}
```

**Explanation**:
- Defines C routine(s): `AliceTurn()`, `for()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Condition Variable (`condition_variable.c`)

File: [`condition_variable.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/condition_variable.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
for the slot to free. 
*/

typedef struct Task {
   int a, b;
} Task;

const int QueueCount = 10;
Task queue[QueueCount];
int insertAt = 0, readAt = 0;
int size = 0, finish = 0;

const int MAX_THREAD_NUM = 4;
pthread_mutex_t lock = PTHREAD_MUTEX_INITIALIZER;
pthread_cond_t empty = PTHREAD_COND_INITIALIZER;
pthread_cond_t full = PTHREAD_COND_INITIALIZER;

// Mutex locking to avoid race condition...
void push(Task task) {
   pthread_mutex_lock(&lock); // lock mutex
   while (size == QueueCount) {
      // We will wait here until a task is done, blocking the push ooperation...
```

**Explanation**:
- Defines C routine(s): `push()`, `while()`, `pop()`, `while()`, `executeTask()`, `startThread()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Create Thread (`create_thread.c`)

File: [`create_thread.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/create_thread.c)  
include<stdio.h> include<unistd.h>.

Relevant code excerpt:

```c
for (int i = 0; i <= 10; i++) {
      sleep(1);
      printf("Alice = %d\n", i);
   }
   return NULL;
}

void* BobTurn(void* args) {
   for (int i = 0; i <= 4; i++) {
      sleep(1);
      printf("Bob = %d\n", i);
   }
   return NULL;
}

int main() {

   pthread_t alice;
   pthread_create(&alice, NULL, AliceTurn, NULL);

   pthread_t bob;
   pthread_create(&bob, NULL, BobTurn, NULL);
```

**Explanation**:
- Defines C routine(s): `AliceTurn()`, `BobTurn()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Join Thread (`join_thread.c`)

File: [`join_thread.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/join_thread.c)  
include<stdio.h> include<unistd.h>.

Relevant code excerpt:

```c
for (int i = 0; i <= 10; i++) {
      sleep(1);
      printf("Alice = %d\n", i);
   }
   return NULL;
}

void* BobTurn(void* args) {
   for (int i = 0; i <= 4; i++) {
      sleep(1);
      printf("Bob = %d\n", i);
   }
   return NULL;
}

int main() {

   pthread_t alice;
   pthread_create(&alice, NULL, AliceTurn, NULL);

   pthread_t bob;
   pthread_create(&bob, NULL, BobTurn, NULL);
```

**Explanation**:
- Defines C routine(s): `AliceTurn()`, `BobTurn()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 5. Mutex Lock (`mutex_lock.c`)

File: [`mutex_lock.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/mutex_lock.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
int main() {

   count = 0;

   pthread_t t;
   pthread_create(&t, NULL, Calc, NULL);

   Calc(NULL);

   pthread_join(t, NULL);
   printf("%d\n", count);

   return 0;
}
```

**Explanation**:
- Defines C routine(s): `Calc()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 6. Rece Condition (`rece_condition.c`)

File: [`rece_condition.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/rece_condition.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
int main() {

   count = 0;

   pthread_t alice;
   pthread_create(&alice, NULL, Calc, NULL);

   Calc(NULL);

   pthread_join(alice, NULL);
   printf("%d\n", count);

   return 0;
}
```

**Explanation**:
- Defines C routine(s): `Calc()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`communite_through_thread.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/communite_through_thread.c) provides or tests `communite_through_thread`.
2. [`condition_variable.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/condition_variable.c) provides or tests `condition_variable`.
3. [`create_thread.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/create_thread.c) provides or tests `create_thread`.
4. [`join_thread.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/join_thread.c) provides or tests `join_thread`.

```text
Caller / Test Runner
  ├──> [communite_through_thread.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [condition_variable.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [create_thread.c] (include<stdio.h> include<unistd.h>...)
  ├──> [join_thread.c] (include<stdio.h> include<unistd.h>...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Compile with GCC
gcc -Wall -Wextra communite_through_thread.c condition_variable.c create_thread.c join_thread.c mutex_lock.c rece_condition.c thread_id.c -o main
./main
```

## Important Details

- **Memory & Scope**: Notice variable allocation, pointer indirection, and lifetime across function boundaries.
- **Error & Return Handling**: Always verify return values and errors before proceeding to prevent panics or undefined behavior.
- **Resource Management**: Check that open files, network sockets, database handles, and heap allocations are properly closed or freed.
- **POSIX Threads Synchronization**: `pthread_mutex_t` guarantees mutual exclusion over shared critical sections; `pthread_cond_t` allows threads to sleep until signaled.
- **Thread Lifecycle**: Every thread created with `pthread_create` must either be joined with `pthread_join` or detached with `pthread_detach` to release internal thread state resources.

## Common Mistakes

- Assuming implicit synchronization or thread safety where none is provided.
- Ignoring error return values or missing zero-value edge cases.
- Misunderstanding pass-by-value versus pointer/reference semantics for complex types.
- Condition variable spurious wakeups: checking conditions using `if` instead of `while (!condition) pthread_cond_wait(...)`.
- Data races caused by accessing shared global variables without holding the associated mutex lock.

## Language Notes

- **Language Features**: Written using idiomatic C paradigms.
- **pthreads**: Direct kernel thread mappings on modern operating systems with shared virtual memory address space.

## Related Concepts

- Data Structures & Algorithms in C
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `communite_through_thread.c`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `communite_through_thread.c` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`communite_through_thread.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/communite_through_thread.c)
- [`condition_variable.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/condition_variable.c)
- [`create_thread.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/create_thread.c)
- [`join_thread.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/join_thread.c)
- [`mutex_lock.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/mutex_lock.c)
- [`rece_condition.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/rece_condition.c)
- [`thread_id.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/thread/thread_id.c)
