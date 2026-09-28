<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Intermediate-01 / Interface / others hizibiji code

## Overview

This directory explores **GO-Intermediate-01 / Interface / others hizibiji code** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Intermediate-01 / Interface / others hizibiji code in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`notification.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface/others_hizibiji_code/notification.go) | Go | Implements `Notifier, Notify, Notify` logic for notification |
| [`payment.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface/others_hizibiji_code/payment.go) | Go | Only rocket payment has discount, How to apply ?? Nothig is here |
| [`temp.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface/others_hizibiji_code/temp.go) | Go | Demonstrates temp implementation mechanics |

## Concepts

### 1. Notification (`notification.go`)

File: [`notification.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface/others_hizibiji_code/notification.go)  
Implements `Notifier, Notify, Notify` logic for notification.

Relevant code excerpt:

```go
type User struct {
	Name string
}

type Admin struct {
	Name string
	Id int
}

type Email struct {
	EmailAddress string
}

type Notification interface {
	Notify(string)
}

func Notifier(n Notification, msg string) {

	// Type assertion - 01
	if admin, ok := n.(Admin); ok {
		msg += fmt.Sprintf(" : Admin id %d", admin.Id)
```

**Explanation**:
- Implements function(s): `Notifier()`, `Notify()`, `Notify()`, `Notify()`, `NotificationManagement()`.
- Defines C routine(s): `Notifier()`, `NotificationManagement()`.
- Defines Go struct(s): `User`, `Admin`, `Email`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Payment (`payment.go`)

File: [`payment.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface/others_hizibiji_code/payment.go)  
Only rocket payment has discount, How to apply ?? Nothig is here.

Relevant code excerpt:

```go
type Paypal struct {
	accountId string
}

func (p Paypal) Pay(amount int) string {
	return fmt.Sprintf("%d amount Paypal is successful for %s", amount, p.accountId)
}

func (p Paypal) CashOut(amount int) {
	// Nothig is here
}

type Bkash struct {
	number string
}

func (b Bkash) Pay(amount int) string {
	return fmt.Sprintf("%d amount Bkash is successful for %s", amount, b.number)
}

func (b Bkash) CashOut(amount int) {
	fmt.Println("Cashout from", b.number)
```

**Explanation**:
- Implements function(s): `Pay()`, `CashOut()`, `Pay()`, `CashOut()`, `Pay()`, `CashOut()`, `Payment()`, `discountCalculate()`, `PaymentMangement()`.
- Defines C routine(s): `Payment()`, `PaymentMangement()`.
- Defines Go struct(s): `Paypal`, `Bkash`, `Rocket`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Temp (`temp.go`)

File: [`temp.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface/others_hizibiji_code/temp.go)  
Demonstrates temp implementation mechanics.

Relevant code excerpt:

```go
type Artifact interface {
	Title() string
	Creators() []string
	CreatedAt() time.Time
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}

type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
```

**Explanation**:
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`notification.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface/others_hizibiji_code/notification.go) provides or tests `notification`.
2. [`payment.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface/others_hizibiji_code/payment.go) provides or tests `payment`.
3. [`temp.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface/others_hizibiji_code/temp.go) provides or tests `temp`.

```text
Caller / Test Runner
  ├──> [notification.go] (Implements `Notifier, Notify, Notif...)
  ├──> [payment.go] (Only rocket payment has discount, H...)
  ├──> [temp.go] (Demonstrates temp implementation me...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run notification.go
```

## Important Details

- **Memory & Scope**: Notice variable allocation, pointer indirection, and lifetime across function boundaries.
- **Error & Return Handling**: Always verify return values and errors before proceeding to prevent panics or undefined behavior.
- **Resource Management**: Check that open files, network sockets, database handles, and heap allocations are properly closed or freed.
- **Interface Representation**: An interface value is represented internally as a two-word pair: `(itab / type_descriptor, data_pointer)`.
- **Laws of Reflection**: 1. Reflection goes from interface value to reflection object (`reflect.TypeOf`, `reflect.ValueOf`). 2. Reflection goes from reflection object to interface value (`v.Interface()`). 3. To modify a reflection object, the value must be settable (`v.CanSet()`).

## Common Mistakes

- Assuming implicit synchronization or thread safety where none is provided.
- Ignoring error return values or missing zero-value edge cases.
- Misunderstanding pass-by-value versus pointer/reference semantics for complex types.
- Calling `.Set()` on a `reflect.Value` created from a non-pointer or unexported struct field, causing a runtime panic.
- Comparing an interface holding a typed nil pointer (`(*MyStruct)(nil)`) with `nil`, which evaluates to `false` because the type component is non-nil.

## Language Notes

- **Language Features**: Written using idiomatic Go paradigms.
- **Go Interfaces**: Implicit structural typing (duck typing) verified at compile time, with dynamic dispatch at runtime.

## Related Concepts

- Data Structures & Algorithms in Go
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `notification.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `notification.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`notification.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface/others_hizibiji_code/notification.go)
- [`payment.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface/others_hizibiji_code/payment.go)
- [`temp.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Intermediate-01/Interface/others_hizibiji_code/temp.go)
