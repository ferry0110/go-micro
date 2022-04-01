package service

import (
	"category/domain/model"
	"category/domain/repository"
)

type CategoryDataService struct {
	CategoryRepository repository.ICategoryRepository
}

// NewCategoryDataService 创建CategoryDataService
func NewCategoryDataService(categoryRepository repository.ICategoryRepository) ICategoryDataService {
	return &CategoryDataService{CategoryRepository: categoryRepository}
}

func (u *CategoryDataService) AddCategory(category *model.Category) (int64, error) {
	return  u.CategoryRepository.CreateCategory(category)
}

func (u *CategoryDataService) DeleteCategory(i int64) error {
	return u.CategoryRepository.DeleteCategory(i)
}

func (u *CategoryDataService) FindByName(categoryname string) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByName(categoryname)
}

func (u *CategoryDataService) UpdateCategory(category *model.Category) (err error) {
	return u.CategoryRepository.UpdateCategory(category)
}

func (u *CategoryDataService) FindById(id int64) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryById(id)
}

func (u *CategoryDataService) FindAllCategory() ([]model.Category, error) {
	return u.CategoryRepository.FindAll()
}

func (u *CategoryDataService) FindByLevel(levelId int32) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByLevel(levelId)
}

func (u *CategoryDataService) FindByParent(parentId int32) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByParent(parentId)
}
