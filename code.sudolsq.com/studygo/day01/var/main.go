package main

import "fmt"

//Go语言推荐用驼峰式命名（小驼峰）

//声明变量
//变量声明格式为：var 变量名 变量类型
//var  name string
//var age int
//var isOk bool

//批量声明
var (
	name string //""
	age  int    //0
	isOk bool   // false
)

func main() {
	name = "crayon"
	age = 18
	isOk = true
	//Go语言中非全局变量声明后必须使用，不使用就编译不过去
	fmt.Print(isOk) //打印
	fmt.Println()
	fmt.Printf("name:%s\n", name) //Printf格式化，%s是占位符，使用name值去替换
	fmt.Println(age)              //打印完指点内容后在后面加一个换行符号

	fmt.Println("===================================================")

	//声明变量同时赋值
	var s1 string = "sudo"
	fmt.Println(s1)
	//类型推到(根据判断该变量是声明类型)
	var s2 = "20"
	fmt.Println(s2)

	fmt.Println("==================================================")

	//简短变量声明,只能在函数里面用
	s3 := "哈哈哈"
	fmt.Println(s3)

	fmt.Println("==============================")

	//匿名变量（anonymousvariable）。 匿名变量用一个下划线_表示
	//匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明

	//1、函数外的每个语句都必须以关键字开始（var、const、func等）
	//2、:=不能使用在函数外。
	//3、_多用于占位，表示忽略值。
}

//同一个作用域（{}）中不能声明同名的变量
