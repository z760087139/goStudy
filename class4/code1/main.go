package main

import (
	"fmt"
	"strings"
)

// 内容：数组、切片  make new函数

// 数组是一个固定长度的特定数据类型序列，长度在定义时候决定，不能修改
// 切片是一个可以调整长度的数组，go语言里面比较常用，但切片的理解需要基于数组

// 定义数组
func f1() {
	// 完整的定义语法
	// var name [number]type
	// 如果赋值的内容 [number]type{...} ...的内容的数量小于number，则会以type数据类型对应的零值存值
	// 数据类型对应的零值列表
	// int			0
	// string		""
	// bool			false
	// map			nil
	var a1 [3]int = [3]int{1, 2, 3}
	// 需要注意 [3]int 部分是对a1的数据类型进行定义，说明 a1 的数据类型时[3]int，但不包括列表数据内容
	// [3]int{1,2,3}代表的是一个由 1,2,3组成的长度为3的数组，然后把这个数组赋值给a1
	// 不能直接使用 {1,2,3}赋值，简化的定义方式如下：

	// 简化定义方式
	a2 := [3]string{"a", "b"}
	// 简化了 var 和 a2 的数据类型定义，只有 [3]int{1,2,3} 这个数组的内容赋值给 a2 的过程，但是会隐式给a2定义数据类型为 [3]int

	// 普通定义及赋值形式 相当于a1 步骤的拆分
	var a3 [3]bool
	a3 = [3]bool{true, false, false}

	// 如果列表的长度需要根据内容来决定
	a4 := [...]int{1, 2, 3, 4} //长度为4，且列表顺序会按照内容排序。
	// 还有另外一种形式：
	a5 := [...]int{1: 10, 2: 1, 5: 9} // 长度为6，内容按照所写的位置进行赋值，为定义的内容都会以零值形式赋值

	// 打印a1 a2 a3的数据类型、值、指针地址
	fmt.Printf("%[1]T,%[1]v,%[2]p\n", a1, &a1)
	fmt.Printf("%[1]T,%[1]v,%[2]p\n", a2, &a2)
	fmt.Printf("%[1]T,%[1]v,%[2]p\n", a3, &a3)
	fmt.Printf("%[1]T,%[1]v,%[2]p\n", a4, &a4)
	fmt.Printf("%[1]T,%[1]v,%[2]p\n", a5, &a5)

}

// 修改数组内容
// 数组长度固定，不能为数组新增内容，只能修改内容
func f2() {
	a1 := [...]int{1, 2, 3, 4}
	// 修改前的内容
	fmt.Println("before : ", a1)
	// 修改索引位置为 0 的内容
	a1[0] = 10
	fmt.Println("after : ", a1)
	// 当数组作为参数进行传递并修改时，会发生什么事情
	f3(a1)
	fmt.Println("array out of f3", a1)

	f4(&a1)
	fmt.Println("array out of f4", a1)
	// 注意函数的区别，并结合第二节课的内容 class2/code1/main.go:130 进行理解
}

// 通过函数修改数组一
func f3(a [4]int) {
	a[0] = 20
	fmt.Println("array in f3", a)
}

// 通过函数修改数组二
func f4(a *[4]int) {
	a[0] = 30
	fmt.Println("array in f4", *a)
}

// -----------------------------------

// 切片(slice)
// 在go语言里面，数组在实际使用过程中，由于固定长度的原因，并没有切片使用来得广泛
// 基本上在go里面想表达一个序列，基本都使用切片。
// 切片内容包括指针、长度、容量
// 切片实际上是通过指针引用了数组的内容，每个切都有对应的数组，存储数据内容
// 长度是当前切片的指针个数
// 容量是当前切片引用的数组，从切片第一个指针在数组中的位置，到数组最后一个元素为止的长度

