package main

import (
	"api/service"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Token string `json:"token"`
	Status string `json:"status"`
}
func main(){
	// e := echo.New()
	// e.GET("/",func(c echo.Context) error{
	// 	return c.String(http.StatusOK,"Hello, World!")
	// })
	// e.POST("/login", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, Response{
	// 		Token: "123456789",
	// 		Status : "success",
	// 	})
	// })

	var (
		e = echo.New()
		redis = ""
		database = ""
	)
	
	services := service.NewHandle(service.NewService(database,redis))
	e.POST("/",services.CallApiLogin)
	e.Logger.Fatal(e.Start(":1323"))
}

