package main

import "fmt"
//布尔值
//1、布尔值类型的默认值为false
//2、Go 语言中不允许将整型强制转换为布尔型.
//3、布尔型无法参与数值运算，也无法与其他类型进行转换。

func main()  {
	b1 := true
	var b2 bool//默认值是false
	fmt.Printf("%T\n",b1)
	fmt.Printf("%T value:%v\n",b2 , b2)//%v是打印变量的值，任意类型的变量
}

