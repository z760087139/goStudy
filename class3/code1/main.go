package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	gbk "golang.org/x/text/encoding/simplifiedchinese"
)

// 内容：字符串

// https://books.studygolang.com/gopl-zh/ch3/ch3-05.html
// https://blog.golang.org/strings go官网博客讲解

// 表示字符串的方式一般有 string []rune []byte 三种形式

// 关于 string
// string是一个不可改变的字节序列，可以理解为一个只读的[]byte（基本等效）
// string必须是用双引号或者是反引号扩起
// 文本字符串通常被解释为采用UTF8编码的Unicode码点序列
// string可以用s[n]进行索引

// 拓展内容
// 回复上节课提到的列表能存储多长的字符串
// 实际上 go语言的字符串是由结构体组成
// https://chai2010.gitbooks.io/advanced-go-programming-book/content/ch1-basic/ch1-03-array-string-and-slice.html
// type StringHeader struct {
//     Data uintptr
//     Len  int
// }
// uintptr is an integer type that is large enough to hold the bit pattern of
// any pointer.

// 关于字符串长度判断（中文、英文）
func f1() {
	s1 := "string"
	s2 := "中文"
	fmt.Println(len(s1))
	fmt.Printf("%T,%v\n", s2, len(s2))
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
// 字符串len(s)表示的长度为字节长度 （不是字符）
// go 语言里面 for range 循环会对多字节组成的字符进行隐式读取，不用担心循环读取问题
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

// []rune
// 从string 内容可以发现，由于go语言使用utf-8编码方式
// 对于unicode码点，可能需要使用两个以上的字节存储
// 对string进行打印时，需要对string内容进行字节内容判断
// 所以go语言里面还有一种扩展的字符形式 rune

// rune 其实是一个int32的数据类型，用于存储unicode的码点

// []byte
// 用切片形式存储的字符串，由于是切片类型，所以在修改的时候比起string更方便
// 存储的内容只能是字节
func f9() {
	a := "世界"
	b := []byte("世界")
	fmt.Println(b)
	fmt.Printf("% x\n", b)
	fmt.Printf("% x\n", a)
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
	f8()
}

// 拓展内容
// 如果需要对字符串内容进行转码，需要使用额外的转码包
// 比如转成GBK
// import "golang.org/x/text/encoding"
// import "golang.org/x/text/encoding/simplifiedchinese"
// 也有另外一种方式
// https://pkg.go.dev/golang.org/x/text/transform?tab=doc

func f8() {
	gbkEncoder := gbk.GBK.NewEncoder()
	newString, _ := gbkEncoder.String("世界")
	for i, v := range newString {
		log.Printf("%d\t%q\t%d\n", i, v, v)
		log.Printf("%d\t%q\t%d\n", i, newString[i], newString[i])
	}
}
