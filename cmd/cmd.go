package cmd

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"server.notion.a1007.wuchuheng.com/m/v2/graph"
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

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
