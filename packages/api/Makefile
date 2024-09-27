build:
	@echo "Building GraphQL API ..."
	@go build -o bin/gographql cmd/server.go

dev:
	@echo "Running GraphQL API ... [watch-mode]"
	air -c .air.toml

gen:
	@echo "Generating GraphQL API ..."
	@go run github.com/99designs/gqlgen generate