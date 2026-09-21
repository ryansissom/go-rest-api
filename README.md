## Go REST API

A small REST API for creating and managing news items. It uses an in-memory
store, so data is reset whenever the server stops.

Requires Go 1.27 or newer.

### Project Structure

The API uses Go's standard `net/http` package with method-aware routing,
request validation, JSON encoding, and UUID-based resource IDs. Handlers depend
on a small store interface, while the current implementation uses a
mutex-protected in-memory store.

Structured logging is applied through HTTP middleware, and the project includes
table-driven tests for request validation and handler behavior.

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
