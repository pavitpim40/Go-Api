package main

import (
	"api/service"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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
	
	fmt.Println("Go MySQL Tutorial")
    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/csi_test")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    } else {
		fmt.Println("Connect success")
	}

    // defer the close till after the main function has finished
    // executing
    defer db.Close()


	services := service.NewHandle(service.NewService(database,redis))
	e.POST("/",services.CallApiLogin)
	e.Logger.Fatal(e.Start(":1323"))
}

