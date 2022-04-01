package common

import (
	"category/domain/model"
	"category/proto/category"
	"encoding/json"
)

// SwapTo 通过json tag 进行数据的转移
func SwapTo(sourceData, targetData interface{}) (err error)   {
	dataByte,err := json.Marshal(sourceData)
	if err!=nil {
		return
	}
	return json.Unmarshal(dataByte,targetData)
}

// CategorySliceToResponse  将CategorySlice转换为response中的Category
func CategorySliceToResponse(categorySlice []model.Category,response *category.FindAllCategoryResponse) error{
	for _ , cg := range categorySlice{
		cur := &category.CategoryResponse{}
		err := SwapTo(cg,cur)
		if err!= nil {
			return err
		}
		response.Category = append(response.Category,cur)
	}
	return nil
}

