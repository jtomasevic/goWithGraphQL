package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/evax/app"
	"github.com/evax/graph"
	"github.com/evax/graph/generated"
	dbMigrate "github.com/evax/app/data_sources/evax/migrations"
	"github.com/go-chi/chi"
	auth "github.com/evax/app/services/auth"
)

const defaultPort = "3001"



func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	app.InitApp()

	router := chi.NewRouter()
	router.Use(auth.Middleware())

	dbMigrate.EvaxDbAutoMigrate()


	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
