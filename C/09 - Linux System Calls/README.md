<!-- AUTO-GENERATED FROM SOURCE CODE -->
# C / Linux System Calls

## Overview

This directory explores **C / Linux System Calls** in **C**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of C / Linux System Calls in C.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`fstat.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/fstat.c) | C | include<stdio.h> include<sys/stat.h> |
| [`mmap_readonly.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/mmap_readonly.c) | C | include<stdio.h> include<sys/stat.h> |
| [`mmap_readwrite.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/mmap_readwrite.c) | C | include<stdio.h> include<sys/stat.h> |

## Concepts

### 1. Fstat (`fstat.c`)

File: [`fstat.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/fstat.c)  
include<stdio.h> include<sys/stat.h>.

Relevant code excerpt:

```c
int main() {

    // File descriptor
    int fd = open("test.txt", O_RDONLY);
    if (fd == -1) {
        // error
        perror("open");
        return 0;
    }
    
    // File status
    struct stat status;
    if (fstat(fd, &status) == -1) {
        // error
        perror("read status");
        close(fd);
        return 0;
    }

    // File size
    printf("%jd\n", (intmax_t)status.st_size);
```

**Explanation**:
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Mmap Readonly (`mmap_readonly.c`)

File: [`mmap_readonly.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/mmap_readonly.c)  
include<stdio.h> include<sys/stat.h>.

Relevant code excerpt:

```c
int main() {

    // File descriptor
    int fd = open("test.txt", O_RDONLY);
    if (fd == -1) {
        // error
        perror("open");
        return 0;
    }
    
    // File status
    struct stat status;
    if (fstat(fd, &status) == -1) {
        // error
        perror("read status");
        close(fd);
        return 0;
    }

    // empty files
    if (status.st_size == 0) {
        perror("mapped area cannot be empty");
```

**Explanation**:
- Defines C routine(s): `main()`, `if()`, `if()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Mmap Readwrite (`mmap_readwrite.c`)

File: [`mmap_readwrite.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/mmap_readwrite.c)  
include<stdio.h> include<sys/stat.h>.

Relevant code excerpt:

```c
int main() {

    // File descriptor
    int fd = open("test.txt", O_RDWR);
    if (fd == -1) {
        // error
        perror("open");
        return 0;
    }
    
    // File status
    struct stat status;
    if (fstat(fd, &status) == -1) {
        // error
        perror("read status");
        close(fd);
        return 0;
    }

    // empty files
    if (status.st_size == 0) {
        perror("mapped area cannot be empty");
```

**Explanation**:
- Defines C routine(s): `main()`, `if()`, `if()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`fstat.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/fstat.c) provides or tests `fstat`.
2. [`mmap_readonly.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/mmap_readonly.c) provides or tests `mmap_readonly`.
3. [`mmap_readwrite.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/mmap_readwrite.c) provides or tests `mmap_readwrite`.

```text
Caller / Test Runner
  ├──> [fstat.c] (include<stdio.h> include<sys/stat.h...)
  ├──> [mmap_readonly.c] (include<stdio.h> include<sys/stat.h...)
  ├──> [mmap_readwrite.c] (include<stdio.h> include<sys/stat.h...)
  └──> Execution Completion / Assertion
```

## System Interaction

```text
User Application (Event Loop)
       │
       ├──> socket() / fcntl(O_NONBLOCK)
       │
       ├──> kevent() / epoll_wait() ──[Syscall Boundary]──> Kernel Event Demux
       │                                                         │
       │<── [FD Ready Notification: EVFILT_READ / EPOLLIN] ──────┘
       │
       └──> non-blocking read() / write()
```

## Compilation & Execution

```bash
# Compile with GCC
gcc -Wall -Wextra fstat.c mmap_readonly.c mmap_readwrite.c -o main
./main
```

## Important Details

- **Memory & Scope**: Notice variable allocation, pointer indirection, and lifetime across function boundaries.
- **Error & Return Handling**: Always verify return values and errors before proceeding to prevent panics or undefined behavior.
- **Resource Management**: Check that open files, network sockets, database handles, and heap allocations are properly closed or freed.
- **Non-Blocking I/O (`O_NONBLOCK`)**: Configured via `fcntl(fd, F_SETFL, flags | O_NONBLOCK)`, preventing `read()` or `accept()` from putting the thread into sleep state.
- **Event Demultiplexing**: Uses `kqueue` (macOS/BSD) or `epoll` (Linux) to monitor readiness events on multiple file descriptors simultaneously in a single event loop.
- **Self-Pipe Trick**: Uses a UNIX pipe `pipe(fd)` to wake up a blocking `select()`/`poll()`/`kqueue()` event loop from an asynchronous signal handler or another worker thread.

## Common Mistakes

- Assuming implicit synchronization or thread safety where none is provided.
- Ignoring error return values or missing zero-value edge cases.
- Misunderstanding pass-by-value versus pointer/reference semantics for complex types.
- Calling `fcntl(fd, F_SETFL, O_NONBLOCK)` directly without reading existing flags via `F_GETFL`, which accidentally wipes out existing file status flags.
- Failing to handle `EAGAIN` / `EWOULDBLOCK` errors on non-blocking sockets when no data is currently available in the receive buffer.

## Language Notes

- **Language Features**: Written using idiomatic C paradigms.
- **POSIX Sockets**: Sockets are standard integer file descriptors managed by the OS kernel network stack.

## Related Concepts

- Data Structures & Algorithms in C
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `fstat.c`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `fstat.c` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`fstat.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/fstat.c)
- [`mmap_readonly.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/mmap_readonly.c)
- [`mmap_readwrite.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/mmap_readwrite.c)
