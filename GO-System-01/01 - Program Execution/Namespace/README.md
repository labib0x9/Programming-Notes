<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-System-01 / Program Execution / Namespace

## Overview

This directory explores **GO-System-01 / Program Execution / Namespace** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-System-01 / Program Execution / Namespace in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`mount_inside_child.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/Namespace/mount_inside_child.go) | Go | NOT WORKING... For OS Thread Locking |
| [`re_exec_change_namespace.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/Namespace/re_exec_change_namespace.go) | Go | mnt namespace id read |

## Concepts

### 1. Mount Inside Child (`mount_inside_child.go`)

File: [`mount_inside_child.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/Namespace/mount_inside_child.go)  
NOT WORKING... For OS Thread Locking.

Relevant code excerpt:

```go
func main() {

	// For OS Thread Locking
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Read parent mnt namespace
	parentNs, err := os.Open("/proc/self/ns/mnt")
	if err != nil {
		panic(err)
	}
	defer parentNs.Close()

	// cmd...
	cmd1 := exec.Command("sh")
	cmd1.Stdin = os.Stdin
	cmd1.Stdout = os.Stdout
	cmd1.Stderr = os.Stderr

	cmd1.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS,
	}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Re Exec Change Namespace (`re_exec_change_namespace.go`)

File: [`re_exec_change_namespace.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/Namespace/re_exec_change_namespace.go)  
mnt namespace id read.

Relevant code excerpt:

```go
func PrintPidAndMntNS(pid int, tag string) {
	fmt.Println(tag, "PID=", pid)

	// mnt namespace id read
	mntNSPath := fmt.Sprintf("/proc/%d/ns/mnt", pid)

	mntNSOut, err := os.Readlink(mntNSPath)
	if err != nil {
		panic(err)
	}

	mntNSID := strings.TrimPrefix(mntNSOut, "mnt:[")
	mntNSID = strings.TrimSuffix(mntNSID, "]")
	fmt.Println(mntNSPath, "=", mntNSID)
}

// self-re-execution pattern
func parent() {
	cmd := exec.Command("/proc/self/exe", "child")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
```

**Explanation**:
- Implements function(s): `PrintPidAndMntNS()`, `parent()`, `child()`, `main()`.
- Defines C routine(s): `PrintPidAndMntNS()`, `parent()`, `child()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`mount_inside_child.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/Namespace/mount_inside_child.go) provides or tests `mount_inside_child`.
2. [`re_exec_change_namespace.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/Namespace/re_exec_change_namespace.go) provides or tests `re_exec_change_namespace`.

```text
Caller / Test Runner
  ├──> [mount_inside_child.go] (NOT WORKING... For OS Thread Lockin...)
  ├──> [re_exec_change_namespace.go] (mnt namespace id read...)
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
go run mount_inside_child.go
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

1. What is the primary role of the functions/routines demonstrated in `mount_inside_child.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `mount_inside_child.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`mount_inside_child.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/Namespace/mount_inside_child.go)
- [`re_exec_change_namespace.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/Namespace/re_exec_change_namespace.go)
