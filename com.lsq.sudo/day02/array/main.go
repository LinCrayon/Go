package main

import "fmt"

//数组

//存放元素的容器
//必须指定存放的元素类型和容量(长度)
//数组的长度是数组类型的一部分
func main()  {
	var a1 [3]bool
	var a2 [4]bool

	fmt.Printf("a1:%T a2:%T\n", a1 , a2)

	//数组的初始化
	//如果不初始化：默认元素都是零值(布尔值：false,整形和浮点型都是0,字符串："")
	fmt.Println(a1 , a2)
	//初始化方法一
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	//初始化数组方法二
	//根据初始值自动推断数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a10)
	//初始化方法三:根据索引来初始化
	a3 := [5]int{0: 1 , 4: 2}
	fmt.Println(a3)

	//数组的遍历
	citys := [...]string{"北京", "上海", "深圳"}
	//1.根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//2.for range遍历
	for i, v := range citys{
		fmt.Println(i , v)
	}

}
