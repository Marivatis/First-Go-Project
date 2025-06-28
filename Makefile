run:
	go run cmd/main.go

.PHONY: test test-cover

test:
	go test -v -count=1 ./internal/repository

test-cover:
	go test -v -cover -count=1 ./internal/repository