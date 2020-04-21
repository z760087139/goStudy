package main

import (
	"fmt"
	"os"
)

// 关于defer 函数的使用

// defer 用于延迟调用。比如在函数执行完成后进行资源回收
// 多个defer 的执行顺序是按照先进后出的堆叠方式

// defer 的一般应用
func f1() {
	// f, err := os.Open("test.txt")
	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 665)
	// 错误处理
	if err != nil {
		fmt.Println(err)
		return
	}
	// 此时不会执行文件关闭，但是在f1函数执行完成后，会执行 f.Close()
	defer f.Close()

	f.WriteString("aaaa")

}

// 关于 defer 的执行顺序  非常重要
// 理解重点：
// 1.GO 语言中 return 不是原子操作，分为 “返回值” 赋值 和 RET 返回指令
// 2.defer 动作在 “返回值” 赋值后    RET 返回指令前
// 即 ： “返回值”赋值  →  defer → RET返回
// 3.不同形式的defer 调用，产生的效果不一样，一定要注意func()的变量作用于使用
// 4.defer 动作在defer 定义的地方会完成参数的传递

// defer func 传参形式与内部变量形式
func f2() {
	n := 1
	// 这里的n 是f2 内部变量，会受到后续的操作影响
	defer func() {
		fmt.Println("defer 1 :", n)
	}()
	// 这里的n传值时是1 ，GO的传值时复制操作，这时函数内的参数n已经定义为1，后续的外部变量更改不会影响
	defer func(n int) {
		fmt.Println("defer 2 :", n)
	}(n)
	n = 2

	// defer 2 : 1
	// defer 1 : 2
}

// 如果defer 的返回值已经有定义
// defer 执行过程中的x是f3内部的x 也是在函数定义时已经定义好的返回值 变量x
// 在RET返回指令前，defer 对 返回值x 进行了值修改，导致返回的值出现变化
func f3() (x int) {
	x = 10
	defer func() {
		x++
	}()
	return
	// 11
}

// 如果返回值还没定义
// 同上，区别在于返回值在函数定义时没有定义
// 导致return x 的时候，返回值是会创建新的变量，就是返回值是复制x的值，但不是x
// 后续defer 对 x进行了操作，但不会影响RET返回指令的返回内容
func f4() int {
	x := 10
	defer func() {
		x++
	}()
	return x
	// 10
}

// 如果返回值变量名与defer不一样
func f5() (y int) {
	x := 10
	defer func() {
		x++
	}()
	return x
}

// defer 命令组合
func f6(number, a, b int) int {
	ret := a + b
	fmt.Println("step :", number, "a:", a, "b:", b, "ret:", ret)
	return ret
}

func f7() {
	a := 1
	b := 2
	defer f6(1, a, f6(2, a, b))
	a = 0
	defer f6(3, a, f6(4, a, b))
	b = 0
}

func main() {
	f7()
}
