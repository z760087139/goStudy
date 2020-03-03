## slice内容介绍
在go语言里面，数组在实际使用过程中，由于固定长度的原因，并没有切片使用来得广泛  
切片内容包括指针、长度、容量  
- 切片实际上是通过指针引用了数组的内容，每个切都有对应的数组，存储数据内容  
- 长度是当前切片的指针个数  
- 容量是当前切片引用的数组，从切片第一个指针在数组中的位置，到数组最后一个元素为止的长度  

**引用内容**  
《GO语言圣经》中的[slice章节](https://books.studygolang.com/gopl-zh/ch4/ch4-02.html)  
### 创建切片
```go
array[begin:end:caps]  
slice[begin:end:caps]  
```
切片内容从begin开始，包括begin，到end结束，不包括end
和python 不同，切片的最后一个并不是步长，是表示新创建的slice的容量  
不填写begin,end的话，默认是第一个元素和len(s)位置元素（不是caps），具体的演示看下面的例子  
  
定义两个函数用于打印数组和切片中的值和指针（f6,f7 打印方式可以先不做研究）  
定义一个数组作为源数组使用
1.  打印源数组的值和地址指针```f6(a2)```
2.  分别使用三种方式进行切片
    1.  有头无尾
    2.  无头有尾
    3.  有头有尾
3.  分别打印三种不同形式创建的切片长度、容量、地址指针

#### 完整例子
``` go
func f5() {
	// https://books.studygolang.com/gopl-zh/ch4/ch4-02.html
	// 创建一个长度为10，内容为对应位置的数组
	a2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 源数组内容指针
	// 打印源数组各个值的地址
	f6(a2)

	// 获取a1的完整内容，但是数据类型是一个切片
	// si1 := a2[:]
	// 打印切片内容
	// f7(si1)

	// 如果是一个片段的切片内容,有头无尾
	// si2 := a2[7:]
	// f7(si2)

	// 无头有尾，注意长度和容量
	si3 := a2[:7]
	f7(si3, "")

	// 有头有尾，注意长度和容量
	// si4 := a2[2:7]
	// f7(si4)

	// 如果循环打印切片长度以外的数据（容量以内）会报错
	// for i := 0; i < cap(si4); i++ {
	// 	fmt.Print(&si4[i], "\t")
	// }
	// fmt.Println()

}

func f6(a2 [10]int) {
	fmt.Println(strings.Repeat("-", 80))
	fmt.Print("源数组内容指针:\t")
	for i := 0; i < len(a2); i++ {
		fmt.Print(&a2[i], "\t")
	}
	fmt.Println()
	fmt.Print("数据值\t")

	for i := 0; i < len(a2); i++ {
		fmt.Print(a2[i], "\t")
	}
	fmt.Println()
}

func f7(s []int, content string) {
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println(content)
	fmt.Println("长度：", len(s))
	fmt.Println("容量: ", cap(s))
	fmt.Print("各数据指针\t")

	for i := 0; i < len(s); i++ {
		fmt.Print(&s[i], "\t")
	}
	fmt.Println()

	fmt.Print("各数据值\t")

	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], "\t")
	}
	fmt.Println()

}
```
### 切片操作
#### 读取/修改切片的值
与数组方式一致，但是不能读取长度以外的索引位置  
***go 里面的数组、切片内容是存放在连续的内存地址中的***  
这点非常重要，如果错误使用切盼会导致数据错误
```go
s[0] = 10
f7(s, "修改[0]后的切片")
```
#### 增加切片内容
切片增加长度需要使用内建函数```append(slice,...type)``` 其中追加的类型必须符合切片定义的类型  
```append``` 函数会对输入的参数进行累加，并返回结果，**传入的参数不会变更**  
所以如果想对原本的变量内容进行修改，需要接收返回值
```go
s = append(s, 11)  
// 如果换成另外一个变量接收返回值
ns := append(s, 12)   
// 打印结果
f7(s, "append一个元素后的切片")
f7(ns, "append新切片")
f6(a2)
```
当追加的内容超过剩余容量，则会自动开辟一个新的内存空间，形成一个新的数组  
容量的扩展公式跟go语言版本有关，具体计算有兴趣可以查看相关资料  
容量达到2048前为翻倍扩容，往后则每次增加512  
go语言代码路径：src/runtime/slice.go
``` go 
ns = append(s, 12, 13, 14, 15, 16)
f7(ns, "append 超过容量")
``` 

#### 完整例子
``` go
// 实际例子
func f5() {
	// https://books.studygolang.com/gopl-zh/ch4/ch4-02.html
	// 创建一个长度为10，内容为对应位置的数组
	a2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 对a1进行切片操作
	// arrary/slice[begin:end:caps] 切片内容从begin开始，包括begin，到end结束，不包括end
	// 和python 不同，切片的最后一个并不是步长，是表示新创建的slice的容量
	// 不填写begin,end的话，默认是第一个元素和len(s)位置元素（不是caps），具体的演示看下面的例子
	// 重点内容 ： go 里面的数组、切片内容是存放在连续的内存地址中的

	// 源数组内容指针
	// 打印源数组各个值的地址
	f6(a2)

	// 获取a1的完整内容，但是数据类型是一个切片
	// si1 := a2[:]
	// 打印切片内容
	// f7(si1)

	// 如果是一个片段的切片内容,有头无尾
	// si2 := a2[7:]
	// f7(si2)

	// 无头有尾，注意长度和容量
	si3 := a2[:7]
	f7(si3, "")
	// f7(si3[:], "")
	// 有头有尾，注意长度和容量
	// si4 := a2[2:7]
	// f7(si4)

	// 如果循环打印切片长度以外的数据（容量以内）会报错
	// for i := 0; i < cap(si4); i++ {
	// 	fmt.Print(&si4[i], "\t")
	// }
	// fmt.Println()

}

func f6(a2 [10]int) {
	fmt.Println(strings.Repeat("-", 80))
	fmt.Print("源数组内容指针:\t")
	for i := 0; i < len(a2); i++ {
		fmt.Print(&a2[i], "\t")
	}
	fmt.Println()
	fmt.Print("数据值\t")

	for i := 0; i < len(a2); i++ {
		fmt.Print(a2[i], "\t")
	}
	fmt.Println()
}

func f7(s []int, content string) {
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println(content)
	fmt.Println("长度：", len(s))
	fmt.Println("容量: ", cap(s))
	fmt.Print("各数据指针\t")

	for i := 0; i < len(s); i++ {
		fmt.Print(&s[i], "\t")
	}
	fmt.Println()

	fmt.Print("各数据值\t")

	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], "\t")
	}
	fmt.Println()

}

// 切片操作
func f8() {
	// 创建原数组
	a2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 对a2进行切片，采集一个长度小于容量的数据
	s := a2[2:6]
	// 打印当前的内存情况
	f7(s, "打印初始切片")
	// 修改切片内容
	s[0] = 10
	f7(s, "修改[0]后的切片")

	// 增加切片内容，需要使用内建函数 append
	// append 函数会对输入的参数进行累加，并返回结果，但对传入的参数不作变更
	// 所以如果想对原本的变量内容进行修改，需要接收返回值
	s = append(s, 11)
	// 注意内存地址
	f7(s, "append一个元素后的切片")
	// 打印原内容
	f6(a2)

	// 如果重新赋值一个变量，留意内存地址是否发生变化
	ns := append(s, 12)
	f7(ns, "append新切片")

	// 如果append的内容比原本的容量要多
	ns = append(s, 12, 13, 14, 15, 16)
	f7(ns, "append 超过容量")
}
```
