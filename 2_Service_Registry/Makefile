run:
	go run cmd/main.go

watch:
	reflex -s -r '\.go$$' make run

lint:
	golangci-lint run --timeout 5m0s