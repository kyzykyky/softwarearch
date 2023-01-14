# Layering

## Book service CRUD implementation in Go

This project uses:

- Sqlite as a database
- NATS as a MQ
- Fiber as a web framework
- Zap as a logger

### Project structure

```bash
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── cfg
│   └── logger.yaml
├── cmd
│   └── booksvc
│       └── main.go
├── internal
│   ├── controller
│   │   └── server
│   │       ├── fiber
│   │       │   ├── docs
│   │       │   │   ├── docs.go
│   │       │   │   └── swagger.json
│   │       │   ├── errors.go
│   │       │   ├── handlers.go
│   │       │   ├── lifecycle.go
│   │       │   ├── routes.go
│   │       │   ├── server.go
│   │       │   └── utils.go
│   │       └── server.go
│   ├── data
│   │   ├── bookrepository
│   │   │   ├── bookrepository.go
│   │   │   └── gorm
│   │   │       ├── entity.go
│   │   │       ├── gorm_sqlite_bookrepo_test.go
│   │   │       ├── init.go
│   │   │       └── repository.go
│   │   ├── config
│   │   │   ├── dbconnect.go
│   │   │   └── gorm
│   │   │       └── sqlite
│   │   │           ├── connect.go
│   │   │           └── errors.go
│   │   └── errors
│   │       └── errors.go
│   ├── domain
│   │   └── domain.go
│   ├── integration
│   │   ├── logger
│   │   │   └── logger.go
│   │   └── mq
│   │       ├── mq.go
│   │       └── nats
│   └── service
│       ├── book.go
│       └── service.go
└── pkg
    └── domain
        ├── book.go
        └── service.go
```

### Service Interface

```go
type BookService interface {
GetBook(ctx context.Context, id int) (Book, error)
GetBooks(ctx context.Context, count, offset int) ([]Book, error)
CreateBook(ctx context.Context, book Book) (Book, error)
UpdateBook(ctx context.Context, book Book) (Book, error)
DeleteBook(ctx context.Context, id int) error
}
```

## Pre-requisites

### Installation

```bash
go mod download
```

### Run

```bash
make run
```

### Build

```bash
make build
```

### Swagger documentation for API

After launching the service, you can access the documentation at:

```bash
http://localhost:3000/docs/index.html
```

OR  
Swagger json file is available at:  
internal/controller/server/fiber/docs/swagger.json

### Dev requirements

- Development reload tool: [reflex](https://github.com/cespare/reflex)
- Swagger documentation generator: [swaggo](https://github.com/swaggo/swag)
- Linter: [golangci-lint](https://github.com/golangci/golangci-lint)

```bash
go install github.com/cespare/reflex@latest
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
```

### Run in dev mode

```bash
make watch
```
