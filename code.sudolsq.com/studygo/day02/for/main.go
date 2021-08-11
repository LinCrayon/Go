package main

import "fmt"

//流程控制跳出for循环

func main()  {
	//当i=5时就跳出for循环(break)
	for i := 0; i < 10; i++ {
		if i == 5 {
			break//跳出for循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	//当i=5时就跳过此次for循环(continue)
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue//继续下一次循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}