// 实际例子
func f5() {
	// https://books.studygolang.com/gopl-zh/ch4/ch4-02.html
	// 创建一个长度为10，内容为对应位置的数组
	a2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 对a1进行切片操作
	// arrary/slice[begin:end:caps] 切片内容从begin开始，包括begin，到end结束，不包括end
	// 和python 不同，切片的最后一个并不是步长，是表示新创建的slice的容量
	// 不填写begin,end的话，默认是第一个元素和len(s)位置元素（不是caps），具体的演示看下面的例子
	// 重点内容 ： go 里面的数组、切片内容是存放在连续的内存地址中的

	// 源数组内容指针
	// 打印源数组各个值的地址
	f6(a2)

	// 获取a1的完整内容，但是数据类型是一个切片
	// si1 := a2[:]
	// 打印切片内容
	// f7(si1)

	// 如果是一个片段的切片内容,有头无尾
	// si2 := a2[7:]
	// f7(si2)

	// 无头有尾，注意长度和容量
	si3 := a2[:7]
	f7(si3, "")
	// f7(si3[:], "")
	// 有头有尾，注意长度和容量
	// si4 := a2[2:7]
	// f7(si4)

	// 如果循环打印切片长度以外的数据（容量以内）会报错
	// for i := 0; i < cap(si4); i++ {
	// 	fmt.Print(&si4[i], "\t")
	// }
	// fmt.Println()

}

func f6(a2 [10]int) {
	fmt.Println(strings.Repeat("-", 80))
	fmt.Print("源数组内容指针:\t")
	for i := 0; i < len(a2); i++ {
		fmt.Print(&a2[i], "\t")
	}
	fmt.Println()
	fmt.Print("数据值\t")

	for i := 0; i < len(a2); i++ {
		fmt.Print(a2[i], "\t")
	}
	fmt.Println()
}

func f7(s []int, content string) {
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println(content)
	fmt.Println("长度：", len(s))
	fmt.Println("容量: ", cap(s))
	fmt.Print("各数据指针\t")

	for i := 0; i < len(s); i++ {
		fmt.Print(&s[i], "\t")
	}
	fmt.Println()

	fmt.Print("各数据值\t")

	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], "\t")
	}
	fmt.Println()

}

// 切片操作
func f8() {
	// 创建原数组
	a2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 对a2进行切片，采集一个长度小于容量的数据
	s := a2[2:6]
	// 打印当前的内存情况
	f7(s, "打印初始切片")
	// 修改切片内容
	s[0] = 10
	f7(s, "修改[0]后的切片")

	// 增加切片内容，需要使用内建函数 append
	// append 函数会对输入的参数进行累加，并返回结果，但对传入的参数不作变更
	// 所以如果想对原本的变量内容进行修改，需要接收返回值
	s = append(s, 11)
	// 注意内存地址
	f7(s, "append一个元素后的切片")
	// 打印原内容
	f6(a2)

	// 如果重新赋值一个变量，留意内存地址是否发生变化
	ns := append(s, 12)
	f7(ns, "append新切片")

	// 如果append的内容比原本的容量要多
	ns = append(s, 12, 13, 14, 15, 16)
	f7(ns, "append 超过容量")
}

// 关于make
func f9() {
	// 正如数组的定义，对一个变量进行定义之后再进行内容的赋值。
	// 但是如何自行初始化一个切片的长度、容量
	// make 函数用于对这种 结构体 的初始化
	// make 函数对 slice map chan 类型的进行初始化并返回对应的值 map chan 类型 以后章节也会用到，这里不作演示
	// 用法 make[type,len,cap]

	// s 变量定义为一个 []int的切片，但是还没赋值内容
	var s []int
	// 只能定义为一个长度3，容量3的切片
	s = []int{1, 2, 3}
	f7(s, "通过固定内容赋值")
	// 自定义一个切片的长度，容量（但是内容都为零值）
	s1 := make([]int, 3, 10)
	// 等价于
	// var s1 []int
	// s1 = make([]int,0,10)
	f7(s1, "make初始化")
}

// new函数
func f10() {
	// 和make 函数类似
	// new 函数同样用于初始化，但内容不限制类型，返回的内容为指针类型
	// new 是对传入的类型分配空间并返回对应的内存地址指针
	n1 := new(int)
	fmt.Printf("%p", n1)
	fmt.Println()
	// 等价于

	var v int
	n2 := &v
	fmt.Printf("%p", n2)

}

func main() {
	f10()
}
