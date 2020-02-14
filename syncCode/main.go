package main

import (
	"fmt"
	"time"
	"sync"
	"sync/atomic"
)

// 原子操作
// 增加操作 Add 
func f1(sum *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(sum, 3)
}

func f1Worker() {
	var wg sync.WaitGroup
	var sum int64 = 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go f1(&sum, &wg)
	}
	wg.Wait()
	fmt.Println(sum)
}

// 初始化操作

func f2() {
	var once sync.Once
	var age int
	once.Do(func() {
		age = 10
	})
	fmt.Println(age)
}

// 读取 Load 覆盖 Store 交换 Swap
var config atomic.Value

func configStore(){
	for {
		// 触发更新config的条件
		time.Sleep(time.Second)
		config.Store("some configuration")
	}
	
}

func f3(){
	// 创建一个更新config的后台进程
	go configStore()
	// 获取当前config
	go func(){
		myConfig := config.Load()
		fmt.Println(myConfig)
	}
}

// 重点内容：
// https://github.com/chai2010/advanced-go-programming-book/blob/master/ch1-basic/ch1-05-mem.md
// 在Go语言中，同一个Goroutine线程内部，顺序一致性内存模型是得到保证的。但是不同的Goroutine之间，并不满足顺序一致性内存模型，
// 需要通过明确定义的同步事件来作为同步的参考。如果两个事件不可排序，那么就说这两个事件是并发的。
// 为了最大化并行，Go语言的编译器和处理器在不影响上述规定的前提下可能会对执行语句重新排序（CPU也会对一些指令进行乱序执行）。

func main() {
	f2()
}
