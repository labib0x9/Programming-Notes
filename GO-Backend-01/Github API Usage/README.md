<!-- AUTO-GENERATED FROM SOURCE CODE -->
# GO-Backend-01 / Github API Usage

## Overview

This directory explores **GO-Backend-01 / Github API Usage** in **Go**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of GO-Backend-01 / Github API Usage in Go.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`add_collaborator.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/add_collaborator.go) | Go | app must be installed on the organization github/organizations/<org-name>/settings |
| [`clone_using_app_key.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/clone_using_app_key.go) | Go | app must be installed on the organization github/organizations/<org-name>/settings |
| [`create_a_repo.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/create_a_repo.go) | Go | app must be installed on the organization github/organizations/<org-name>/settings |
| [`create_repo_from_template.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/create_repo_from_template.go) | Go | app must be installed on the organization template repository te `Template repository` select krte hobe |
| [`delete_repo.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/delete_repo.go) | Go | app must be installed on the organization github/organizations/<org-name>/settings |

## Concepts

### 1. Add Collaborator (`add_collaborator.go`)

File: [`add_collaborator.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/add_collaborator.go)  
app must be installed on the organization github/organizations/<org-name>/settings.

Relevant code excerpt:

```go
func NewClient(org string, appId, instalationId int64) *github.Client {
	transport, err := gh.NewKeyFromFile(
		http.DefaultTransport,
		appId,
		instalationId,
		"hikma-test-service.2026-08-12.private-key.pem",
	)
	if err != nil {
		panic(err)
	}

	return github.NewClient(&http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	})
}

func CreateRepository(ctx context.Context, client *github.Client, org, template, repo string) error {
	trepo, resp, err := client.Repositories.CreateFromTemplate(
		ctx,
		org,
		template, &github.TemplateRepoRequest{
```

**Explanation**:
- Implements function(s): `NewClient()`, `CreateRepository()`, `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Clone Using App Key (`clone_using_app_key.go`)

File: [`clone_using_app_key.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/clone_using_app_key.go)  
app must be installed on the organization github/organizations/<org-name>/settings.

Relevant code excerpt:

```go
func NewToken(org string, appId, instalationId int64) string {
	transport, err := gh.NewKeyFromFile(
		http.DefaultTransport,
		appId,
		instalationId,
		"hikma-test-service.2026-08-12.private-key.pem",
	)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	token, err := transport.Token(ctx)
	if err != nil {
		panic(err)
	}
	return token
}

func main() {
```

**Explanation**:
- Implements function(s): `NewToken()`, `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Create A Repo (`create_a_repo.go`)

File: [`create_a_repo.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/create_a_repo.go)  
app must be installed on the organization github/organizations/<org-name>/settings.

Relevant code excerpt:

```go
func main() {

	// from github app, replace with real id
	appId := 0773456
	instalationId := 478367040
	org := "ZERO9xz"
	repoName := "hikma-from-github-api-test"

	transport, err := gh.NewKeyFromFile(
		http.DefaultTransport,
		int64(appId),
		int64(instalationId),
		"hikma-test-service.2026-08-12.private-key.pem",
	)
	if err != nil {
		panic(err)
	}

	client := github.NewClient(&http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	})
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Create Repo From Template (`create_repo_from_template.go`)

File: [`create_repo_from_template.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/create_repo_from_template.go)  
app must be installed on the organization template repository te `Template repository` select krte hobe.

Relevant code excerpt:

```go
func main() {

	// from github app
	appId := 0773456
	instalationId := 478367040
	org := "ZERO9xz"
	repoName := "test-02"

	transport, err := gh.NewKeyFromFile(
		http.DefaultTransport,
		int64(appId),
		int64(instalationId),
		"hikma-test-service.2026-08-12.private-key.pem",
	)
	if err != nil {
		panic(err)
	}

	client := github.NewClient(&http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	})
```

**Explanation**:
- Implements function(s): `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.

### 5. Delete Repo (`delete_repo.go`)

File: [`delete_repo.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/delete_repo.go)  
app must be installed on the organization github/organizations/<org-name>/settings.

Relevant code excerpt:

```go
func NewClient(org string, appId, instalationId int64) *github.Client {
	transport, err := gh.NewKeyFromFile(
		http.DefaultTransport,
		appId,
		instalationId,
		"hikma-test-service.2026-08-12.private-key.pem",
	)
	if err != nil {
		panic(err)
	}

	return github.NewClient(&http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	})
}

func CreateRepository(ctx context.Context, client *github.Client, org, template, repo string) error {
	trepo, resp, err := client.Repositories.CreateFromTemplate(
		ctx,
		org,
		template, &github.TemplateRepoRequest{
```

**Explanation**:
- Implements function(s): `NewClient()`, `CreateRepository()`, `AddCollaborator()`, `main()`.
- Defines C routine(s): `main()`.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`add_collaborator.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/add_collaborator.go) provides or tests `add_collaborator`.
2. [`clone_using_app_key.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/clone_using_app_key.go) provides or tests `clone_using_app_key`.
3. [`create_a_repo.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/create_a_repo.go) provides or tests `create_a_repo`.
4. [`create_repo_from_template.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/create_repo_from_template.go) provides or tests `create_repo_from_template`.

```text
Caller / Test Runner
  ├──> [add_collaborator.go] (app must be installed on the organi...)
  ├──> [clone_using_app_key.go] (app must be installed on the organi...)
  ├──> [create_a_repo.go] (app must be installed on the organi...)
  ├──> [create_repo_from_template.go] (app must be installed on the organi...)
  └──> Execution Completion / Assertion
```

## Compilation & Execution

```bash
# Run with Go
go run .
# Or run specific file
go run add_collaborator.go
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

1. What is the primary role of the functions/routines demonstrated in `add_collaborator.go`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `add_collaborator.go` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`add_collaborator.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/add_collaborator.go)
- [`clone_using_app_key.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/clone_using_app_key.go)
- [`create_a_repo.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/create_a_repo.go)
- [`create_repo_from_template.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/create_repo_from_template.go)
- [`delete_repo.go`](file:///Users/labib0x9/Desktop/Programming-Notes/GO-Backend-01/Github API Usage/delete_repo.go)
