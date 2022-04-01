package controller

import (
	"category/common"
	"category/domain/model"
	"category/domain/service"
	"category/proto/category"
	"context"
)

type CategoryController struct {
	CategoryService service.ICategoryDataService
}

// CreateCategory 添加分类
func (c CategoryController) CreateCategory(ctx context.Context, request *category.CategoryRequest, response *category.CreateCategoryResponse) error {
	curCategory := &model.Category{}
	err := common.SwapTo(request, curCategory)
	if err != nil {
		return err
	}
	categoryId, err := c.CategoryService.AddCategory(curCategory)
	if err != nil {
		return err
	}
	response.Message = "分类添加成功"
	response.CategoryId = categoryId
	return nil
}

// UpdateCategory 更新分类
func (c CategoryController) UpdateCategory(ctx context.Context, request *category.CategoryRequest, response *category.UpdateCategoryResponse) error {
	curCategory := &model.Category{}
	err := common.SwapTo(request, curCategory)
	if err != nil {
		return err
	}
	err = c.CategoryService.UpdateCategory(curCategory)
	if err != nil {
		return err
	}
	response.Message = "分类更新成功"
	return nil
}

// DeleteCategory  删除分类
func (c CategoryController) DeleteCategory(ctx context.Context, request *category.DeleteCategoryRequest, response *category.DeleteCategoryResponse) error {
	curCategoryId := request.CategoryId
	curCategory, err := c.CategoryService.FindById(curCategoryId)
	if err != nil {
		return err
	}
	err = c.CategoryService.DeleteCategory(curCategoryId)
	if err != nil {
		return err
	}
	response.CategoryName = "删除:" + curCategory.CategoryName + "成功"
	return nil
}

// FindCategoryByName  查找分类ByName
func (c CategoryController) FindCategoryByName(ctx context.Context, request *category.FindCategoryByNameRequest, response *category.CategoryResponse) error {
	curCategory, err := c.CategoryService.FindByName(request.CategoryName)
	if err != nil {
		return err
	}
	return common.SwapTo(response, curCategory)
}

// FindCategoryByID 查找分类ByID
func (c CategoryController) FindCategoryByID(ctx context.Context, request *category.FindCategoryByIDRequest, response *category.CategoryResponse) error {
	curCategory, err := c.CategoryService.FindById(request.CategoryId)
	if err != nil {
		return err
	}
	return common.SwapTo(response, curCategory)
}

// FindCategoryByLevel 查找分类ByLevel
func (c CategoryController) FindCategoryByLevel(ctx context.Context, request *category.FindCategoryByLevelRequest, response *category.FindAllCategoryResponse) error {
	categorys, err := c.CategoryService.FindByLevel(request.CategoryLevel)
	if err != nil {
		return err
	}
	err = common.CategorySliceToResponse(categorys, response)
	if err != nil {
		return err
	}
	return nil
}

// FindCategoryByParent 查找分类ByParent
func (c CategoryController) FindCategoryByParent(ctx context.Context, request *category.FindCategoryByParentRequest, response *category.FindAllCategoryResponse) error {
	categorys, err := c.CategoryService.FindByLevel(request.CategoryParent)
	if err != nil {
		return err
	}
	err = common.CategorySliceToResponse(categorys, response)
	if err != nil {
		return err
	}
	return nil
}

// FindAllCategory 查找所有分类
func (c CategoryController) FindAllCategory(ctx context.Context, request *category.FindAllCategoryRequest, response *category.FindAllCategoryResponse) error {
	categorys, err := c.CategoryService.FindAllCategory()
	if err != nil {
		return err
	}
	err = common.CategorySliceToResponse(categorys, response)
	if err != nil {
		return err
	}
	return nil
}
