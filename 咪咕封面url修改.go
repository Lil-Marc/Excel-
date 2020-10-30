package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)
func main() {
	f, err := excelize.OpenFile("咪咕分成封面.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	//newBookIdMap := make(map[string]interface{}, 0)
	// Get all the rows in the Sheet1.
	rows := f.GetRows("SheetJS")


	for i, row := range rows {
		//按照格式输出
		if i >= 0   {
			//fmt.Println(row[0])
			fmt.Printf("UPDATE `sync_book` SET `cover`  = 'https://midu-book-image.oss-cn-beijing.aliyuncs.com/book/coverbyyunpeng/%v.jpg' WHERE `cover`  = '%v';\n",i,row[0])
		}

	}



}
