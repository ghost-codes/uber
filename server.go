package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	db "github.com/ghost-codes/uber/db/sqlc"
	"github.com/ghost-codes/uber/graph"
	resolver "github.com/ghost-codes/uber/graph/resolver"
	"github.com/ghost-codes/uber/util"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {

	err, config := util.LoadConfig(".")
	if err != nil {
		log.Fatal("unable to load config: ", err)
	}

    auth:= util.SetupFirebaseClient(config.FirebaseConfigPath)

	// init store
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("error unable to connect to database: ", err)
	}

	store := db.NewStore(conn)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{
		Store:  store,
		Config: config,
        FirebaseAuth: auth,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
