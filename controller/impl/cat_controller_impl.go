package impl

import (
	"github.com/corysakti/cats-social-go/controller"
	"github.com/corysakti/cats-social-go/service"
)

type CatController struct {
}

func NewCatController(categoryService service.CategoryService) controller.CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}
