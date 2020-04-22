package main

import "fmt"

// 方法类似函数，但是方法只属于某个或者某些类型，只能通过类型来调用
// 方法也是类型的一个成员，在等下说到的接口部分也会被作为接口内容使用
// 简单理解，就像面向对象的编程里面 类 的方法 比如python 里面的class 定义的方法

// 举个例子，比如所有的Person 类型都可以执行Run 方法
type Person struct {
	name string
	age  int
}

// 方法定义跟函数相似，只是在前面先注释这个方法属于哪个类型
// (self Person) 这个用于说明这个方法属于 Person 类型，self 是Person 类型的简称，用户函数内部表示Person 类型自身
// self 的命名不固定，一般常用的有self 或者 是类型的第一个字母 (比如Person 用 p)表示
func (self Person) Run() {
	fmt.Println(self.name, "is running")
}
