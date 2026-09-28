<!-- AUTO-GENERATED FROM SOURCE CODE -->
# DSA / Trie

## Overview

This directory explores **DSA / Trie** in **C, Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of DSA / Trie in C, Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`trie.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/Trie/trie.c) | C | https://codeforces.com/contest/1902/problem/E include<stdio.h> |
| [`trie.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/Trie/trie.go) | Go | https://codeforces.com/contest/1902/problem/E x, _ := strconv.ParseFloat(strings.TrimSpace(s), 64) |

## Concepts

### 1. Trie (`trie.c`)

File: [`trie.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/Trie/trie.c)  
https://codeforces.com/contest/1902/problem/E include<stdio.h>.

Relevant code excerpt:

```c
struct Node {
	char exit;
	int freq;
	struct Node* child[26];
};

struct Node* NewNode() {
	struct Node* node = malloc(sizeof(struct Node));
	node->exit = 0;
	node->freq = 0;
	for (int i = 0; i < 26; i++) {
		node->child[i] = NULL;
	}
	return node;
}

struct Node* Trie() {
	return NewNode();
};

void insert(struct Node* root, const char* s, int n) {
	struct Node* cur = root;
```

**Explanation**:
- Defines C routine(s): `NewNode()`, `Trie()`, `insert()`, `LCP()`, `freeTrie()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Trie (`trie.go`)

File: [`trie.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/Trie/trie.go)  
https://codeforces.com/contest/1902/problem/E x, _ := strconv.ParseFloat(strings.TrimSpace(s), 64).

Relevant code excerpt:

```go
func ReadInt() int {
	s, _ := reader.ReadString('\n')
	x, _ := strconv.Atoi(strings.TrimSpace(s))
	// x, _ := strconv.ParseFloat(strings.TrimSpace(s), 64)
	return x
}

func ReadPairs() (int64, int64, int64, int64, int64) {
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	u, _ := strconv.ParseInt(parts[0], 10, 64)
	v, _ := strconv.ParseInt(parts[1], 10, 64)
	x, _ := strconv.ParseInt(parts[2], 10, 64)
	y, _ := strconv.ParseInt(parts[3], 10, 64)
	z, _ := strconv.ParseInt(parts[4], 10, 64)

	return u, v, x, y, z
}

func ReadArray() []string {
	s, _ := reader.ReadString('\n')
	fields := strings.Fields(s)
```

**Explanation**:
- Implements function(s): `ReadInt()`, `ReadPairs()`, `ReadArray()`, `ReadString()`, `Write()`, `NewNode()`, `Insert()`, `LCP()`, `dfs()`, `NewTrie()`, `solve()`, `main()`.
- Defines C routine(s): `Write()`, `dfs()`, `solve()`, `main()`.
- Defines Go struct(s): `node`, `Trie`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`trie.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/Trie/trie.c) provides or tests `trie`.
2. [`trie.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/Trie/trie.go) provides or tests `trie`.

```text
Caller / Test Runner
  ├──> [trie.c] (https://codeforces.com/contest/1902...)
  ├──> [trie.go] (https://codeforces.com/contest/1902...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run trie.c
```

```bash
# Compile with GCC
gcc -Wall -Wextra trie.c -o main
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

- **Language Features**: Written using idiomatic C, Go paradigms.
- **Data Structures**: Focuses on memory locality, cache-friendly array traversals, and pointer-linked tree node structures.

## Related Concepts

- Data Structures & Algorithms in C, Go
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `trie.c`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `trie.c` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`trie.c`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/Trie/trie.c)
- [`trie.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/Trie/trie.go)
