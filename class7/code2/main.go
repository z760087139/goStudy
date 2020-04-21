package main

import (
	"fmt"
	"runtime"
)

// 虽然Go的panic机制类似于其他语言的异常，但panic的适用场景有一些不同。
// 由于panic会引起程序的崩溃，因此panic一般用于严重错误
func main() {
	// recover
	defer func() {
		if r := recover(); r != nil {
			// print track back
			stack := make([]byte, 4096)
			n := runtime.Stack(stack[:], false)
			fmt.Println("message from stack :")
			fmt.Println(string(stack[:n]))

			// print recover
			fmt.Println("recover :", r)
		}
	}()
	f(3)

}

// 递归函数  当x =0 时候panic
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
