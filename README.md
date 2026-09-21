## Go REST API

A small REST API for creating and managing news items. It uses an in-memory
store, so data is reset whenever the server stops.

I built this project while learning Go to strengthen my software engineering
skills as I work toward platform engineering. It applies core Go concepts in a
complete, testable service while relying primarily on the standard library.

Requires Go 1.27 or newer.

### Focus

The project intentionally uses Go's standard library where possible and focuses
on:

- HTTP routing with `net/http`
- Dependency inversion through small interfaces
- Request validation and error handling
- Structured logging with `log/slog`
- Concurrency-safe in-memory storage
- UUID-based resource identifiers
- Table-driven HTTP handler tests

### Architecture

| Layer | Responsibility |
| --- | --- |
| `cmd/main.go` | Creates the router, store, and middleware |
| Router | Connects HTTP methods and paths to handlers |
| Handlers | Decode requests, validate input, and return responses |
| `NewsStorer` | Keeps handlers independent of storage details |
| In-memory store | Stores news items behind a mutex |

### Run

```sh
go run ./cmd/main.go
```

The server listens on `http://localhost:8080`.

### Endpoints

| Method | Path | Description |
| --- | --- | --- |
| `POST` | `/news` | Create a news item |
| `GET` | `/news` | List all news items |
| `GET` | `/news/{news_id}` | Get one news item |
| `PUT` | `/news/{news_id}` | Update a news item |
| `DELETE` | `/news/{news_id}` | Delete a news item |

Request bodies use these fields: `author`, `title`, `summary`, `content`,
`created_at` (RFC3339), `source` (URL), and at least one `tags` value. For
updates, the ID in the URL determines which item is changed.

Example request:

```sh
curl -X POST http://localhost:8080/news \
	-H "Content-Type: application/json" \
	-d '{"author":"Ada","title":"A headline","summary":"A summary","content":"The article body","created_at":"2024-01-01T00:00:00Z","source":"https://example.com","tags":["go"]}'
```

### Checks

```sh
go test ./...
go vet ./...
```

The Makefile also provides:

```sh
make run
make fmt
make tidy
```
