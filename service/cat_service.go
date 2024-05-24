package service

import (
	"context"
	"github.com/corysakti/cats-social-go/model/web/request"
	"github.com/corysakti/cats-social-go/model/web/response"
)

type CatService interface {
	Create(ctx context.Context, request request.CatRequest) response.CatResponse
	Update(ctx context.Context, request request.CatRequest, catId int) response.CatResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) response.CatResponse
	FindAll(ctx context.Context) []response.CatResponse
}
