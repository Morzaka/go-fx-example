package server

import (
	"context"
	"net/http"

	"go.uber.org/fx"
)

func SModule() fx.Option {
	return fx.Invoke(
		RegisterHandlers,
		InitServer,
	)
}

func RegisterHandlers(hanler http.Handler) {
	http.Handle("/", hanler)
}

func InitServer(lifecycle fx.Lifecycle) {
	server := &http.Server{
		Addr: ":8080",
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Close()
		},
	})
}
