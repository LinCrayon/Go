package main

import "fmt"

//if条件判断
func main()  {
	/*age := 19
	if age > 18 {
		fmt.Println("允许进入赌场")
	}else {
		fmt.Println("滚蛋")
	}*/
	age := 19
	if age := 19 ; age > 18 {
		fmt.Println("允许进入赌场")
	}else {
		fmt.Println("滚蛋")
	}
	//fmt.Println(age)注意作用域

	//多个判断条件
	if age > 35 {
		fmt.Println("中年")
	}else if age > 18 {
		fmt.Println("青年")
	}else{
		fmt.Println("小屁孩")
	}
}
