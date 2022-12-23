package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	configAppName := os.Getenv("APP_NAME")
	if configAppName == "" {
		e.Logger.Fatal("APP_NAME config is required")
	}

	confServerPort := os.Getenv("SERVER_PORT")
	if confServerPort == "" {
		e.Logger.Fatal("SERVER_PORT config is required")
	}

	e.GET("/index", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, true)
	})

	server := new(http.Server)
	server.Addr = ":" + confServerPort

	if confServerReadTimeout := os.Getenv("SERVER_READ_TIMEOUT_IN_MINUTE"); confServerReadTimeout != "" {
		duration, _ := strconv.Atoi(confServerReadTimeout)
		server.ReadTimeout = time.Duration(duration) * time.Minute
	}

	if confServerWriteTimeout := os.Getenv("SERVER_WRITE_TIMEOUT_IN_MINUTE"); confServerWriteTimeout != "" {
		duration, _ := strconv.Atoi(confServerWriteTimeout)
		server.WriteTimeout = time.Duration(duration) * time.Minute
	}

	e.Logger.Print("Starting", configAppName)
	e.Logger.Fatal(e.StartServer(server))
}
