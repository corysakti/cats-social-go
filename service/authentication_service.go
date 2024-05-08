package service

import (
	"context"
	"github.com/corysakti/cats-social-go/model/web/request"
	"github.com/corysakti/cats-social-go/model/web/response"
)

type AuthenticationService interface {
	Login(ctx context.Context, request request.LoginRequest) response.AuthenticationResponse
	Register(ctx context.Context, request request.RegisterRequest) response.AuthenticationResponse
}
