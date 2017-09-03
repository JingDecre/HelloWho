package main

import "fmt"

func SayHello() {
	var age int = 12
	v_name := "12"
	v_name1, v_name2, v_name3 := "a", "b", "c"
	println(age, v_name, v_name1, v_name2, v_name3)
	fmt.Println("hello decre!")
}

var x, y int
var ( //这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"
func MultipleVariables() {
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)

}
