package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"firebase.google.com/go/auth"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	db "github.com/ghost-codes/uber/db/sqlc"
	"github.com/ghost-codes/uber/graph"
	"github.com/ghost-codes/uber/graph/directives"
	resolver "github.com/ghost-codes/uber/graph/resolver"
	"github.com/ghost-codes/uber/token"
	"github.com/ghost-codes/uber/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	//xyz
	err, config := util.LoadConfig(".")
	if err != nil {
		log.Fatal("unable to load config: ", err)
	}

	auth := util.SetupFirebaseClient(config.FirebaseConfigPath)

	// init store
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("error unable to connect to database: ", err)
	}

	store := db.NewStore(conn)
	maker, err := token.NewJWTMaker(config.Secret)
	if err != nil {
		log.Fatal("error unable to initialize token maker: ", err)
	}

	// writeConfig := kafka.WriterConfig{
	//     Brokers: []string{config.KafkaHost},
	//     Topic: "driver-location",
	//     Balancer: &kafka.LeastBytes{},
	// }

	// writer:=kafka.NewWriter(writeConfig)
	// defer writer.Close()

	// go func(){
	//     long:= 0.0;
	//     for{
	//         long=long+0.00001
	//         location:=model.CarLocation{
	//             Location: &model.Location{
	//                 Lat:0.0,
	//                 Long: long,
	//             },
	//             CarType: model.CarTypeLuxury,
	//             Driver: &db.Driver{
	//                 ID: int64(12),
	//             },
	//         }
	//         g,err:=json.Marshal(location)

	//         if err!=nil{
	//             log.Println(err)
	//             break;
	//         }

	//         kafkaMessage:=kafka.Message{
	//             Value: g,
	//         }
	//         cerr:=writer.WriteMessages(context.Background(),kafkaMessage)

	//         if err != nil {
	//             fmt.Println("Error writing message:", cerr)
	//         } else {
	//             //fmt.Println("Wrote message:", kafkaMessage)
	//         }

	//     // Wait for one second before sending the next location update
	//     time.Sleep(time.Second)
	//     }

	// }()

	router := gin.Default()
	// router.Use(middleware.AuthMiddleware(*store, auth))

	router.Use(util.GinContextToContextMiddleware())
	router.Any("/graphql", graphqlHandler(store, config, auth, maker))
	router.GET("/", playgroundHandler())

	log.Printf("connect to http://localhost:8080/ for GraphQL playground")
	router.Run()
}

func graphqlHandler(store *db.Store, config util.Config, auth *auth.Client, maker token.Maker) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{
		Store:        store,
		Config:       config,
		FirebaseAuth: auth,
		Maker:        maker,
	},
		Directives: graph.DirectiveRoot{
			Auth:         directives.UserAuthDirective,
			Authenticate: directives.AuthenticateDriverfunc(maker),
		},
	}))

	srv.AddTransport(transport.SSE{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.Use(extension.Introspection{})

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
