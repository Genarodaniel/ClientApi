package main

import "clientApi/config/server"

// @title Swagger Hank API
// @version 1.0
// @description API responsible for managing Q2Bank

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	s := server.Init()
	s.Run()
}
