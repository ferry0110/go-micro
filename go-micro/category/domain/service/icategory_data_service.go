package service

import "category/domain/model"

type ICategoryDataService interface {
	AddCategory(*model.Category) (int64, error)

	DeleteCategory(int64) error

	UpdateCategory(category *model.Category) (err error)

	FindByName(categoryname string) (*model.Category, error)

	FindById(id int64) (*model.Category, error)

	FindAllCategory() ([]model.Category, error)

	FindByLevel(levelId int32) ([]model.Category, error)

	FindByParent(parentId int32) ([]model.Category, error)

}
