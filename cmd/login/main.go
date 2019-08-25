package main

import (
	"stub/api"
	"stub/route"
	"stub/service"
)

func main() {
	loginService := &service.LoginService{}
	loginAPI := &api.LoginAPI{
		LoginService: loginService,
	}
	server := route.NewServer(loginAPI)
	server.Run(":8888")
}
