# MAKE 和 NEW 函数
## MAKE 函数
### 作用
make 函数用于对这种 结构体 的初始化  
make 函数对 slice map chan 类型的进行初始化并返回对应的值。 map chan 类型 以后章节也会用到，这里不作演示
### 语法
``` go
make(type,len,cap)
```
type 为数据类型，len为长度，cap为容量
### 完整例子
``` go
func f9() {
	// s 变量定义为一个 []int的切片，但是还没赋值内容
	var s []int
	// 只能定义为一个长度3，容量3的切片
	s = []int{1, 2, 3}
	f7(s, "通过固定内容赋值")
	// 自定义一个切片的长度，容量（但是内容都为零值）
	s1 := make([]int, 3, 10)
	// 等价于
	// var s1 []int
	// s1 = make([]int,0,10)
	f7(s1, "make初始化")
}
```
## NEW 函数
### 作用
new 函数同样用于初始化，但内容不限制类型，返回的内容为指针类型  
new 是对传入的类型分配空间并返回对应的内存地址指针，内容为对应类型的零值
### 语法
``` go 
new(Type)
```
### 例子
``` go
func f10() {
	n1 := new(int)
	fmt.Printf("%p", n1)
	fmt.Println()
	// 等价于
	var v int
	n2 := &v
	fmt.Printf("%p", n2)
}
``` 