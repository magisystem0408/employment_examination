package main

import (
	"context"
	"graphql-server/db"
	"graphql-server/graph"
	"graphql-server/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	store := db.NewKeyValueStore()
	rslv := graph.NewResolver(store)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: rslv}))

	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		oc := graphql.GetOperationContext(ctx)
		switch oc.OperationName {
		case "":
			log.Println("no operation named ''")
		case "IntrospectionQuery":
		default:
			log.Printf("operationName: %s", oc.OperationName)
		}
		return next(ctx)
	})

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))

	mux.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "header", r.Header)

		srv.ServeHTTP(w, r.WithContext(ctx))
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
