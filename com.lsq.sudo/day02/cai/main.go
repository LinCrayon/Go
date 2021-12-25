//package是关键字，声明作用，要是main包会生成可执行文件
package main

//导入语句
import "fmt"

//函数外面只能放置标识符(变量\常量)的声明
//fmt.Println("Hello")//非法

//程序的入口函数
func main() {
	fmt.Println("helloWorld")
}
