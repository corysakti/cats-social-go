package impl

import (
	"context"
	"database/sql"
	"github.com/corysakti/cats-social-go/helper"
	"github.com/corysakti/cats-social-go/model/entity"
	"github.com/corysakti/cats-social-go/model/web/request"
	"github.com/corysakti/cats-social-go/model/web/response"
	"github.com/corysakti/cats-social-go/repository"
	"github.com/corysakti/cats-social-go/service"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validator          *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validator *validator.Validate) service.CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validator:          validator}
}

func (service CategoryServiceImpl) Create(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := entity.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}

func (service CategoryServiceImpl) Update(ctx context.Context, request request.CategoryUpdateRequest) response.CategoryResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := entity.Category{
		Id:   request.Id,
		Name: request.Name,
	}

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)
	service.CategoryRepository.Delete(ctx, tx, category)

}

func (service CategoryServiceImpl) FindById(ctx context.Context, categoryId int) response.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) FindAll(ctx context.Context) []response.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryListResponse(categories)
}
