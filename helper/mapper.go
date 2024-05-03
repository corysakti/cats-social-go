package helper

import (
	"github.com/corysakti/cats-social-go/model/entity"
	"github.com/corysakti/cats-social-go/model/web/response"
)

func ToCategoryResponse(category entity.Category) response.CategoryResponse {
	return response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryListResponse(categories []entity.Category) []response.CategoryResponse {
	var categoryResponses []response.CategoryResponse

	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}
