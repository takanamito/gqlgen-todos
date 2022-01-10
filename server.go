package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"

	"github.com/takanamito/gqlgen-todos/ent"
	"github.com/takanamito/gqlgen-todos/graph"
	"github.com/takanamito/gqlgen-todos/graph/generated"
)

const defaultPort = "8080"

func main() {
	// マイグレーション
	client, err := ent.Open("postgres", "host=localhost port=5432 user=admin dbname=todos password=admin sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if err := client.Debug().Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// サーバーの初期化
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
