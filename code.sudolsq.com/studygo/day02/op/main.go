package main

import "fmt"

//运算符

func main()  {
	var (
		a = 5
		b = 2
	)
	//算数运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独的语句，不能放在=的右边赋值 ==> a = a + 1
	b-- //单独的语句，不能放在=的右边赋值 ==> b = b - 1
	fmt.Println(a)

	//关系运算符
	fmt.Println(a == b)//Go语言是强类型，相同类型的变量才能比较
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a < b)
	fmt.Println(a > b)

	//逻辑运算符
	//如果年龄大于18岁并且年龄小于60岁
	age := 19
	if age > 18 && age < 60{
		fmt.Println("上班的苦逼")
	}else {
		fmt.Println("自由人")
	}
	//如果年龄小于18岁或者年龄大于60岁
	if age < 18 || age > 60{
		fmt.Println("自由人")
	}else {
		fmt.Println("上班的苦逼")
	}
	//not取反
	isMarried := false
	fmt.Println(isMarried)//false
	fmt.Println(!isMarried)//true
}
