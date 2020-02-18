package main

import "fmt"

// 内容：类型

// 相当于类，不作详细解释
// 类型的使用会在方法、函数、结构体部分有更加丰富的使用，这里只引出基本的用法

// 类型定义语句：
// type 类型名称 底层类型

// 对于每一个类型T，都有一个对应的类型转换操作T(x)，用于将x转为T类型
// 只有当两个类型的底层基础类型相同时，才允许这种转型操作，或者是两者都是指向相同底层结构的指针类型，这些转换只改变类型而不会影响值本身。
// 如果T是指针类型，可能会需要用小括弧包装T，比如(*int)(0)

// 底层类型
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// float32 float64 complex128 complex64
// bool byte rune string error

// 除了类似一般语言里面的 类 的用法，在go里面，底层类型不能增加方法，需要通过创建新的类型使用，这里只做简单演示，在方法讲解时候再次提及

// 演示例子
// 定义一个新的 int 类型
type myInt int
type anotherInt int

// 为myInt 增加一个方法
func (m myInt) add() {
	m++
	// 打印m++后的值
	fmt.Println(m)
}

func f1() {
	// myInt()是类型转换操作，不是函数调用。不会改变值本身，但会改变值的类型
	a := myInt(3)
	a.add()
	// 拓展内容：这时的a 应该是多少？
	// fmt.Println(a)

	// 类型转换例子
	b := anotherInt(5)
	c := myInt(b)
	fmt.Printf("%T,%v\n", b, b)
	fmt.Printf("%T,%v\n", c, c)
}

func main() {
	f1()
}
