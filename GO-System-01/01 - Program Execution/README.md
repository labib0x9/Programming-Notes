<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-System-01 / Program Execution

## Overview

This directory explores **GO-System-01 / Program Execution** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-System-01 / Program Execution in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`basic.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/basic.go) | Go | io.Write implements stores output in buf |
| [`child_pid.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/child_pid.go) | Go | output via terminal / stdio Parent PID |
| [`error_code.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/error_code.go) | Go | output via terminal / stdio cmd.Stdin = os.Stdin |
| [`ex01.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/ex01.go) | Go | Explain what is happening here, why the exeecution hangs ?? |
| [`ex02.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/ex02.go) | Go | Explain what is happening here, why the execution stops ?? With - Ignore SITTOU Signal |
| [`ex03.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/ex03.go) | Go | Demonstrates main entry point and workflow for ex03 |
| [`foregound_trnasfer.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/foregound_trnasfer.go) | Go | Demonstrates main entry point and workflow for foregound trnasfer |
| [`get_pid.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/get_pid.go) | Go | Demonstrates main entry point and workflow for get pid |
| [`monitor.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/monitor.go) | Go | // set timeout for 10sec ctx, cancel := context.WithTimeout( |
| [`move_child_to_foregoround.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/move_child_to_foregoround.go) | Go | Demonstrates main entry point and workflow for move child to foregoround |
| [`multiple_run.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/multiple_run.go) | Go | Implements `Run, Start, Wait` logic for multiple run |
| [`new_pgid.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/new_pgid.go) | Go | HERE IS SOME CONFUSION>> NEED TO CLEARR.... PID(./sum)= 50662 |
| [`output.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/output.go) | Go | io.Reader implements // I/O mechanism |
| [`process_state.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/process_state.go) | Go | Demonstrates main entry point and workflow for process state |
| [`sh_cmd.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/sh_cmd.go) | Go | output via terminal / stdio cmd.Stdin = os.Stdin |
| [`timeout.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/timeout.go) | Go | set timeout for 10sec |

## Concepts

### 1. Basic (`basic.go`)

File: [`basic.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/basic.go)  
io.Write implements stores output in buf.

Relevant code excerpt:

```go
func main() {
	cmd := exec.Command("./sum") // sum is a ext file in the current dir with no arg

	// io.Write implements
	// stores output in buf
	out := strings.Builder{}

	// io.Reader implements
	in := strings.NewReader("1 2\n")

	// I/O mechanism
	cmd.Stdout = &out
	cmd.Stdin = in

	// runs a program and wait to finish
	// Run() = Start() + Wait()
	if err := cmd.Run(); err != nil {
		fmt.Println("Run Error" + err.Error())
		return
	}

	fmt.Println(out.String())
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Child Pid (`child_pid.go`)

File: [`child_pid.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/child_pid.go)  
output via terminal / stdio Parent PID.

Relevant code excerpt:

```go
func main() {

	cmd := exec.Command("./sum")

	// output via terminal / stdio
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Parent PID
	fmt.Println("Parent PID=", os.Getpid())

	// // No way to get the child pid
	// if err := cmd.Run(); err != nil {
	// 	fmt.Println("Run error:" + err.Error())
	// 	return
	// }

	if err := cmd.Start(); err != nil {
		fmt.Println("Start error:" + err.Error())
		return
	}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Error Code (`error_code.go`)

File: [`error_code.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/error_code.go)  
output via terminal / stdio cmd.Stdin = os.Stdin.

Relevant code excerpt:

```go
func main() {

	cmd := exec.Command("sh", "-c", "pwd")

	// output via terminal / stdio
	// cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// // Parent PID
	// fmt.Println("Parent PID=", os.Getpid())

	if err := cmd.Run(); err != nil {
		fmt.Println("Run error:" + err.Error())
		return
	}

	
	// Handle error code
	err := cmd.Wait()
	if errExit, ok := err.(*exec.ExitError); ok {
		fmt.Println("Wait error:" + err.Error())
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Ex01 (`ex01.go`)

File: [`ex01.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/ex01.go)  
Explain what is happening here, why the exeecution hangs ??.

Relevant code excerpt:

```go
func main() {

	cmd1 := exec.Command("./sum")
	cmd1.Stdin = os.Stdin
	cmd1.Stdout = os.Stdout
	cmd1.Stderr = os.Stderr

	cmd1.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	if err := cmd1.Start(); err != nil {
		panic(err)
	}

	fmt.Println("PID=", cmd1.Process.Pid)

	err := cmd1.Wait()
	if errExit, ok := err.(*exec.ExitError); ok {
		fmt.Println("Run Error=", err.Error())
		fmt.Println("Exit Error=", errExit.ExitCode())
	} else {
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 5. Ex02 (`ex02.go`)

File: [`ex02.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/ex02.go)  
Explain what is happening here, why the execution stops ?? With - Ignore SITTOU Signal.

Relevant code excerpt:

```go
func main() {

	// signal.Ignore(syscall.SIGTTOU)

	cmd1 := exec.Command("./sum")
	cmd1.Stdin = os.Stdin
	cmd1.Stdout = os.Stdout
	cmd1.Stderr = os.Stderr

	cmd1.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	if err := cmd1.Start(); err != nil {
		panic(err)
	}

	childPID := cmd1.Process.Pid
	fmt.Println("Child Process ID=", childPID)

	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Assembly procedure(s): `Child` with DOS interrupt interactions.
- Demonstrates step-by-step logic and runtime behavior.

### 6. Ex03 (`ex03.go`)

File: [`ex03.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/ex03.go)  
Demonstrates main entry point and workflow for ex03.

Relevant code excerpt:

```go
func main() {

	cmd1 := exec.Command("./sum")
	cmd1.Stdin = os.Stdin
	cmd1.Stdout = os.Stdout
	cmd1.Stderr = os.Stderr

	cmd1.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	var err error
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer tty.Close()

	if err := cmd1.Start(); err != nil {
		panic(err)
	}
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Assembly procedure(s): `Child, child` with DOS interrupt interactions.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`basic.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/basic.go) provides or tests `basic`.
2. [`child_pid.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/child_pid.go) provides or tests `child_pid`.
3. [`error_code.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/error_code.go) provides or tests `error_code`.
4. [`ex01.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/ex01.go) provides or tests `ex01`.

```text
Caller / Test Runner
  ├──> [basic.go] (io.Write implements stores output i...)
  ├──> [child_pid.go] (output via terminal / stdio Parent ...)
  ├──> [error_code.go] (output via terminal / stdio cmd.Std...)
  ├──> [ex01.go] (Explain what is happening here, why...)
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
go run basic.go
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

1. What is the primary role of the functions/routines demonstrated in `basic.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `basic.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`basic.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/basic.go)
- [`child_pid.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/child_pid.go)
- [`error_code.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/error_code.go)
- [`ex01.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/ex01.go)
- [`ex02.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/ex02.go)
- [`ex03.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/ex03.go)
- [`foregound_trnasfer.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/foregound_trnasfer.go)
- [`get_pid.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/get_pid.go)
- [`monitor.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/monitor.go)
- [`move_child_to_foregoround.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/move_child_to_foregoround.go)
- [`multiple_run.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/multiple_run.go)
- [`new_pgid.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/new_pgid.go)
- [`output.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/output.go)
- [`process_state.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/process_state.go)
- [`sh_cmd.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/sh_cmd.go)
- [`timeout.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-System-01/01 - Program Execution/timeout.go)
