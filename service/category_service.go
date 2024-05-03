package service

import (
	"context"
	"github.com/corysakti/cats-social-go/model/web/request"
	"github.com/corysakti/cats-social-go/model/web/response"
)

type CategoryService interface {
	Create(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse
	Update(ctx context.Context, request request.CategoryUpdateRequest) response.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) response.CategoryResponse
	FindAll(ctx context.Context) []response.CategoryResponse
}
