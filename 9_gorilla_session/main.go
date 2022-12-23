package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/context"
	"github.com/labstack/echo"
)

const SESSION_ID = "id"

func newPostgresStore() *pgstore.PGStore {
	url := "postgres://fatakhul:1234@127.0.0.1:5432/fatakhul?sslmode=disable"
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store, err := pgstore.NewPGStore(url, authKey, encryptionKey)
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(0)
	}

	return store
}

func main() {
	store := newPostgresStore()

	e := echo.New()

	e.Use(echo.WrapMiddleware(context.ClearHandler))

	e.GET("/set", func(ctx echo.Context) error {
		session, _ := store.Get(ctx.Request(), SESSION_ID)
		session.Values["message1"] = "hello"
		session.Values["message2"] = "world"
		session.Save(ctx.Request(), ctx.Response())

		return ctx.Redirect(http.StatusTemporaryRedirect, "/get")
	})

	e.GET("/get", func(ctx echo.Context) error {
		session, _ := store.Get(ctx.Request(), SESSION_ID)

		if len(session.Values) == 0 {
			return ctx.String(http.StatusOK, "empty result")
		}

		return ctx.String(http.StatusOK, fmt.Sprintf(
			"%s %s",
			session.Values["message1"],
			session.Values["message2"],
		))
	})

	e.GET("/delete", func(ctx echo.Context) error {
		session, _ := store.Get(ctx.Request(), SESSION_ID)
		session.Options.MaxAge = -1
		session.Save(ctx.Request(), ctx.Response())

		return ctx.Redirect(http.StatusTemporaryRedirect, "/get")
	})

	e.Logger.Fatal(e.Start(":9000"))
}
