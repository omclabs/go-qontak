package middlewares

import (
	"net/http"
)

type GeneralMiddleware struct {
	Handler http.Handler
}

func NewGeneralMiddleware(handler http.Handler) *GeneralMiddleware {
	return &GeneralMiddleware{Handler: handler}
}

func (middleware *GeneralMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")

	middleware.Handler.ServeHTTP(writer, request)
}
