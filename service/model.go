package service

type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
 }

 
 type Response struct {
	 Token string `json:"token"`
 }