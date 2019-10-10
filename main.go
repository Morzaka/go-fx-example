package main

import (
	"github.com/Morzaka/go-fx-example/handlers"
	"github.com/Morzaka/go-fx-example/logger"
	"github.com/Morzaka/go-fx-example/server"

	"go.uber.org/fx"
)

func mainfx() fx.Option {
	return fx.Options(
		server.SModule(),
		logger.LModule(),
		handlers.HModule(),
	)
}

func main() {
	app := fx.New(
		mainfx(),
	)
	app.Run()
}
