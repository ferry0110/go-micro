package repository

import (
	"category/domain/model"
	"github.com/jinzhu/gorm"
)

// CategoryRepository Category数据库实体类
type CategoryRepository struct {
	mysqlDb *gorm.DB
}

// NewCategoryRepository 创建CategoryRepository
func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDb: db}
}

// InitTable 初始化表
func (u *CategoryRepository)InitTable() error{
	return u.mysqlDb.CreateTable(&model.Category{}).Error
}

// FindCategoryById 查询类别byId
func (u *CategoryRepository)FindCategoryById(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	return category,u.mysqlDb.First(category,id).Error
}

// FindCategoryByName 通过类别名查询类别
func (u *CategoryRepository)FindCategoryByName(name string) (category *model.Category, err error)  {
	category = &model.Category{}
	return category,u.mysqlDb.Where("category_name = ?",name).Find(category).Error
}

// FindCategoryByParent 查询类别byParent
func (u *CategoryRepository) FindCategoryByParent(parentId int32) (category []model.Category,err error) {
	return category,u.mysqlDb.Where("category_parent = ?",parentId).Find(category).Error
}

// FindCategoryByLevel 查询类别byLevel
func (u *CategoryRepository) FindCategoryByLevel(levelId int32) (category []model.Category,err  error) {
	return category,u.mysqlDb.Where("category_level = ?",levelId).Find(category).Error
}

// CreateCategory 创建新类别
func (u *CategoryRepository)CreateCategory(category *model.Category) (id int64, err error) {
	return category.Id, u.mysqlDb.Create(category).Error
}

// DeleteCategory 删除类别
func (u *CategoryRepository)DeleteCategory(id int64) error {
	return u.mysqlDb.Where("id = ?",id).Delete(&model.Category{}).Error
}

// UpdateCategory 更新类别信息
func (u *CategoryRepository)UpdateCategory(category *model.Category) error{
	return u.mysqlDb.Model(category).Update(&category).Error
}

// FindAll 查找所有类别
func  (u *CategoryRepository)FindAll() (categorys []model.Category, err error)  {
	return categorys,u.mysqlDb.Find(&categorys).Error
}


















