package main

import (
	"fmt"
	"strings"
)

//字符串

func main()  {
	path1 := "\"C:\\Code\\Go\\src\\code.sudolsq.com\\studygo\\day01\""
	path2 := "'C:\\Code\\Go\\src\\code.sudolsq.com\\studygo\\day01'"
	fmt.Println(path1)
	fmt.Println(path2)

	s := "I'm ok"
	fmt.Println(s)

	//多行的字符串
	s2 := `
	hhh
	  lll
	xxx
	`
	fmt.Println(s2)

	s3 := `C:\Code\Go\src\code.sudolsq.com\studygo\day01`//(``)是反引号
	fmt.Println(s3)

	//字符串的相关操作
	fmt.Println(len(s3))//字符串长度

	//字符串拼接
	name := "linshengqian"
	world := "jiexiaowei"

	ss := name +world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s",name , world)
	//fmt.Println("%s%s",name , world)
	fmt.Println(ss1)

	//分隔
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)

	//包含
	fmt.Println(strings.Contains(ss,"linshengqian"))
	fmt.Println(strings.Contains(ss,"jiexiaowei"))

	//前缀和后缀判断
	fmt.Println(strings.HasPrefix(ss,"linshengqian"))
	fmt.Println(strings.HasSuffix(ss,"linshengqian"))

	//字符出现的位置
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4,"c"))
	fmt.Println(strings.LastIndex(s4,"b"))

	//拼接
	fmt.Println(strings.Join(ret,"+"))

}
