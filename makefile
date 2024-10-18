build:
	@go build -o bin/GOSE cmd/main.go


test:
	@go test -v ./...

run:build
	@./bin/GOSE