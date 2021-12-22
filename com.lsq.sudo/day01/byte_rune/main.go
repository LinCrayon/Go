package main

import "fmt"
//byte和rune类型
//Go语言中为了处理非ASCII码类型的字符 定义了新的rune类型
func main(){
	s := "hello"
	//len()求得byte字节的数量
	n := len(s)
	fmt.Println(n)

	for _, c :=range s {//从字符串中拿出具体的字符
		fmt.Printf("%c\n",c)//%c:字符
	}
	//字符串的修改
	s2 := "小薇"
	s3 := []rune(s2)//把字符串强制转换成了一个切片
	s3[0] = '大'
	fmt.Println(string(s3))//把rune切片强制转换成字符串

	c1 := "薇"	//string
	c2 := '薇'	//int32
	fmt.Printf("c1:%T c2:%T\n" , c1 , c2)

	//类型转换
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n",f)

}

//1、uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
//1、rune类型，代表一个 UTF-8字符。