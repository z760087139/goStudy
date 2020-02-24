package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// 内容：字符串

// https://books.studygolang.com/gopl-zh/ch3/ch3-05.html

// 字符串是一个不可改变的字节序列
// 字符串必须是用双引号或者是反引号扩起
// 文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列
// 字符串可以用s[n]进行索引

// 关于字符串长度判断（中文、英文）
func f1() {
	s1 := "string"
	s2 := "中文"
	fmt.Println(len(s1))
	fmt.Printf("%T,%v", s2, len(s2))
	// 第i个字节并不一定是字符串的第i个字符，因为对于非ASCII字符的UTF8编码会要两个或多个字节
	fmt.Println(s2[2])
}

// 修改字符串
// 通过切片拼接方式或者将string 转成[]byte形式
func f2() {
	s1 := "string"
	// s1[0] = "S" //无法执行，会报错
	s2 := "S" + s1[1:]
	fmt.Println(s2)
}

// 遍历字符串的问题
// go 语言默认编码是utf8，字符串是由ascii和rune组成（非ascii能表示的，以utf-8存储的内容）
// rune 组成的字节数根据实际情况存在不同的长度
// 但是字符串len(s)表示的长度为字节长度
// go 语言里面 for range 循环会对rune 进行隐式读取，不用担心循环读取问题
func f3() {
	s1 := "hello,世界"
	for i, s := range s1 {
		fmt.Printf("%d\t%q\t%d\n", i, s, s)
	}
}

// 统计字符储量
func f4() {
	s1 := "hello,世界"
	num := utf8.RuneCountInString(s1)
	fmt.Println(num)
}

// 常用的字符串处理内建库
// strings bytes strconv unidcode path/filepath
// 一般的replace index  join contain count 等方法 包含在bytes 和strings 函数中
func f5() {
	a := "hello,世界"
	fmt.Println(strings.Contains(a, "世界"))
}

// 字符串/数字转换 strconv
func f6() {
	// 为什么需要用strconv函数？
	// 通过string()直接转换函数的时候，会把数字识别成码点进行翻译
	fmt.Println(string(89))

	// 数字转字符串
	s := strconv.FormatInt(89, 10)
	fmt.Printf("数字转字符串 %T,%v\n", s, s)
	// 字符串转数字
	i, _ := strconv.ParseInt(s, 10, 64)
	fmt.Printf("字符串转数字 %T,%v\n", i, i)

}

// 关于常量
// 常量表达式的值在编译期计算，而不是在运行期。每种常量的潜在类型都是基础类型：boolean、string或数字。
// 常量的值不可修改，
const (
	pi = 3.14159
	// 量的声明也可以包含一个类型和一个值
	timeout               = time.Second
	time2   time.Duration = 2
	// 如果和上面的变量相同时，可以忽略不写
	time3
	// iota ，在第一行常量声明的地方，iota置为0，后面的每行常量声明时+1
	i0 = 1 << iota // 1 << 4
	i1             // 1 << 5
	i2             // 1 << 6
	i3             // 1 << 7
	i4             // i << 8
	i5             // i << 9
)

func f7() {
	fmt.Printf("%b\n", i0)
	fmt.Printf("%b\n", i1)
	fmt.Printf("%b\n", i2)
	fmt.Printf("%b\n", i3)
	fmt.Printf("%b\n", i4)
	fmt.Printf("%b\n", i5)
}

func main() {
	f7()
}
