package exception

import (
	"net/http"
	"restfullapi/golang/helper"
	"restfullapi/golang/model/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	internalSeverError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {

	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true

	} else {
		return false
	}
}

func internalSeverError(writer http.ResponseWriter, _ *http.Request, err interface{}) {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:    http.StatusInternalServerError,
		Status:  "Internal Server Error",
		Message: err.(error).Error(),
	}

	helper.WriteToResponseBody(writer, webResponse)
}
