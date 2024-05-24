package middleware

import (
	"net/http"
	"restfullapi/golang/helper"
	"restfullapi/golang/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "Unauthorized",
			Message: "Invalid X-API-Key",
		}
		helper.WriteToResponseBody(writer, webResponse)
	}
}
