package main

import (
	"github.com/crossplatform-mapdb/graphql-api/graphql/generated"
	customMiddleware "github.com/crossplatform-mapdb/graphql-api/middleware"
	"github.com/crossplatform-mapdb/graphql-api/postgres"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v9"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/crossplatform-mapdb/graphql-api/graphql"
)

const defaultPort = "8080"

func main() {
	DB := postgres.New(&pg.Options{
		Addr:     os.Getenv("ADDR"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASS"),
		Database: os.Getenv("DB"),
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := defaultPort

	userRepo := postgres.UsersRepo{DB: DB}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		//AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(userRepo))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{
		PlacesRepo: postgres.PlacesRepo{DB: DB},
		UsersRepo:  userRepo,
	}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
