package repository

import (
	"category/domain/model"
)

// ICategoryRepository 操作用户数据库接口
type ICategoryRepository interface {

	InitTable() error

	CreateCategory(*model.Category) (int64, error)

	DeleteCategory(int64) error

	UpdateCategory(*model.Category) error

	FindAll() ([]model.Category,error)

	FindCategoryByName(string) (*model.Category, error)

	FindCategoryById(int64) (*model.Category, error)

	FindCategoryByParent(int32) ([]model.Category,error)

	FindCategoryByLevel(int32) ([]model.Category,error)

}