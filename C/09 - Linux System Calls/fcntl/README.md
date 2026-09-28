<!-- AUTO-GENERATED FROM SOURCE CODE -->
# C / Linux System Calls / fcntl

## Overview

This directory explores **C / Linux System Calls / fcntl** in **C**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of C / Linux System Calls / fcntl in C.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`non_blocking_io.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/fcntl/non_blocking_io.c) | C | include<stdio.h> include<unistd.h> |
| [`override_or_combine_flags.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/fcntl/override_or_combine_flags.c) | C | include<stdio.h> include<unistd.h> |

## Concepts

### 1. Non Blocking Io (`non_blocking_io.c`)

File: [`non_blocking_io.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/fcntl/non_blocking_io.c)  
include<stdio.h> include<unistd.h>.

Relevant code excerpt:

```c
int main() {
    int fd = socket(AF_INET, SOCK_STREAM, 0);
    if (fd == -1) {
        perror("socket");
        return 1;
    }
    if (set_fd_nonblock(fd) == -1) {
        perror("set_fd_nonblock");
        close(fd);
        return 1;
    }
    printf("Socket set to non-blocking mode\n");
    close(fd);
    return 0;
}
```

**Explanation**:
- Defines C routine(s): `set_fd_nonblock()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Override Or Combine Flags (`override_or_combine_flags.c`)

File: [`override_or_combine_flags.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/fcntl/override_or_combine_flags.c)  
include<stdio.h> include<unistd.h>.

Relevant code excerpt:

```c
int main() {    

    // pipe creation
    int ok = pipe(pipe_fd);
    if (ok == -1) {
        perror("Pipe");
        return 0;
    }

    printf("=========================\n");

    // Combined flag
    int flags = fcntl(pipe_fd[0], F_GETFL, 0);
    printf("READER FLAG= %d\n", flags);
    fcntl(pipe_fd[0], F_SETFL, flags | O_NONBLOCK);
    printf("READER FLAG= %d\n", fcntl(pipe_fd[0], F_GETFL, 0));


    // Flag override, can be dangerous
    flags = fcntl(pipe_fd[1], F_GETFL, 0);
    printf("WRITER FLAG= %d\n", flags);
    fcntl(pipe_fd[1], F_SETFL, O_NONBLOCK);
```

**Explanation**:
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`non_blocking_io.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/fcntl/non_blocking_io.c) provides or tests `non_blocking_io`.
2. [`override_or_combine_flags.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/fcntl/override_or_combine_flags.c) provides or tests `override_or_combine_flags`.

```text
Caller / Test Runner
  ├──> [non_blocking_io.c] (include<stdio.h> include<unistd.h>...)
  ├──> [override_or_combine_flags.c] (include<stdio.h> include<unistd.h>...)
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
gcc -Wall -Wextra non_blocking_io.c override_or_combine_flags.c -o main
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

1. What is the primary role of the functions/routines demonstrated in `non_blocking_io.c`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `non_blocking_io.c` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`non_blocking_io.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/fcntl/non_blocking_io.c)
- [`override_or_combine_flags.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/09 - Linux System Calls/fcntl/override_or_combine_flags.c)
