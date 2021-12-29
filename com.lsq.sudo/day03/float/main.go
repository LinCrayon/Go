package main

//o语言提供了两种精度的浮点数 float32 和 float64，它们的算术规范由 IEEE754 浮点数国际标准定义，该浮点数规范被所有现代的 CPU 支持。
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
}
