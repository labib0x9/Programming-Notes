<!-- AUTO-GENERATED FROM SOURCE CODE -->
# DSA / array / leetcode solve

## Overview

This directory explores **DSA / array / leetcode solve** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of DSA / array / leetcode solve in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`905.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/leetcode solve/905.go) | Go | https://leetcode.com/problems/sort-array-by-parity/description use two pointer to solve. |

## Concepts

### 1. 905 (`905.go`)

File: [`905.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/leetcode solve/905.go)  
https://leetcode.com/problems/sort-array-by-parity/description use two pointer to solve..

Relevant code excerpt:

```go
// https://leetcode.com/problems/sort-array-by-parity/description/

// use two pointer to solve.
// one for odd, one for even
// O(n) time complexity
// O(1) memory complexity
func sortArrayByParity(nums []int) []int {
	odd, even := 0, len(nums)-1
	for odd < even {
		for nums[odd]%2 == 0 && odd < even { // skip evens
			odd++
		}
		for nums[even]%2 == 1 && odd < even { // skip odds
			even--
		}
		nums[odd], nums[even] = nums[even], nums[odd]
		odd, even = odd+1, even-1
	}
	return nums
}
```

**Explanation**:
- Implements function(s): `sortArrayByParity()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. Entry into [`905.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/leetcode solve/905.go) (`main` or top-level routine).
2. Initializes internal state / variables and executes core logic.
3. Processes inputs, performs transformations or system calls, and outputs results.

```text
main() [905.go]
  │
  ├── Initialize state / data
  ├── Execute core operations
  └── Output result / Terminate
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run 905.go
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

- **Language Features**: Written using idiomatic Go paradigms.
- **Data Structures**: Focuses on memory locality, cache-friendly array traversals, and pointer-linked tree node structures.

## Related Concepts

- Data Structures & Algorithms in Go
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `905.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `905.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`905.go`](file:///Users/labib0x9/Desktop/Programming-Notes/DSA/array/leetcode solve/905.go)
