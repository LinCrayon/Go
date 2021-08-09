package main

import "fmt"

//for循环
func main() {
	//基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//变种1
	var i1 = 5
	for ; i1 < 10; i1++ {
		fmt.Println(i1)
	}
	//变种2
	var i2 = 5
	for i2 < 10 {
		fmt.Println(i2)
		i2++
	}
	//无线循环
	/*for {
		fmt.Println("123")
	}*/

	//for range循环
	s := "Hello 沙河"
	for i , v := range s{	//索引
		fmt.Printf("%d %c\n",i ,v)
	}
}