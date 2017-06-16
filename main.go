package main

import (
	"net/http"

	"hyangit/echohello/hello"
	"hyangit/echohello/singleton"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger(), middleware.Recover(), middleware.Gzip(), middleware.CORS())
	e.Validator = singleton.GetValidator()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	e.GET("/hello/:id", hello.Hello)
	e.GET("/index", hello.Index)

	e.Server.Addr = ":8888"
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}
