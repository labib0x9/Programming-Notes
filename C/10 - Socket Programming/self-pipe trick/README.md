<!-- AUTO-GENERATED FROM SOURCE CODE -->
# C / Socket Programming / self-pipe trick

## Overview

This directory explores **C / Socket Programming / self-pipe trick** in **C**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of C / Socket Programming / self-pipe trick in C.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`basic.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/10 - Socket Programming/self-pipe trick/basic.c) | C | include<stdio.h> include<unistd.h> |

## Concepts

### 1. Basic (`basic.c`)

File: [`basic.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/10 - Socket Programming/self-pipe trick/basic.c)  
include<stdio.h> include<unistd.h>.

Relevant code excerpt:

```c
int main() {    

    // pipe creation
    int pipe_fd[2];
    int ok = pipe(pipe_fd);
    if (ok == -1) {
        perror("Pipe");
        return 0;
    }

    // Data flows, fd[1] -> fd[0]
    printf("Read end, fd[0] = %d\n", pipe_fd[0]);
    printf("Write end, fd[1] = %d\n", pipe_fd[1]);

    // write to pipe
    int n = write(pipe_fd[1], "Hello Labib\n", 12);
    printf("%d bytes are written to fd[1]\n", n);

    // read from pipe
    char buf[16];
    n = read(pipe_fd[0], &buf, sizeof(buf));
    printf("%d bytes are read from fd[0]\n", n);
```

**Explanation**:
- Defines C routine(s): `main()`.
- Assembly procedure(s): `child, same` with DOS interrupt interactions.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. Entry into [`basic.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/10 - Socket Programming/self-pipe trick/basic.c) (`main` or top-level routine).
2. Initializes internal state / variables and executes core logic.
3. Processes inputs, performs transformations or system calls, and outputs results.

```text
main() [basic.c]
  │
  ├── Initialize state / data
  ├── Execute core operations
  └── Output result / Terminate
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
gcc -Wall -Wextra basic.c -o main
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

1. What is the primary role of the functions/routines demonstrated in `basic.c`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `basic.c` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`basic.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/10 - Socket Programming/self-pipe trick/basic.c)
