<!-- AUTO-GENERATED FROM SOURCE CODE -->
# C / Socket Programming

## Overview

This directory explores **C / Socket Programming** in **C**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of C / Socket Programming in C.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`setsockopt_client_id.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/10 - Socket Programming/setsockopt_client_id.c) | C | include<stdio.h> include<stdlib.h> |
| [`socket.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/10 - Socket Programming/socket.c) | C | include<stdio.h> include<stdlib.h> |

## Concepts

### 1. Setsockopt Client Id (`setsockopt_client_id.c`)

File: [`setsockopt_client_id.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/10 - Socket Programming/setsockopt_client_id.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
struct sockaddr_storage addr;
} client_t;

void handle_conn(client_t client) {

    // set recv() timeout on client socket fd
    struct timeval tv = {.tv_sec = 5, .tv_usec = 0};    // 5seconds
    if( setsockopt(client.fd, SOL_SOCKET, SO_RCVTIMEO, &tv, sizeof(tv)) < 0) {
        perror("setsockopt SO_RCVTIMEO");
        close(client.fd);
        return;
    }

    // ipv4, client address print
    if (client.addr.ss_family == AF_INET) {
        // // Method - 01
        // struct sockaddr_in temp_client;
        // memcpy(&temp_client, &client.addr, sizeof(temp_client));
        // char client_ip[INET_ADDRSTRLEN];
        // inet_ntop(AF_INET, &temp_client->sin_addr, client_ip, INET_ADDRSTRLEN);
        // printf("[CLIENT] %s:%d\n", client_ip, htons(temp_client->sin_port));
```

**Explanation**:
- Defines C routine(s): `handle_conn()`, `if()`, `if()`, `main()`, `while()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Socket (`socket.c`)

File: [`socket.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/10 - Socket Programming/socket.c)  
include<stdio.h> include<stdlib.h>.

Relevant code excerpt:

```c
int main() {

	// create a socket endpoint
	int fd = socket(AF_INET, SOCK_STREAM, 0);
	if (fd < 0) {
		perror("socket create");
		return 0;
	}
	printf("%d\n", fd);

	// struct sockaddr_storage addr; // ipv4 + ipv6

	struct sockaddr_in addr;
	socklen_t addr_len = sizeof(addr);
	// bzero(&addr, sizeof(addr));	// fill with zero

	memset(&addr, 0, sizeof(addr));

	addr.sin_family = AF_INET;	// ipv4
	addr.sin_port = htons(8080);	// port -> endianess handles
	addr.sin_addr.s_addr = INADDR_ANY;	// accept any local interface
```

**Explanation**:
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`setsockopt_client_id.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/10 - Socket Programming/setsockopt_client_id.c) provides or tests `setsockopt_client_id`.
2. [`socket.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/10 - Socket Programming/socket.c) provides or tests `socket`.

```text
Caller / Test Runner
  ├──> [setsockopt_client_id.c] (include<stdio.h> include<stdlib.h>...)
  ├──> [socket.c] (include<stdio.h> include<stdlib.h>...)
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
gcc -Wall -Wextra setsockopt_client_id.c socket.c -o main
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

1. What is the primary role of the functions/routines demonstrated in `setsockopt_client_id.c`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `setsockopt_client_id.c` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`setsockopt_client_id.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/10 - Socket Programming/setsockopt_client_id.c)
- [`socket.c`](file:///Users/labib0x9/Desktop/Programming-Notes/C/10 - Socket Programming/socket.c)
