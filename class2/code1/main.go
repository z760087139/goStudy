// 声明所属包
package main

import "fmt"
import "sync"

// 内容：声明定义及指针基础

// GO语言主要有四种类型声明语句
// var		定义变量
// const	定义常量
// type		定义类型
// func		定义函数

// 主要讲解声明及指针类型

// go语言里面，源文件第一级内容都必须是合规关键字开头

// 关键字
// break      default       func     interface   select
// case       defer         go       map         struct
// chan       else          goto     package     switch
// const      fallthrough   if       range       type
// continue   for           import   return      var

// 内建常量
// true false iota nil

// 内建类型（底层类型）
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// float32 float64 complex128 complex64
// bool byte rune string error

// 内建函数
// make len cap new append copy close delete
// complex real imag
// panic recover

// 以下都是在package main里面可以直接互相访问的全局变量定义，具体关于包的内容在后面章节讲解

// 声明语法
// var 变量名字 类型 = 表达式

// 声明一个变量
var a1 string

// 声明多个变量
var (
	b1 string
	b2 string
	e1,e2,e3 int
	// 多个重复类型的变量
	c1, c2, c3 string = "c1", "c2", "c3"
	// 声明推导
	d1, d2 = 10, "string"
)

// 声明例子
func f1() {
	// 打印类型和变量值，关于fmt的打印格式会在后面章节具体讲解，这里不作详细描述
	fmt.Printf("type:%T,value:%v\n", d1, d1)
	// 如果是一些需要进行传递的变量类型，个人不建议用声明推导方式创建
	// 因为在传递的时候有部分函数会要求指定的变量类型，int 不能传递给int64 int32

	// 声明方式二
	// 简短变量声明
	// 该方式只能在非第一级使用
	e1, e2 := 10, 20
	// 等价于
	var e3, e4 = 30, 40

	// 简短变量声明中，必须至少有一个是新的变量声明，如果已经有的变量则为赋值操作
	e1, e5 := 11, 50
	// 可以尝试注释下面一句之后运行程序
	fmt.Println(e1, e2, e3, e4, e5)

}

// 其他需要注意的地方：
// 1.除了全局变量以外，已经声明的变量都必须被使用，否则无法编译
// 2.如果函数返回的内容不需要使用，可以用 _ 接收
// 3.go 语言中 字符串表示只能使用 "" 扩起，详细内容在后面章节讲解

// 指针例子
// GO 语言里面虽然可以获取变量的指针地址，但是GO 语言不支持指针地址操作
func f2() {
	// 定义一个变量为一个类型的指针类型
	// 在类型前加 * 表示这个类型(int)的指针类型
	var p1 *int
	// fmt.Println(*p1) //打印p1存放的值会报错，p1只分配了空间，还没进行赋值
	// & 表示该变量的指针地址
	// d1的指针地址赋值给p1
	p1 = &d1
	// 打印d1的指针地址
	fmt.Println(&d1)
	// 打印p1的内容
	fmt.Println(p1)

	// 反引用
	// 通过已知的内存地址，获取内容
	// 反引用使用 *
	fmt.Println(*p1)

	// 指针操作演示
	// GO不能操作内存地址类型
	// p1++ // 这个操作会报错，无法执行
	*p1++ //该操作实际是对d1 的值（10）进行自加1，d1 改为了11
	fmt.Println(d1)
	// p1 作为一个变量，也同样是存放在某一个内存地址中，可以打印p1的内存地址
	// fmt.Println(&p1)

	// new()函数
	// 也是一种变量声明形式，等同于函数方式申明的var,但生成的是指针类型，并且对应的内存地址存放零值
	p2 := new(int)
	fmt.Println(p2)
	fmt.Println(*p2)
}

// 赋值内容
func f3() {
	// 声明并赋值一个变量
	x, y := 10, 20
	// 重新赋值
	x = 5
	// 元组赋值是另一种形式的赋值语句，同时进行多个数值的赋值操作。赋值语句会先对右边的表达式进行求值后统一更新左边对应的变量值
	x, y = y, x
	fmt.Println(x, y)
}

// GO 语言里面，所有的赋值操作，其实都是copy，copy的内容包括指针
// 不同的数据类型，在copy的过程中，实际上的内容也不一样，现在只作了解即可，在了解所有数据类型之后，会做一个专门的讲解
// type		syntactic sugar for
// array	the array
// string	struct holding len + a pointer to the backing array
// slice	struct holding len, cap + a pointer to the backing array
// map		pointer to a struct
// channel	pointer to a struct

// 关于赋值操作的例子
func f4() {
	s := [5]int{1, 2, 3, 4, 5}
	for _, v := range s {
		fmt.Printf("%p\t", &v)
	}
	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("%p\t", &s[i])
	}
}

type myInt int

func (self myInt) p(){
	fmt.Println(self)
}

func f5(){
	s := [10]myInt{1,2,3,4,5,6,7,8,9,10}
	var wg sync.WaitGroup
	for _,v := range s{
		wg.Add(1)
		go func(){
			v.p()
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	f5()
}
