package main

import (
	"fmt"
	"strings"
	"study/class8/code1/types"
)

// 方法类似函数，但是方法只属于某个或者某些类型，只能通过类型来调用
// 方法也是类型的一个成员，在等下说到的接口部分也会被作为接口内容使用
// 简单理解，就像面向对象的编程里面 类 的方法 比如python 里面的class 定义的方法

// 分隔符号
var split string = strings.Repeat("-", 80)

// 方法的继承、使用例子
func f1() {
	s := types.NewStudent()
	// 更改名字
	s.RenameFailed("failed")
	fmt.Println("rename failed", s.Name())
	s.Rename("new jszc")
	fmt.Println("rename success", s.Name())

	fmt.Println(split)

	s.Run()
}

func f2() {
	// f := func(p PlayGame) {
	// 	fmt.Println(p.Play())
	// }
	f := func(l types.LOL) {
		fmt.Println(l.Play())
	}
	lol := types.NewLOL()
	f(lol)
	// dota := types.NewDOTA()
	// f(dota)
}

func main() {
	f1()
}
