<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-System-01 / Container

## Overview

This directory explores **GO-System-01 / Container** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-System-01 / Container in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`chroot.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/02-Container/chroot.go) | Go | self-re-execution pattern |
| [`container-01.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/02-Container/container-01.go) | Go | //go:build linux "fmt" |

## Concepts

### 1. Chroot (`chroot.go`)

File: [`chroot.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/02-Container/chroot.go)  
self-re-execution pattern.

Relevant code excerpt:

```go
func logger(reason string, err error) {
	if err == nil {
		return
	}
	panic(fmt.Sprintf("[%s] %v", reason, err))
}

// self-re-execution pattern
func parent() {
	exePath, err := os.Executable()
	logger("parent exe path", err)

	cmd := exec.Command(exePath, "child")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUSER | // NEWUSER for rootless containersyscall.CLONE_NEWPID |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWNS,
```

**Explanation**:
- Implements function(s): `logger()`, `parent()`, `child()`, `main()`.
- Defines C routine(s): `logger()`, `parent()`, `child()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Container-01 (`container-01.go`)

File: [`container-01.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/02-Container/container-01.go)  
//go:build linux "fmt".

Relevant code excerpt:

```go
// //go:build linux

// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"path/filepath"
// 	"strings"
// 	"syscall"

// 	"golang.org/x/sys/unix"
// )

// func PrintPidAndMntNS(pid int, tag string) string {
// 	fmt.Println(tag, "PID=", pid)

// 	// mnt namespace id read
// 	mntNSPath := fmt.Sprintf("/proc/%d/ns/mnt", pid)

// 	mntNSOut, err := os.Readlink(mntNSPath)
```

**Explanation**:
- Implements function(s): `PrintPidAndMntNS()`, `pivotRoot()`, `parent()`, `child()`, `main()`.
- Defines C routine(s): `parent()`, `child()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`chroot.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/02-Container/chroot.go) provides or tests `chroot`.
2. [`container-01.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/02-Container/container-01.go) provides or tests `container-01`.

```text
Caller / Test Runner
  ├──> [chroot.go] (self-re-execution pattern...)
  ├──> [container-01.go] (//go:build linux "fmt"...)
  └──> Execution Completion / Assertion
```

## System Interaction

```text
Host Parent Process (PID 1234)
       │
       ├──> syscall.ForkExec() / os/exec [Cloneflags: CLONE_NEWPID | CLONE_NEWNS]
       │         │
       │         ▼
       │    Child Process (Host PID 5678, Container PID 1)
       │         │
       │         ├──> chroot("/path/to/rootfs")
       │         ├──> mount("proc", "/proc", "proc", 0, "")
       │         └──> execve("/bin/sh")
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run chroot.go
```

## Important Details

- **Memory & Scope**: Notice variable allocation, pointer indirection, and lifetime across function boundaries.
- **Error & Return Handling**: Always verify return values and errors before proceeding to prevent panics or undefined behavior.
- **Resource Management**: Check that open files, network sockets, database handles, and heap allocations are properly closed or freed.
- **Linux Namespaces & Re-exec**: Leverages `syscall.SysProcAttr` with `CLONE_NEWPID` and `CLONE_NEWNS` to spawn child processes inside isolated process ID and mount namespaces.
- **Filesystem Isolation (`chroot` / `pivot_root`)**: Jails a process within a sub-tree filesystem, restricting filesystem access to the container root.
- **Process Group & Foreground Management**: Manipulates `Setpgid` and terminal control (`TIOCSPGRP`) to move background child processes to the foreground.
- **Docker API Client**: Interacts with the Docker daemon socket (`/var/run/docker.sock`) to pull images and manage container lifecycles programmatically.

## Common Mistakes

- Assuming implicit synchronization or thread safety where none is provided.
- Ignoring error return values or missing zero-value edge cases.
- Misunderstanding pass-by-value versus pointer/reference semantics for complex types.
- Running Linux namespace syscalls on macOS or Windows; namespaces are Linux-kernel-specific features requiring a Linux host or VM.
- Forgetting to mount a private `/proc` filesystem inside a new PID namespace, causing `ps` to still reflect host processes.

## Language Notes

- **Language Features**: Written using idiomatic Go paradigms.
- **Go System Syscalls**: The `syscall` and `golang.org/x/sys/unix` packages expose raw POSIX/Linux kernel interfaces.

## Related Concepts

- Data Structures & Algorithms in Go
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `chroot.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `chroot.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`chroot.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/02-Container/chroot.go)
- [`container-01.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/02-Container/container-01.go)
