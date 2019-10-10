package logger

import (
	"go.uber.org/fx"

	"log"
	"os"
)

func LModule() fx.Option {
	return fx.Options(
			fx.Provide(NewLogger),
		)
}

type Logger interface {
	Println(v ...interface{})
}

func NewLogger() Logger {
	return log.New(os.Stdout, "[ACME]", 0)
}