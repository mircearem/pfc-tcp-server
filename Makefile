build:
	@go build -o bin/server.bin

run: build
	@./bin/server.bin
