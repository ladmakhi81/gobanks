build:
	@go build -o ./bin/output

run: build
	@./bin/output

test:
	@go test -v ./...