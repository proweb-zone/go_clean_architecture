package main

import (
	"log/slog"

	"clean/architector/internal/app"
	"clean/architector/internal/domain/repository"
	"clean/architector/internal/domain/usecase"
)

type IproductUseCase interface {
	Run(productId int)
}

func main() {

	cfg := app.MustLoad()           // load config file local.yaml
	log := app.SetupLogger(cfg.Env) // setup custom Logger

	log.Info(
		"starting url-shortener",
		slog.String("env", cfg.Env),
		slog.String("version", "123"),
	)

	log.Debug("debug messages are enabled")

	var c *app.Server = app.NewServer(cfg)
	c.StartServer()

	return

	var productId int = 2
	var iProductUseCase IproductUseCase = usecase.NewViewProductUseCase(repository.NewProductRepo(productId))
	iProductUseCase.Run(productId)
}
