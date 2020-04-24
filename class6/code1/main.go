package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// 标准函数形式

// func name(parameter-list) (result-list) {
//     body
// }

// 函数的类型被称为函数的标识符。如果两个函数形式参数列表和返回值列表中的变量类型一一对应，
// 那么这两个函数被认为有相同的类型或标识符。
// 形参和返回值的变量名不影响函数标识符，也不影响它们是否可以以省略参数类型的形式表示。

// 每一次函数调用都必须按照声明顺序为所有参数提供实参（参数值），没有任何方法可以通过参数名指定形参，形参和返回值的变量名对于函数调用者而言没有意义
// 实参通过值的方式传递，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。但是，如果实参包括引用类型，如指针，slice(切片)、map、function、channel等类型，实参可能会由于函数的间接引用被修改。

// 关于作用域部分不作描述
// 除了常规的 arg type 方式赋值，go里面的赋值可以将相同类型的放在一起。GO语言里面，没有默认参数值
// 如果需要动态接收参数类型，可以使用 ...type形式。
func f1(a, b int, args ...interface{}) {
	fmt.Printf("args type :%T\n", args)
	fmt.Printf("args value :%v\n", args)
}

// 关于返回值
// 1.返回多个值的时候，需要通过括号括起来
func f2() (string, int) {
	return "f2", 10
}

// 如果只有一个返回值时候可以忽略
func f3() string {
	return "f2"
}

// 返回的值可以定义名称，这时候这个名称在函数内部已经完成定义并可以使用
// 1. return 可以不需要写对应的变量名
// 2. 如果过程中没有对返回的变量名进行赋值，返回的都会类型对应的零值
func f4(k bool) (name string, age int) {
	if k {
		name = "f4"
		age = 10
		return
	}
	return
}

// 递归函数 可以在函数内部调用自身，形成递归
func f5(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f5(x - 1)
}

// 重点
// 匿名函数
// go 语言里面，func 是一个第一类值，就行相当于普通的int,map，可以被赋值给变量
// go 语言的func 内部不能创建新的func，但是可以通过上面提到的特性，在函数内部将新的函数赋值给变量，形成在函数内部创建函数
// 由于这种方式创建的函数没有对应的名字，一般称为匿名函数
func nameless() func(int, string) {
	// func inner (){}  // illegal
	sum := 0
	sumFunc := func(a int, content string) {
		localTotal := 10
		sum += a
		localTotal += a
		fmt.Printf("%-20s: global : %v, local: %v\n", content, sum, localTotal)
	}
	sumFunc(2, "func in nameless")
	return sumFunc
}

func namelessTest() {
	f1 := nameless()
	f2 := nameless()
	b := 2
	f1(b, "f1 first time ")
	f1(b, "f1 second time ")
	f2(b, "f2 first time ")
}

// 错误接收及一般处理策略
// GO语言里面没有像其他语言一样的异常处理机制，比如python 里面的try except。
// 所有的已知错误都需要通过error进行传递，发生未知错误时使用panic
func f6() {
	// 比如在打开文件时候，会返回error
	f, err := os.Open("error.txt")
	if err != nil {
		// 错误类型可以直接打印
		fmt.Println(err)
		// fmt.Printf("%v", err)
	}
	defer f.Close()
}

// 错误处理方式一  上报且添加信息
// 传入文件名，打开并读取文件内容
func f7(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		// 错误时，返回nil ，并用fmt.Errorf 拼接新的错误内容
		return nil, fmt.Errorf("Open %s Failed : %v", filename, err)
	}
	defer f.Close()
	b, _ := ioutil.ReadAll(f)
	return b, nil
}

// 错误处理方式二  重试
func f8(filename string) ([]byte, error) {
	// 重复次数
	const t = 3
	for tries := 0; tries < t; tries++ {
		b, err := f7(filename)
		if err == nil {
			return b, nil
		}
		fmt.Println(err, "retry...")
	}
	// 根据时间重试
	const timeout = 3 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		b, err := f7(filename)
		if err == nil {
			return b, nil
		}
		fmt.Println(err, "retry by time")
		time.Sleep(time.Second)
	}
	return nil, fmt.Errorf("Open Failed in %d times", t)
}

func main() {
	// f1(1, 2, 3, 4, 5, 6)

	// fmt.Println(f4(true))
	// fmt.Println(f4(false))

	// _, err := f8("error.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	namelessTest()
}
