package cmd

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"server.notion.a1007.wuchuheng.com/m/v2/graph"
	"time"
)

func Start() {
	app := &cli.App{
		Name:  "server",
		Usage: "Start service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "port",
				Value:   "8080",
				Usage:   "The server listening port number",
				Aliases: []string{"p"},
			},
		},
		Action: func(cCtx *cli.Context) error {
			startServer(cCtx)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func startServer(cCtx *cli.Context) {
	port := cCtx.String("port")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
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
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
