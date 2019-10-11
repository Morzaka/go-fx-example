package handlers

import (
	"io"
	"net/http"

	"github.com/Morzaka/go-fx-example/logger"

	"go.uber.org/fx"
)

func HModule() fx.Option {
	return fx.Provide(NewHandler)
}

func NewHandler(logger logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Println("Handler called")
		_, _ = io.WriteString(w, "Hello, world!\n")
	})
}
