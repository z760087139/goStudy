package main

import "fmt"

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

func main() {
	// f1(1, 2, 3, 4, 5, 6)
	fmt.Println(f4(true))
	fmt.Println(f4(false))
}
