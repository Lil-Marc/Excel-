package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)
func main() {
	f, err := excelize.OpenFile("咪咕分类映射.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	//newBookIdMap := make(map[string]interface{}, 0)
	// Get all the rows in the Sheet1.
	rows := f.GetRows("Sheet1")


	for i, row := range rows {
		//按照格式输出
		if i > 0  {
			//fmt.Println(row[2])//cp方分类
			//fmt.Println(row[3])//米读一级分类ID
			//fmt.Println(row[5])//米读二级分类ID
			//fmt.Println(row[7])//米读三级分类ID
			fmt.Printf("        {\n            \"copyright_category\":\"%v\",\n            \"stack_first_category_id\":%v,\n            \"stack_second_category_id\":%v,\n            \"stack_third_category_id\":%v\n        },\n",row[2],row[3],row[5],row[7])
		}

	}



}
