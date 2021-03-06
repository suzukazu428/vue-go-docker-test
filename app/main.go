package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// echo導入
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8082"))
	// e.Startの中はdocker-composeのgoコンテナで設定したportsを指定
}
