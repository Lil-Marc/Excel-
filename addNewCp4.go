package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)
//甜橙
func main() {
	f, err := excelize.OpenFile("【清单】甜橙授权米读60本原创（已验证）.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	//newBookIdMap := make(map[string]interface{}, 0)
	// Get all the rows in the Sheet1.
	rows := f.GetRows("Sheet1")


	for i, row := range rows {
		//按照格式输出
		if i > 0 {
			//fmt.Println(row[1])//作者
			//fmt.Println(row[0])//title
			//fmt.Println(row[2])//字数
			//fmt.Println(row[8])//简介
			fmt.Printf("    {\n        \"author\": \"%v\",\n        \"description\": \"%v\",\n        \"id\": \"%v\",\n        \"title\": \"%v\"\n    },\n",row[1],row[8],i+500,row[0])
		}

	}



}
