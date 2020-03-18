package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fanadol/graphql-go-example/internal/constants"
	"github.com/fanadol/graphql-go-example/internal/graph"
	"github.com/fanadol/graphql-go-example/internal/graph/generated"
	userpb "github.com/fanadol/graphql-go-example/internal/user/proto"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = constants.DefaultPort
	}

	conn, err := grpc.Dial(constants.GRPCPort, grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	usergrpcclient := userpb.NewUsersClient(conn)

	resolver := &graph.Resolver{
		GRPCUserClient: usergrpcclient,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(constants.DefaultPort, nil))
}
