package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
	_ "net/http"
)

func main() {
	fmt.Printf("Hello world")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world my old friend!!!!!")
	})

	e.Start(":8080")
}
