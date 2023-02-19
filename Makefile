server:
	go run cmd/server/server.go

generate:
	go run github.com/99designs/gqlgen generate

db:
	docker run --rm -it  -v $(pwd):/data keinos/sqlite3 sqlite3 data/data.db
