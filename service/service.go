package service

import "github.com/labstack/echo/v4"

type Service struct {
	database string
	redis	string
}

type Servicer interface {
	 Login(c echo.Context, req Request) (*Response,error)
}
func NewService(database, redis string) *Service {
	return &Service{}
}

func (s *Service)Login(c echo.Context, req Request) (*Response,error) {
	return &Response{Token: req.Username + req.Password},nil
}