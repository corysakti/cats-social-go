package impl

import (
	"github.com/corysakti/cats-social-go/controller"
	"github.com/corysakti/cats-social-go/helper"
	request2 "github.com/corysakti/cats-social-go/model/web/request"
	"github.com/corysakti/cats-social-go/model/web/response"
	"github.com/corysakti/cats-social-go/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AuthenticationControllerImpl struct {
	AuthenticationService service.AuthenticationService
}

func NewAuthenticationController(authenticationService service.AuthenticationService) controller.AuthenticationController {
	return &AuthenticationControllerImpl{
		AuthenticationService: authenticationService,
	}
}

func (controller AuthenticationControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	registerRequest := request2.RegisterRequest{}
	helper.ReadFromRequestBody(request, &registerRequest)

	dataResponse := controller.AuthenticationService.Register(request.Context(), registerRequest)
	responseTemplate := response.DataResponse{
		Message: "User registered successfully",
		Data:    dataResponse,
	}

	helper.WriteToResponseBodyAndStatusCode(writer, 201, responseTemplate)
}

func (controller AuthenticationControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := request2.LoginRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	dataResponse := controller.AuthenticationService.Login(request.Context(), loginRequest)
	responseTemplate := response.ResponseTemplate{
		Code:   200,
		Status: "OK",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, responseTemplate)
}
