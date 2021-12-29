package main

import "fmt"

//在编码过程中，可能会遇到没有名称的变量、类型或方法。虽然这不是必须的，但有时候这样做可以极大地增强代码的灵活性，这些变量被统称为匿名变量。
//匿名变量的特点是一个下画线“_”，“_”本身就是一个特殊的标识符，被称为空白标识符。
//它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给它），但任何赋给这个标识符的值都将被抛弃，
//因此这些值不能在后续的代码中使用，也不可以使用这个标识符作为变量对其它变量进行赋值或运算。使用匿名变量时，只需要在变量声明的地方使用下画线替换即可。
func GetData() (int, int) {
	return 100, 200
}
func main() {
	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)
}

//匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。