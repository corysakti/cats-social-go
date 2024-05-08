package exception

import (
	"github.com/corysakti/cats-social-go/helper"
	"github.com/corysakti/cats-social-go/model/web/response"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	switch err := err.(type) {
	case NotFoundError:
		notFoundError(writer, request, err)
	case validator.ValidationErrors:
		validationErrors(writer, request, err)
	case AlreadyExistError:
		alreadyExistError(writer, request, err)
	case PasswordIsWrongError:
		passwordIsWrongError(writer, request, err)
	default:
		internalServerError(writer, request, err)
	}
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		responseTemplate := response.DataResponse{
			Message: "Validation Error!",
			Data:    exception.Error,
		}

		helper.WriteToResponseBody(writer, responseTemplate)
		return true
	} else {
		return false
	}
}

func passwordIsWrongError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(PasswordIsWrongError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		responseTemplate := response.DataResponse{
			Message: "password is wrong",
			Data:    exception.Error,
		}

		helper.WriteToResponseBody(writer, responseTemplate)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		responseTemplate := response.DataResponse{
			Message: "Data Not Error!",
			Data:    exception.Error,
		}

		helper.WriteToResponseBody(writer, responseTemplate)
		return true
	} else {
		return false
	}

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	responseTemplate := response.ResponseTemplate{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, responseTemplate)
}

func alreadyExistError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(AlreadyExistError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		responseTemplate := response.DataResponse{
			Message: "conflict if email exists",
			Data:    exception.Error,
		}

		helper.WriteToResponseBody(writer, responseTemplate)
		return true
	} else {
		return false
	}
}
