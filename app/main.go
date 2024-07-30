package main

import (
	"clean/architector/internal/app"
)

type IproductUseCase interface {
	Run(productId int)
}

func main() {
	var config *app.Config = app.InitConfig()
	var c *app.Server = app.NewServer(config)
	c.StartServer()
}
