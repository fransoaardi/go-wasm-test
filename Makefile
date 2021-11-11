.PHONY: all
all: build run

.PHONY: build
build:
	GOOS=js GOARCH=wasm go build -o static/main.wasm cmd/wasm/main.go
	go build -o backend cmd/backend/main.go

.PHONY: run
run:
	./backend