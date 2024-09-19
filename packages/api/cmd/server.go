package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/db"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/graph"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/helpers"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/middleware"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
)

const defaultPort = "4040"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var g errgroup.Group

	routerV1 := Init()

	addr := fmt.Sprintf("0.0.0.0:%s", defaultPort)

	println("Server is running on http://", addr)

	s := &http.Server{
		Addr:         addr,
		Handler:      routerV1,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     log.New(os.Stderr, "", 0),
	}

	g.Go(func() error {
		return s.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func graphqlHandler(introspectionEnabled bool) gin.HandlerFunc {
	RedisAddress := os.Getenv("REDIS_ADDRESS")

	cacheAPQ, err := helpers.NewAPQCache(RedisAddress, 24*time.Hour)
	if err != nil {
		log.Fatalf("cannot create APQ redis cache: %v", err)
	}

	cacheSQC, err := helpers.NewSQCCache(RedisAddress, 24*time.Hour)
	if err != nil {
		log.Fatalf("cannot create APQ redis cache: %v", err)
	}

	conn := db.ConnectMongo()
	awsSession := helpers.ConnectS3()

	h := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					Repositories: repository.Init(
						conn,
						awsSession,
					),
				},
			},
		),
	)

	h.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.MultipartForm{})

	h.SetQueryCache(cacheSQC)

	h.Use(extension.AutomaticPersistedQuery{
		Cache: cacheAPQ,
	})

	if introspectionEnabled {
		h.Use(extension.Introspection{})
	}

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func Init() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	// setup cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"Accept", "Accept-CH", "Accept-Charset", "Accept-Datetime", "Accept-Encoding", "Accept-Ext", "Accept-Features", "Accept-Language", "Accept-Params", "Accept-Ranges", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Origin", "Access-Control-Expose-Headers", "Access-Control-Max-Age", "Access-Control-Request-Headers", "Access-Control-Request-Method", "Authorization", "Content-Type"}
	corsConfig.AllowAllOrigins = true

	router.Use(cors.New(corsConfig))

	introspectionEnabled := true

	healthHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK!",
		})
	}
	router.Use(middleware.GinContextToContext())

	router.GET("/playground", playgroundHandler())

	router.GET("/health", healthHandler)

	router.POST("/query", middleware.NewJWT().Auth(context.Background()), graphqlHandler(introspectionEnabled))

	return router
}
