package middlewares

// import (
// 	"net/http"
// 	"omclabs/go-qontak/app/helpers"
// 	"omclabs/go-qontak/app/models/web"
// )

// type AuthMiddleware struct {
// 	Handler http.Handler
// }

// func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
// 	return &AuthMiddleware{Handler: handler}
// }

// func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	if "RAHASIA" == request.Header.Get("X-API-Key") {
// 		// ok
// 		middleware.Handler.ServeHTTP(writer, request)
// 	} else {
// 		// error
// 		writer.Header().Set("Content-Type", "application/json")
// 		writer.WriteHeader(http.StatusUnauthorized)

// 		webResponse := web.ApiResponse{
// 			Code:   http.StatusUnauthorized,
// 			Status: "UNAUTHORIZED",
// 		}

// 		helpers.WriteToResponseBody(writer, webResponse)
// 	}
// }
