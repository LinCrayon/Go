package main

import "fmt"

func main() {
	//1.编程最简单的算法之一，莫过于变量交换。交换变量的常见算法需要一个中间变量进行变量的临时保存。
	var a int = 100
	var b int = 200
	var t int
	t = a
	a = b
	b = t
	fmt.Println(a, b)

	//2.中间变量的使用

	var c int = 100
	var d int = 200
	c = c ^ d
	d = d ^ c
	c = c ^ d
	fmt.Println(c, d)

	//3.使用 Go 的“多重赋值”特性
	var e int = 100
	var f int = 200
	e, f = f, e
	fmt.Println(e, f)
}
