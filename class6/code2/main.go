package main

import (
	"fmt"
	"strings"
)

// go 语言关于流程控制有下面几种方式
// 1.for 循环  （没有while)
// 2.条件判断  if / switch / select
// 3.跳转      continue / break / goto(tag)

// for 语法
// 完整语法
// for 初始化;条件子语句;后置子语句{循环内容}

// 简化写法
// 1. 忽略后置子语句
// for 初始化;条件子语句;{...}

// 2. 忽略初始化，后置子语句
// for 条件子语句{}

// 3. 忽略初始化
// for ;条件子语句;后置子语句{}

// for 例子
func f1() {
	i := 0

	// 完整命令
	for i := 0; i < 4; i++ {
		fmt.Println("i in for 1", i)
	}

	// 是否改变了i
	fmt.Println("i after for 1", i)

	// 其他形式
	for i := 0; i < 4; {
		i++
		fmt.Println("i in for 2", i)
	}

	for ; i < 4; i++ {
		fmt.Println("i in for 3", i)
	}
	// 是否改变了i
	fmt.Println("i after for 2", i)

	for i < 6 {
		i++
		fmt.Println("i in for 4", i)
	}
	fmt.Println(strings.Repeat("-", 80))

	// 关于 for range 部分，在第二节课应该有相关内容说明，这边只作简单回顾
	// 通过 for range 语法，返回当前循环的索引位置和值内容 或者值返回索引位置
	b := []int{1, 2, 3}
	for index, value := range b {
		fmt.Println("current index :", index, "value", value)
	}
	for index := range b {
		fmt.Println("current index :", index)
	}
	fmt.Println(strings.Repeat("-", 80))

	// 同样可以用for range 遍历 map ，返回的为key 和value
	// 但是由于map 是hash表结构，每次返回的key ,value 是随机的
	c := map[int]string{0: "a", 1: "b", 2: "c", 3: "d"}
	for key, value := range c {
		fmt.Println("key :", key, "value :", value)
	}

	// Range expression                          1st value          2nd value

	// array or slice  a  [n]E, *[n]E, or []E    index    i  int    a[i]       E
	// string          s  string type            index    i  int    see below  rune
	// map             m  map[K]V                key      k  K      m[k]       V
	// channel         c  chan E, <-chan E       element  e  E
}

// 条件判断

// if 语句
// if 初始化;条件子语句{} 初始化可忽略
// if 跟一般的语言中的if 相同，区别在于 go语言没有python 中的elif ，需要紧接在if 花括号后使用else 或者 else if
func f2() {
	a := 0
	if a >= 1 {
		fmt.Println("if 1 a >= 1 ")
	} else if a <= -1 {
		fmt.Println("if 1 a <= 1")
	} else {
		fmt.Println("if 1 else and make a = 10")
		a = 10
	}
	if a := -1; a < 1 {
		fmt.Println("if 2 a < 1")
	} else {
		fmt.Println("if 2 a > 1")
	}
	// 是否改变了a
	fmt.Println(a)
}

// switch 语法
// switch 初始化/指定某个变量 {
// case 条件 :
// }
// switch 相当于多个if else ... 有一个关键字 fallthrough
func f3() {
	grade := 80
	switch grade {
	// case <60: 指定变量之后，不能用表达式
	case 60:
		fmt.Println("不及格")
	case 70:
		fmt.Println("70分")
	default:
		fmt.Println("及格")
	}
	switch {
	case grade <= 60:
		fmt.Println("不及格")
	case grade > 60:
		fmt.Println("及格")
		fallthrough
	case grade <= 80:
		fmt.Println("良")
		// fallthrough
	default:
		fmt.Println("优秀")
	}

	// switch 还可以用于断言，该部分在接口部分详细说明
}

// select 语法
// select
// select 与 switch 类似，但是select 采用随机执行方式，当随机一个条件可以执行，就会按照该方式执行
// 一般用于管道传输或者超时等待方面使用，该部分在goroutine 管道部分说明，例子可以自行尝试
func f4() {
	a := make(chan int, 1024)
	b := make(chan int, 1024)

	for i := 0; i < 10; i++ {
		fmt.Printf("第%d次", i)
		a <- 1
		b <- 1

		select {
		case <-a:
			fmt.Println("from a")
		case <-b:
			fmt.Println("from b")
		}
	}
}

// 流程跳转、中断 标签
// break ,continue,goto
// break 能在switch /select 流程中被应用。嵌套循环中只能终止当前的循环，不能终止外层的循环
// 不推荐使用label 和 goto 方式，容易导致混乱
// 如果想查看label 的详细内容，可以看下面的页面介绍
// https://medium.com/golangspec/labels-in-go-4ffd81932339
func f5() {
	i := 0

	for i := 0; i < 10; i++ {
		fmt.Println("in for, times:", i)
		switch {
		case i < 2:
			fmt.Println("in switch, times:", i)
		case i == 20:
			goto LAB
		default:
			continue
			// break
			// fmt.Println("something")
		}
	}
	fmt.Println("finish")

LAB:
	fmt.Println("goto times :", i)
	if i < 3 {
		i++
		goto LAB
	}
}

func main() {
	f5()
}
