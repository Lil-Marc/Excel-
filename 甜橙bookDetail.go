package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"strconv"
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
			//id , _ := strconv.Atoi(row[0])//书籍序号从501开始，查出的id+500
			wordCount , _ := strconv.ParseFloat(row[2],64)//小说字数转化为float
			wordCount = wordCount

			content := fmt.Sprintf("{\n    \"id\": %v,\n    \"title\": \"%v\",\n    \"author\": \"%v\",\n    \"description\": \"%v\",\n    \"cover\": \"http://img.midukanshu.com/browser/copyright/tianchengData/%v.jpg\",\n    \"wordCount\": %v\n}",i+500,row[0],row[1],row[8],row[0],wordCount)
			//fmt.Printf("{\n    \"id\": %v,\n    \"title\": \"%v\",\n    \"author\": \"%v\",\n    \"description\": \"%v\",\n    \"cover\": \"http://img.midukanshu.com/browser/copyright/tianchengData/%v\",\n    \"wordCount\": %v\n}",i+500,row[0],row[1],row[8],row[0],wordCount)
			filePath := fmt.Sprintf("/Users/marc/Documents/java课/java自学/廖雪峰/src/甜橙/章节信息/%v/bookDetail.txt",row[0])

			response := ConvertToByte(string(content), "utf8", "utf-8")
			//fmt.Println(response)



			//写入
			err := ioutil.WriteFile(filePath, response,0644)
			if err != nil {
				panic(err)
			}
		}


	}

}
//src为要转换的字符串，srcCode为待转换的编码格式，targetCode为要转换的编码格式
func ConvertToByte(src string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}
