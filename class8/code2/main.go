package main

import (
	"study/class8/code2/types"
)

// 关于接口
// 接口类型是一种抽象的类型。它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合；它们只会表现出它们自己的方法。
// 也就是说当你有看到一个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的方法来做什么。

// 官方解析
// An interface type specifies a method set called its interface.
// A variable of interface type can store a value of any type with a method set that is any superset of the interface.
// Such a type is said to implement the interface. The value of an uninitialized variable of interface type is nil.

// 简单翻译
// 接口是方法的集合
// 只要实现了指定的方法，任意类型都能被接口接收(转成接口类型)
// 未初始化的接口为nil

// 接口的基本定义语法
// interface { "MethodName" | "InterfaceName" ;  }
// 即 接口内容可以是方法名、接口名

// 接口调用例子
func f1(p types.Person) {
	p.Eat("rich")
	// p.Run()  // 由于接口定义没有说明持有Run 方法，
	p.Talk()
}

// 使用例子
func f2() {
	s := types.Student{
		// name: "jszc",
	}
	f1(s)
}

// 指针方法使用例子
func f3() {
	m := types.Man{
		// name: "jszc",
	}
	// f1(m) // 无法执行
	f1(&m)
}

func main() {
	f2()
}
