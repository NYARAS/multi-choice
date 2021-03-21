init:
	go run github.com/99designs/gqlgen init

generate:
	go run github.com/99designs/gqlgen && go run ./graph/models/model_tags/model_tags.go