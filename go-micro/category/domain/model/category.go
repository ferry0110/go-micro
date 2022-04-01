package model

type Category struct {
	Id                  int64  `gorm:"primary_key;not_null;auto_increment"`
	CategoryName        string `gorm:"unique_index;not_null" json:"category_name"`
	CategoryLevel       int32  `json:"category_level"`
	CategoryParent      int32  `json:"category_parent"`
	CategoryImage       string `json:"category_image"`
	CategoryDescription string `json:"category_description"`
}
