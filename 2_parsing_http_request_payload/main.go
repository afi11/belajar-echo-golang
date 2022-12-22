package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func main() {
	r := echo.New()
	// routes here

	// BISA POST atau GET
	r.Any("/user", func(ctx echo.Context) (err error) {
		u := new(User)
		if err = ctx.Bind(u); err != nil {
			return
		}

		return ctx.JSON(http.StatusOK, u)
	})

	fmt.Println("server started at :9000")
	r.Start(":9000")
}
