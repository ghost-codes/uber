package main

import (
	"database/sql"
	"log"

	"firebase.google.com/go/auth"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	db "github.com/ghost-codes/uber/db/sqlc"
	"github.com/ghost-codes/uber/graph"
	directives "github.com/ghost-codes/uber/graph/directives"
	resolver "github.com/ghost-codes/uber/graph/resolver"
    "github.com/ghost-codes/uber/middleware"
	"github.com/ghost-codes/uber/util"
	"github.com/gin-gonic/gin"
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

    router:= gin.Default()
    router.Use(middleware.AuthMiddleware(*store,auth))
    
    router.POST("/graphql",graphqlHandler(store,config,auth))
    router.GET("/",playgroundHandler())

	log.Printf("connect to http://localhost:8080/ for GraphQL playground")
    router.Run()
}

func graphqlHandler(store *db.Store,config util.Config, auth *auth.Client) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
    
    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{
        Store:  store,
        Config: config,
        FirebaseAuth: auth,
    },
    Directives: graph.DirectiveRoot{
        Auth: directives.UserAuthDirective,
    },
}))


	return func(c *gin.Context) {
	    srv.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
