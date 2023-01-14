build: swagdocs
	go build -o build/booksvc cmd/booksvc/main.go

run:
	go run cmd/booksvc/main.go

watch:
	reflex -s -r '\.go$$' make run

lint:
	golangci-lint run --timeout 5m0s

swagdocs:
	swag fmt && swag init --generalInfo server.go --dir internal/controller/server/fiber,pkg/domain --output internal/controller/server/fiber/docs --outputTypes go,json
