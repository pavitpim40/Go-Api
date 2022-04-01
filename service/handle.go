package service

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handle struct {
	service Servicer
}

func NewHandle(service Servicer) *Handle {
	return &Handle{service: service}
}

func (h *Handle) CallApiLogin(c echo.Context) error {
	// var res Response
	var req Request 
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	if len(req.Password) < 8 {
		c.JSON(http.StatusBadRequest, "Password must be at least 8 characters")
	}
	data := c.Request().Header.Get("app-id")
	fmt.Println((data))
	fmt.Printf("req: %+v\n", req)
	response,_ := h.service.Login(c,req)
	return c.JSON(http.StatusOK, response)
}