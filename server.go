package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ericgrandt/gqlgen-example/database"
	"github.com/ericgrandt/gqlgen-example/graph/generated"
	"github.com/ericgrandt/gqlgen-example/graph/resolver"
	"github.com/go-chi/chi/v5"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := database.GetOrCreateDatabase("./gqlgenexample.db")
	defer db.Close()

	router := chi.NewRouter()

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: resolver.NewResolver(db)}))
	srv.AddTransport(transport.POST{})
	router.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
