package main

import (
	"fmt"
	"unsafe"
)


//常量枚举
//常量可以用len(),cap(),unsafe.Sizeof()常量计算表达式的值
//此为全局变量，当其他go文件已存在相同全局变量命名时会报错
//在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。
//第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；
const (
 	Unknown = 0
	Female  = 1
	Male    = 2
	a1      = "abc"
	b1      = len(a1)
	c1      = unsafe.Sizeof(a1)
)

func ConstVariables() {
	const length int = 10
	const width int = 5
	var area int
	//const a, b, c = 1, false, "str"

	area = length * width

	fmt.Printf("面积为：%d", area)
	println()
	println(a1, b1, c1)
}
