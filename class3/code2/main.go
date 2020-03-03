package main

import (
	"fmt"
	"math"
)

// 内容：数字类型
// int int8 int16 int32 int64 uint uint32 uint64...
// float float32 float64

// 其中rune 其实等价于 int32

// 其中有符号整数采用2的补码形式表示，也就是最高bit位用来表示符号位，
// 一个n-bit的有符号数的值域是从$-2^{n-1}$到$2^{n-1}-1$。无符号整数的所有bit位都用于表示非负数，值域是0到$2^n-1$。
// 例如，int8类型整数的值域是从-128到127，而uint8类型整数的值域是从0到255。

// 下面是Go语言中关于算术运算、逻辑运算和比较运算的二元运算符，它们按照优先级递减的顺序排列：

// *      /      %      <<       >>     &       &^
// +      -      |      ^
// ==     !=     <      <=       >      >=
// &&
// ||

// %取模运算符的符号和被取模数的符号总是一致的，因此-5%3和-5%-3结果都是-2
// 除法运算符/的行为则依赖于操作数是否为全为整数，比如5.0/4.0的结果是1.25，但是5/4的结果是1，因为整数除法会向着0方向截断余数。

// 整数运算例子
func f1() {
	var a int = 5
	var b int64 = 4
	// fmt.Println(a / b)  类型不一致，无法执行
	fmt.Println(a / int(b))
	var c float32 = 3
	var d float32 = 2
	fmt.Println(c / d)
}

// 位运算
func f2() {
	var a int8 = 1 << 6
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b", a<<1)

}

// go 语言里面 float 使用标准的IEEE754表示浮点数 会存在不准确的情况
// float32 小数位大约是6位精度 float64 大约是15位进度
// float 转int 的时候，优先往0方向靠拢
func f3() {
	var f32 = float32(0)
	var f64 = float64(0)
	for i := 0; i < 10; i++ {
		f32 += 0.1
		f64 += 0.1
	}
	fmt.Println(f32, f64)
	fmt.Println(int(f32), int(f64))
	// 打印精度
	fmt.Printf("%1.3f\n", f64)
	// 截取精度
	fmt.Println(f4(float64(123.4567), 3))
	fmt.Println(f4(float64(123.4567), 1))
}

// 自己编写截取函数
func f4(source float64, n int) float64 {
	truncNum := math.Pow10(n)
	return math.Trunc((source+0.5/truncNum)*truncNum) / truncNum
}

// 关于算数运算精度，go 语言官方没有demical类型。tidb实现了相关数据类型
// https://github.com/pingcap/tidb/blob/0dc91069d8b94300a449837bcd6d001af9117120/util/codec/decimal.go

// 复数
// 因为这部分预计不常用，只作简单介绍
// 复数定义
func f5() {
	var x complex128 = complex(1, 2) // 1 +2i
	var y complex128 = complex(3, 4) // 3 + 4i
	fmt.Println(x+y, x*y)
	fmt.Println(real(y))
	fmt.Println(imag(y))
}

// 布尔值
// go 语言的布尔值只有true 和 false。0,1 不能当做布尔值使用


func main() {
	f5()
}
