package main

import (
	"net/http"

	bug "entgo.io/bug"
	"entgo.io/bug/db"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	client, err := db.GetClient()
	logger := e.Logger

	if err != nil {
		logger.Info("Database could not initialize")
	}

	//defer client.Close()

	engine := handler.NewDefaultServer(bug.NewSchema(client))

	http.Handle("/",
		playground.Handler("User", "/query"),
	)
	http.Handle("/query", engine)
	logger.Info("listening on :8081")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		logger.Fatal("http server terminated", err)
	}
}
