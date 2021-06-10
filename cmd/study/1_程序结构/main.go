package main

import "fmt"

func main() {
	// 1、声明
	// 声明：var const type func
	var i, j = 123, 234
	fmt.Println(i, j)

	var x, y string
	fmt.Println(x, y)

	// 短变量声明
	m, n := 1.23, 2.34
	fmt.Println(m, n)

	// 2、指针
	a := &m // 加上了&才是指针类型
	*a = 3.45
	fmt.Println(&a)
	fmt.Println(*a)

	// 3、new 函数
	var c = new(int) // 返回地址
	*c = 12
	fmt.Println(c)
	fmt.Println(*c)

	// 4、多重赋值
	i, j, m = 1, 2, 3
	fmt.Println(i, j, m)

	// 5、类型声明
	type Celsius float64
	var absoluteZeroC Celsius = -273.15
	fmt.Println(absoluteZeroC)
}
