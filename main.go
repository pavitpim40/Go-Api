package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Token string `json:"token"`
	Status string `json:"status"`
}
func main(){
	e := echo.New()
	e.GET("/",func(c echo.Context) error{
		return c.String(http.StatusOK,"Hello, World!")
	})
	e.POST("/login", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Response{})
	})
	e.Logger.Fatal(e.Start(":1323"))
}