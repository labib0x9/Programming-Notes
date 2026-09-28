<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-System-01 / Docker Api

## Overview

This directory explores **GO-System-01 / Docker Api** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-System-01 / Docker Api in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`basics.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/03-Docker Api/basics.go) | Go | Demonstrates main entry point and workflow for basics |
| [`create_a_container.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/03-Docker Api/create_a_container.go) | Go | Demonstrates main entry point and workflow for create a container |
| [`interractive.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/03-Docker Api/interractive.go) | Go | Demonstrates main entry point and workflow for interractive |
| [`pull_an_image.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/03-Docker Api/pull_an_image.go) | Go | Demonstrates main entry point and workflow for pull an image |

## Concepts

### 1. Basics (`basics.go`)

File: [`basics.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/03-Docker Api/basics.go)  
Demonstrates main entry point and workflow for basics.

Relevant code excerpt:

```go
func main() {

	imageN := "alpine:3.20"
	containerId := ""

	// --- // Connect to dockerhub
	client, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Pull image if not present...
	imgResp, err := client.ImageInspect(context.Background(), imageN)
	if err != nil {
		if strings.Contains(err.Error(), "No such image:") {
			reader, err := client.ImagePull(context.Background(), imageN, image.PullOptions{})
			if err != nil {
				panic(err)
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Create A Container (`create_a_container.go`)

File: [`create_a_container.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/03-Docker Api/create_a_container.go)  
Demonstrates main entry point and workflow for create a container.

Relevant code excerpt:

```go
func main() {

	// --- // Connect to dockerhub
	client, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// --- // Create Container
	containerN := "container-vv2"
	resp, err := client.ContainerCreate(context.Background(),
		&container.Config{
			Image: "alpine:3.20", // image has to be present.. or pull
			Cmd:   []string{"echo", "Hello host"},
		},
		nil, nil, nil, containerN,
	)
	if err != nil {
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Interractive (`interractive.go`)

File: [`interractive.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/03-Docker Api/interractive.go)  
Demonstrates main entry point and workflow for interractive.

Relevant code excerpt:

```go
func main() {

	imageN := "alpine:3.20"
	containerId := ""

	// --- // Connect to dockerhub
	client, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Pull image if not present...
	imgResp, err := client.ImageInspect(context.Background(), imageN)
	if err != nil {
		if strings.Contains(err.Error(), "No such image:") {
			reader, err := client.ImagePull(context.Background(), imageN, image.PullOptions{})
			if err != nil {
				panic(err)
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Pull An Image (`pull_an_image.go`)

File: [`pull_an_image.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/03-Docker Api/pull_an_image.go)  
Demonstrates main entry point and workflow for pull an image.

Relevant code excerpt:

```go
func main() {

	// --- // Connect to dockerhub
	client, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	reader, err := client.ImagePull(context.Background(), "alpine:3.20", image.PullOptions{})
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	io.Copy(os.Stdout, reader)
}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`basics.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/03-Docker Api/basics.go) provides or tests `basics`.
2. [`create_a_container.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/03-Docker Api/create_a_container.go) provides or tests `create_a_container`.
3. [`interractive.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/03-Docker Api/interractive.go) provides or tests `interractive`.
4. [`pull_an_image.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/03-Docker Api/pull_an_image.go) provides or tests `pull_an_image`.

```text
Caller / Test Runner
  ├──> [basics.go] (Demonstrates main entry point and w...)
  ├──> [create_a_container.go] (Demonstrates main entry point and w...)
  ├──> [interractive.go] (Demonstrates main entry point and w...)
  ├──> [pull_an_image.go] (Demonstrates main entry point and w...)
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
go run basics.go
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

1. What is the primary role of the functions/routines demonstrated in `basics.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `basics.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`basics.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/03-Docker Api/basics.go)
- [`create_a_container.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/03-Docker Api/create_a_container.go)
- [`interractive.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/03-Docker Api/interractive.go)
- [`pull_an_image.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/03-Docker Api/pull_an_image.go)
