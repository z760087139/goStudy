package pk1

import (
	"fmt"
)

// 包里的所有文件都可以写初始化函数，但是执行顺序是按照文件名排序，不建议大量使用
// 程序运行时就执行，不需要调用
func init() {
	fmt.Println("pk1 init")
}

// 由于需要将MyPrint函数提供给pk1包以外的地方使用，需要用大写字母开头命名
// 在GO 语言里面，所有用大写字母开头命名的一级变量，都是可以被外部使用的意思
// 小写字母无法被外部引用

func MyPrint() {
	fmt.Println("Message from pk1")
	// 相同的包之间互相引用不需要大写( printString()的调用 )
	printString()
}
