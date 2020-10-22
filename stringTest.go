package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name string
	age int
	sex string
}

func main() {


	var str = "how do you do"
	var strSlice = strings.Split(str," ")
	fmt.Println(strSlice)

	var strMap = make(map[string]int)
	for _, v := range strSlice{
		strMap[v]++
	}
	fmt.Println(strMap)

	//求两数之和
	sum1 := sumFn(12, 8)
	fmt.Println(sum1)

	sumFn1(1,3,5)

	//函数作为参数传递
	result := calc(5,2,sub)
	fmt.Println(result)

	//闭合函数
	fn := adder()
	fmt.Println(fn(10))
	fmt.Println(fn(10))
	fmt.Println(fn(10))


	//fmt.Println("开始")
	//defer fmt.Println("1")//延时执行
	//defer fmt.Println("2")
	//defer fmt.Println("3")
	//fmt.Println("结束")


	var p1 Person
	p1.name = "marc"
	p1.age = 21
	p1.sex = "men"


	fmt.Printf("值%#v 类型%T\n",p1,p1)

	//第二种创建结构体的方法,指针类型
	var p2 = new(Person)
	p2.name = "kris"
	p2.age = 22
	p2.sex = "men"
	//可以使用指针直接操作变量体里面的值
	(*p2).name = "koma"
	fmt.Printf("值%#v 类型%T\n",p2,p2)

	//第三种创建结构体的方法，指针类型
	var p3 = &Person{}
	p3.name = "tomas"
	p3.age = 22
	p3.sex = "men"
	//可以使用指针直接操作变量体里面的值
	(*p3).name = "tomase"
	fmt.Printf("值%#v 类型%T\n",p3,p3)

	//第四种创建结构体的方法
	var p4 = Person{
		name: "lux",
		age: 32,
		sex: "gril",
	}
	fmt.Printf("值%#v 类型%T\n",p4,p4)

	//第5种创建结构体的方法,指针类型
	var p5 = &Person{
		name: "lux",
		age: 32,
		sex: "gril",
	}
	fmt.Printf("值%#v 类型%T\n",p5,p5)

}
func sumFn(x int,y int) int{
	sum := x+y
	return sum
}

func sumFn1(x ...int)  {//将输入变成切片
	fmt.Printf("%v--%T",x,x)
}

func sub(x ,y int) int{
	sub := x-y
	return sub
}

type calcType func(int,int) int
//函数作为参数传递
func calc(x, y int, cb calcType) int {
	return cb(x,y)
}

//闭包 函数里面嵌套一个函数 最后返回里面的函数
//使变量常驻内存，不污染全局
func adder() func(y int)int  {
	var i = 10
	return func(y int)int{
		i += y
		return i
	}
}



