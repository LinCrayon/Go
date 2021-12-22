package main

import "fmt"

//常量，定义了就不能修改了
const pi = 3.1415926

//批量声明常量
const (
	statusOK = 200
	notFound = 404
)
//批量声明常量，如果省略了值则表示和上面一行的值相同。
const (
	n1 = 100
	n2
	n3
)

//iota是go语言的常量计数器，只能在常量的表达式中使用。(类似枚举)
//iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。
const (
	a1 = iota	//0
	a2 = iota	//1
	a3 = iota	//2
)

const (
	b1 = iota  //0
	b2			//1
	_			//2
	b3			//3
)

//插队
const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota  //2
	c4		  //3
)

//多个产量声明在一行
const (
	d1 , d2 = iota + 1 ,iota + 2 //d1：1  d2 :2  d3:2  d4:3
	d3 , d4 = iota + 1 ,iota + 2
)

//定义数量级
const (
	_  = iota					//扔了
	KB = 1 << (10 * iota) 		//1左移十位，2的10次方
	MB = 1 << (10 * iota)		//1左移二十位，2的20次方
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)


func main()  {
	//pi = 123
	fmt.Println("n1", n1)		//按ALT
	fmt.Println("n2", n2)
	fmt.Println("n3", n3)

	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
	fmt.Println("a3", a3)

	fmt.Println("b1", b1)
	fmt.Println("b2", b2)
	fmt.Println("b3", b3)

	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	fmt.Println("c3", c3)
	fmt.Println("c4", c4)

	fmt.Println("d1", d1)
	fmt.Println("d2", d2)
	fmt.Println("d3", d3)
	fmt.Println("d4", d4)
}