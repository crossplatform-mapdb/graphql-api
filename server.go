package main

import (
	"github.com/go-pg/pg/v9"
	"github.com/zackartz/go-graphql-api/postgres"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/zackartz/go-graphql-api/graphql"
	"github.com/zackartz/go-graphql-api/graphql/generated"
)

const defaultPort = "8080"

func main() {
	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "pog_dev",
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{
		PlacesRepo: postgres.PlacesRepo{DB: DB},
		UsersRepo:  postgres.UsersRepo{DB: DB},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
