server:
	go run cmd/server/server.go

generate:
	go run github.com/99designs/gqlgen generate
