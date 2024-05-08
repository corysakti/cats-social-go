package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AuthenticationController interface {
	Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
