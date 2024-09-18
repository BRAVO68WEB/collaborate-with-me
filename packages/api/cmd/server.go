package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/db"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/helpers"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/repository"
	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/graph"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn := db.ConnectMongo()
	awsSession := helpers.ConnectS3()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Repositories: repository.Init(
			conn,
			awsSession,
		),
	}}))

	srv.AddTransport(&transport.Websocket{})

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", authMiddleware(srv))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("authorization")

		if header == "" {
			next.ServeHTTP(w, r)
			return
		}

		isValid, claims := helpers.VerifyJWT(header)

		println(isValid)

		if !isValid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
