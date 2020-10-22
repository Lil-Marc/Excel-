package main

import (
	"fmt"
)


func main() {
	//var str = "123-456-789"
	//arr := strings.Split(str,"-")
	//fmt.Println(arr)//简单理解切片就是数组，在golang中切片和数组
	//
	//str1 := strings.Join(arr,"*")
	//fmt.Println(str1)
	//
	//s := "hello张三"
	//for _,i := range s {//通过rune循环
	//	fmt.Printf("%v(%c)\n",i,i)
	//}

	//str3 := "1234"
	//fmt.Printf("%v--%T\n",str3,str3)
	//
	//
	////_匿名变量代表error，转化为整数int
	//num,_ := strconv.ParseInt(str3,10,64)
	//fmt.Printf("%v--%T\n",num,num)
	//
	////_匿名变量代表error，转化为float
	//floats,_ := strconv.ParseFloat(str3,64)
	//fmt.Printf("%v--%T\n",floats,floats)
	//
	////余数=被除数 - （被除数/除数） * 除数
	//fmt.Println(-10%3)

	//golang中赋值运算符不能与自增自减同时使用

/*	var a = 12
	b := a++
	fmt.Println(b)*/

	//var a = 12
	//a++
	//fmt.Println("a=",a)
	//
	//for i := 0; i <= 10; i++ {
	//	fmt.Println(i)
	//}
	//
	//var arr = []string{"php","java","node","golang"}
	//
	//for key,val := range arr{
	//	fmt.Println(key,val)
	//}
	
/*	var arr = [...]int{-1,2,12,65,11}
	var max = arr[0]
	var index = 0
	for i := 0;i < len(arr);i++{
		if  max < arr[i]{
			max = arr[i]
			index = i
		}
	}
	fmt.Printf("最大值:%v 最大值对应的索引值:%v",max,index)*/

	//基于数组定义切片
	a := [5]int{12,15,45,55}
	b := a[:]//获取数组里的所有值

	fmt.Printf("%v-%T\n",b,b)

	sliceA := []int{1,2,3,45}
	sliceB := make([]int,4,4)
	copy(sliceB,sliceA)
	fmt.Println(sliceA)
	fmt.Println(sliceB)
	sliceB[0] = 222
	fmt.Println("修改后:")
	fmt.Println(sliceA)
	fmt.Println(sliceB)

	var userinfo = map[string]string{
		"username":"marc",
		"age":"2O",
		"sex":"boy",
	}

	for k,v  := range userinfo{
		fmt.Printf("key:%v value:%v\n",k,v)
	}





}