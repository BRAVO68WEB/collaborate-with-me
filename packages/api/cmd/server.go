package main

import (
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/db"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/graph"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connecting to the database
	conn := db.ConnectMongo()

	// Setting up graphql server endpoint "/"
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Conn: &conn,
	}}))

	// Setting up the playground
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	// Fetching the port number from the env file and connecting the server to the port
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
