package impl

import (
	"context"
	"database/sql"
	"github.com/corysakti/cats-social-go/exception"
	"github.com/corysakti/cats-social-go/helper"
	"github.com/corysakti/cats-social-go/model/entity"
	"github.com/corysakti/cats-social-go/model/web/request"
	"github.com/corysakti/cats-social-go/model/web/response"
	"github.com/corysakti/cats-social-go/repository"
	"github.com/corysakti/cats-social-go/service"
	"github.com/go-playground/validator/v10"
	"time"
)

type CatServiceImpl struct {
	CatRepository     repository.CatRepository
	UserCatRepository repository.UserCatRepository
	DB                *sql.DB
	Validator         *validator.Validate
}

func NewCatService(catRepository repository.CatRepository, userCatRepository repository.UserCatRepository, DB *sql.DB, validator *validator.Validate) service.CatService {
	return &CatServiceImpl{
		CatRepository:     catRepository,
		UserCatRepository: userCatRepository,
		DB:                DB,
		Validator:         validator}
}

func (service CatServiceImpl) Create(ctx context.Context, request request.CatRequest) response.CatResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cat := entity.Cat{
		Name:        request.Name,
		Race:        request.Race,
		Sex:         request.Sex,
		Description: request.Description,
		AgeInMonth:  int32(request.AgeInMonth),
		ImageUrls:   helper.ConvertArrToString(request.ImageUrls),
		CreatedAt:   time.Time{},
	}

	cat = service.CatRepository.Save(ctx, tx, cat)

	return response.CatResponse{
		Id:        int(cat.Id),
		CreatedAt: cat.CreatedAt,
	}
}

func (service CatServiceImpl) Update(ctx context.Context, request request.CatRequest, catId int) response.CatResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	existCat, err := service.CatRepository.FindById(ctx, tx, catId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	existMatched, err := service.UserCatRepository.FindByCatId(ctx, tx, catId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	existCat = service.CatRepository.Update(ctx, tx, existCat)

	return response.CatResponse{}
}

func (service CatServiceImpl) Delete(ctx context.Context, catId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CatRepository.FindById(ctx, tx, catId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	helper.PanicIfError(err)
	service.CatRepository.Delete(ctx, tx, category)

}

func (service CatServiceImpl) FindById(ctx context.Context, categoryId int) response.CatResponse {
	//TODO implement me
	panic("implement me")
}

func (service CatServiceImpl) FindAll(ctx context.Context) []response.CatResponse {
	//TODO implement me
	panic("implement me")
}
