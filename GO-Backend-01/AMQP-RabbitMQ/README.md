<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Backend-01 / AMQP-RabbitMQ

## Overview

This directory explores **GO-Backend-01 / AMQP-RabbitMQ** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Backend-01 / AMQP-RabbitMQ in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`confirm_publish.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/confirm_publish.go) | Go | exchange name + type queue name + routing key |
| [`direct_exchange.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/direct_exchange.go) | Go | exchange name + type queue name + routing key |
| [`fanout_exchange.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/fanout_exchange.go) | Go | exchange name + type queue name + routing key |
| [`frefetch_message.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/frefetch_message.go) | Go | how many concurrent msg to handle at once need to undestand the global parameter |
| [`max_channel_per_connection.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/max_channel_per_connection.go) | Go | Demonstrates main entry point and workflow for max channel per connection |
| [`producer-consumer.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/producer-consumer.go) | Go | exchange name + type queue name + routing key |
| [`topic_exchange.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/topic_exchange.go) | Go | exchange name + type queue name + routing key |

## Concepts

### 1. Confirm Publish (`confirm_publish.go`)

File: [`confirm_publish.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/confirm_publish.go)  
exchange name + type queue name + routing key.

Relevant code excerpt:

```go
func declareQueue(conn *amq.Connection, exchange string, queue, route string) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// _, err = ch.QueueDelete(queue, false, false, false)
	// var amqpErr *amq.Error
	// if !(errors.As(err, &amqpErr) && amqpErr.Code == amq.NotFound) {
	// 	return err
	// }

	q, err := ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	fmt.Println(q.Name)

	if err := ch.QueueBind(q.Name, route, exchange, false, nil); err != nil {
		return err
```

**Explanation**:
- Implements function(s): `declareQueue()`, `declareExchange()`, `publish()`, `main()`.
- Defines C routine(s): `publish()`, `main()`, `func()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Direct Exchange (`direct_exchange.go`)

File: [`direct_exchange.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/direct_exchange.go)  
exchange name + type queue name + routing key.

Relevant code excerpt:

```go
func declareQueue(ch *amq.Channel, exchange, typee string, queue, route string) error {
	if err := ch.ExchangeDeclare(exchange, typee, false, false, false, false, nil); err != nil {
		return err
	}

	q, err := ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	fmt.Println(q.Name)

	if err := ch.QueueBind(q.Name, route, exchange, false, nil); err != nil {
		return err
	}
	return nil
}

func consume(conn *amq.Connection, queue string, consumer string) {
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(consumer, "ERROR:", err)
```

**Explanation**:
- Implements function(s): `declareQueue()`, `consume()`, `publish()`, `main()`.
- Defines C routine(s): `consume()`, `publish()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Fanout Exchange (`fanout_exchange.go`)

File: [`fanout_exchange.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/fanout_exchange.go)  
exchange name + type queue name + routing key.

Relevant code excerpt:

```go
func declareQueue(ch *amq.Channel, exchange string, queue, route string) error {
	q, err := ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	fmt.Println(q.Name)

	if err := ch.QueueBind(q.Name, route, exchange, false, nil); err != nil {
		return err
	}
	return nil
}

// remove exchange if presents, then create a new one
func declareExchange(conn *amq.Connection, exchange, typee string) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
```

**Explanation**:
- Implements function(s): `declareQueue()`, `declareExchange()`, `consume()`, `publish()`, `main()`.
- Defines C routine(s): `consume()`, `publish()`, `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Frefetch Message (`frefetch_message.go`)

File: [`frefetch_message.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/frefetch_message.go)  
how many concurrent msg to handle at once need to undestand the global parameter.

Relevant code excerpt:

```go
func consume(conn *amq.Connection, queue string, consumer string) {
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(consumer, "ERROR:", err)
		return
	}
	defer ch.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// how many concurrent msg to handle at once
	// need to undestand the global parameter
	err = ch.Qos(10, 0, false) // per consumer limit
	// ch.Qos(10, 0, true)  // per channel limit
	if err != nil {
		fmt.Println(consumer, "ERROR:", err)
		return
	}

	msgs, err := ch.ConsumeWithContext(ctx, queue, consumer, true, false, false, false, nil)
	if err != nil {
		fmt.Println(consumer, "ERROR:", err)
```

**Explanation**:
- Implements function(s): `consume()`.
- Defines C routine(s): `consume()`.
- Demonstrates step-by-step logic and runtime behavior.

### 5. Max Channel Per Connection (`max_channel_per_connection.go`)

File: [`max_channel_per_connection.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/max_channel_per_connection.go)  
Demonstrates main entry point and workflow for max channel per connection.

Relevant code excerpt:

```go
func main() {

	conn, err := amq.DialConfig(
		"amqp://guest:guest@localhost:5672",
		amq.Config{
			ChannelMax: 10,
			Recovery:   &amq.Recovery{},
		})
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var wg sync.WaitGroup
	for i := range 20 {
		wg.Add(1)
		go func(id int, wg *sync.WaitGroup) {
			defer wg.Done()
			ch, err := conn.Channel()
			if err != nil {
				fmt.Println("ID:", id, "Error", err, "Coonection", conn.IsClosed())
				return
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`, `func()`.
- Demonstrates step-by-step logic and runtime behavior.

### 6. Producer-Consumer (`producer-consumer.go`)

File: [`producer-consumer.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/producer-consumer.go)  
exchange name + type queue name + routing key.

Relevant code excerpt:

```go
func declareQueue(ch *amq.Channel, exchange, typee string, queue, route string) error {
	// ----	// Exchange
	if err := ch.ExchangeDeclare(exchange, typee, false, false, false, false, nil); err != nil {
		return err
	}

	// ---- // Queue
	q, err := ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	fmt.Println(q.Name)

	// ---- // Bind
	if err := ch.QueueBind(q.Name, route, exchange, false, nil); err != nil {
		return err
	}
	return nil
}

func main() {
```

**Explanation**:
- Implements function(s): `declareQueue()`, `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`confirm_publish.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/confirm_publish.go) provides or tests `confirm_publish`.
2. [`direct_exchange.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/direct_exchange.go) provides or tests `direct_exchange`.
3. [`fanout_exchange.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/fanout_exchange.go) provides or tests `fanout_exchange`.
4. [`frefetch_message.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/frefetch_message.go) provides or tests `frefetch_message`.

```text
Caller / Test Runner
  ├──> [confirm_publish.go] (exchange name + type queue name + r...)
  ├──> [direct_exchange.go] (exchange name + type queue name + r...)
  ├──> [fanout_exchange.go] (exchange name + type queue name + r...)
  ├──> [frefetch_message.go] (how many concurrent msg to handle a...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run confirm_publish.go
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

1. What is the primary role of the functions/routines demonstrated in `confirm_publish.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `confirm_publish.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`confirm_publish.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/confirm_publish.go)
- [`direct_exchange.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/direct_exchange.go)
- [`fanout_exchange.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/fanout_exchange.go)
- [`frefetch_message.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/frefetch_message.go)
- [`max_channel_per_connection.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/max_channel_per_connection.go)
- [`producer-consumer.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/producer-consumer.go)
- [`topic_exchange.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/AMQP-RabbitMQ/topic_exchange.go)
