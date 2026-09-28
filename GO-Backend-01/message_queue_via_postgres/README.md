<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Backend-01 / message queue via postgres

## Overview

This directory explores **GO-Backend-01 / message queue via postgres** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Backend-01 / message queue via postgres in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`mq.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/message_queue_via_postgres/mq.go) | Go | // rq "github.com/rabbitmq/amqp091-go" |

## Concepts

### 1. Mq (`mq.go`)

File: [`mq.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/message_queue_via_postgres/mq.go)  
// rq "github.com/rabbitmq/amqp091-go".

Relevant code excerpt:

```go
func createSchema(ctx context.Context, conn *sqlx.DB, name string) {
	schemaCreateQuery := fmt.Sprintf(`create schema if not exists %s`, name)
	_, err := conn.ExecContext(ctx, schemaCreateQuery)
	if err != nil {
		panic(err)
	}
}

func listSchema(ctx context.Context, conn *sqlx.DB) {
	var schemas []string
	err := conn.SelectContext(ctx, &schemas, `
        SELECT schema_name
        FROM information_schema.schemata
        WHERE schema_name NOT IN ('pg_catalog', 'information_schema')
          AND schema_name NOT LIKE 'pg_toast%'
        ORDER BY schema_name
    `)
	if err != nil {
		panic(err)
	}

	fmt.Println("Schemas::List::")
```

**Explanation**:
- Implements function(s): `main()`, `migration()`, `createSchema()`, `listSchema()`, `listSchemaTables()`, `createEnum()`, `listEnums()`, `createTable()`, `createEnqueueFunction()`, `enqueue()`, `createDequeueFunction()`, `dequeue()`, `listFunctions()`, `main()`.
- Defines C routine(s): `main()`, `createSchema()`, `listSchema()`, `listSchemaTables()`, `createEnum()`, `listEnums()`, `createTable()`, `createEnqueueFunction()`, `enqueue()`, `createDequeueFunction()`, `dequeue()`, `listFunctions()`, `main()`, `func()`.
- Defines Go struct(s): `EnumInfo`, `QueueRow`, `Payload`, `FunctionInfo`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. Entry into [`mq.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/message_queue_via_postgres/mq.go) (`main` or top-level routine).
2. Initializes internal state / variables and executes core logic.
3. Processes inputs, performs transformations or system calls, and outputs results.

```text
main() [mq.go]
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
go run mq.go
```

## Important Details

- **Memory & Scope**: Notice variable allocation, pointer indirection, and lifetime across function boundaries.
- **Error & Return Handling**: Always verify return values and errors before proceeding to prevent panics or undefined behavior.
- **Resource Management**: Check that open files, network sockets, database handles, and heap allocations are properly closed or freed.
- **AMQP Exchange Routing**: RabbitMQ directs messages via direct, topic (wildcard `*`, `#`), or fanout exchange bindings with publisher confirms.
- **Server-Sent Events (SSE)**: Unidirectional HTTP streaming with `Content-Type: text/event-stream`, utilizing `http.Flusher` to push instant updates to browser clients.
- **PostgreSQL Transactional Queues**: Implements persistent message queuing using `SELECT ... FOR UPDATE SKIP LOCKED` or advisory locks to prevent concurrent worker collisions.

## Common Mistakes

- Assuming implicit synchronization or thread safety where none is provided.
- Ignoring error return values or missing zero-value edge cases.
- Misunderstanding pass-by-value versus pointer/reference semantics for complex types.
- Failing to cast `http.ResponseWriter` to `http.Flusher` in SSE endpoints before attempting to stream data chunks.
- Leaving database transactions open or omitting rollback on error paths, holding row locks indefinitely.

## Language Notes

- **Language Features**: Written using idiomatic Go paradigms.
- **Go Web Stack**: `net/http` handles each inbound connection in an isolated lightweight goroutine.

## Related Concepts

- Data Structures & Algorithms in Go
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `mq.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `mq.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`mq.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/message_queue_via_postgres/mq.go)
