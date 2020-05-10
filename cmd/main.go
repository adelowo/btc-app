package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/adelowo/queryapp/graph"
	"github.com/adelowo/queryapp/graph/generated"
	"github.com/friendsofgo/graphiql"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		log.Fatalf("Could not set up graphiql...%v", err)
	}

	http.Handle("/graphiql", graphiqlHandler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
