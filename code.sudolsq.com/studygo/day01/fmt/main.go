package main

import "fmt"

//fmt占位符
func main()  {
	var n = 100
	//查看类型
	fmt.Printf("%T\n",n)//类型
	fmt.Printf("%v\n",n)//查看值(任意)
	fmt.Printf("%b\n",n)//二进制
	fmt.Printf("%d\n",n)//整数
	fmt.Printf("%o\n",n)//八进制
	fmt.Printf("%x\n",n)//十六进制
	var s = "Hello 沙河"
	fmt.Printf("字符串：%s\n",s)//字符串
	fmt.Printf("字符串：%v\n",s)//查看值(任意)  //字符串：Hello 沙河
	fmt.Printf("字符串：%#v\n",s)//显示值和类型  //字符串："Hello 沙河"
}

