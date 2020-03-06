# MAP
## 简介
在Go语言中，一个map就是一个哈希表的引用  
map类型可以写为map[K]V，其中K和V分别对应key和value。map中所有的key都有相同的类型，所有的value也有着相同的类型，但是key和value之间可以是不同的数据类型。  
其中K对应的key必须是支持==比较运算符的数据类型

## golang原文
The key can be of any type for which the equality operator is defined, such as integers, floating point and complex numbers, strings, pointers, interfaces (as long as the dynamic type supports equality), structs and arrays  
 Like slices, maps hold references to an underlying data structure.  If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller.

 ### 简单翻译
 key的值可以是所有支持== 运算符的数据类型，比如   
 int,*float,complex,string,pointer,interface,struct,array。（但是不支持slice）  
和slice相似，map是指向一个底层数据结构的，在函数传参时候，对map内容的修改只是在caller可见

## MAP使用
### 定义及赋值
在上一次课提到的make函数可以用于创建map结构，因为map需要指向一个底层数据结构，初始化时候需要使用map函数，否则单纯定义后无法使用
#### 定义语法
map[K]V  
#### 例子
``` go
// 定义
var m1 map[int]string
// 赋值 
m1 = map[int]string{}
m1[0] = "A"
// 方式二
m2 := make(map[string]int)
// 方式三
m3 := map[int]string{}
```

### 寻值
如果寻找的key 不存在，会根据value 的数据类型返回零值。
```go
m2 := make(map[string]int)
fmt.Println(m2["A"])
```
因此，可以把map当做set使用
```go 
set := make(map[string]bool)
set["created"] =true
fmt.Println(set["created"],set["none"])
```
同时，在寻值时候可以接受一个参数，判断是否存在该值（因为无论是否存在都有返回值）
```
if _,ok := set["none"];!ok{
    fmt.Println("none not exists")
}
```
### 扩展内容
map可以存储slice、struct等内容。不是所有value 类型都能直接修改里面的内容。  
创建一个map[string][]int的类型，并直接修改内容
``` go
m := make(map[string][]int)
m["m"] = []int{1,2,3}
m["m"][0] = 0
fmt.Println(m)
```
创建一个map[string]Person类型 Person是一个存放name string的结构体  
```go
type Person struct{
	name string
}

func f9(){
	m := make(map[string]Person)
	m["m"] = Person{
		name : "m",
	}
	// m["m"].name = "M"
	fmt.Println(m["m"].name)
}
```
[问题解答](https://stackoverflow.com/questions/17438253/accessing-struct-fields-inside-a-map-value-without-copying)

### 作为函数传参时
如果函数的参数需要传入map，并且对map内容进行了修改
``` go
func f5(){
	m := make(map[string]Person)
	m["M"] = Person{
		name :"M",
	}
	fmt.Println("f5 before f6",m)
	f6(m)
	fmt.Println("f5 after f6",m)
	m2 := make(map[string][]int)
	m2["M"] = []int{1,2,3}
	// fmt.Println(len(m2["M"]),cap(m2["M"]))
	f7(m2)
	fmt.Println(m2)
}

func f6(m map[string]Person){
	p := m["M"]
	p.name = "m"
	fmt.Println("f6 change1",m)
	m["M"] = p
	fmt.Println("f6 chagne2",m)
}

func f7(m map[string][]int){
	d := m["M"]
	// 修改d[0]的值是否对m产生影响
	d[0] = 0 
	fmt.Println(m)
	// 增加d的值是否对m产生影响
	d = append(d,4)
	fmt.Println(m)
}
```
## 完整例子
```go
package main

import (
	"fmt"
)

func f1(){
	var m1 map[int]string
	m1 = map[int]string{}
	m1[0] = "A"
	m2 := make(map[string]int)
	m3 := map[int]string{}
	fmt.Println(m1,m2,m3)
	fmt.Println(m2["A"])
}
func f2(){
	set := make(map[string]bool)
	set["created"] =true
	fmt.Println(set["created"],set["none"])
	if _,ok := set["none"];!ok{
		fmt.Println("none not exists")
	}
}

func f3(){
	m := make(map[string]int,0)
	fmt.Printf("first :%p\n",m)
	m["A"] = 1
	fmt.Printf("second :%p\n",m)
	f4(m)
	fmt.Printf("fifth :%p\n",m)
	fmt.Println(m)
}

func f4(m map[string]int){
	fmt.Printf("third %p \n",m)
	m["A"] = 10
	fmt.Println(m["A"])
	m["B"] = 20
	fmt.Printf("fourth %p\n",m)
}

type Person struct{
	name string
}

func f5(){
	m := make(map[string]Person)
	m["M"] = Person{
		name :"M",
	}
	fmt.Println("f5 before f6",m)
	f6(m)
	fmt.Println("f5 after f6",m)
	m2 := make(map[string][]int)
	m2["M"] = []int{1,2,3}
	// fmt.Println(len(m2["M"]),cap(m2["M"]))
	f7(m2)
	fmt.Println(m2)
}

func f6(m map[string]Person){
	p := m["M"]
	p.name = "m"
	fmt.Println("f6 change1",m)
	m["M"] = p
	fmt.Println("f6 chagne2",m)
}

func f7(m map[string][]int){
	d := m["M"]
	// 修改d[0]的值是否对m产生影响
	d[0] = 0 
	fmt.Println(m)
	// 增加d的值是否对m产生影响
	d = append(d,4)
	fmt.Println(m)
}

func f8(){
	m := make(map[string][]int)
	m["m"] = []int{1,2,3}
	m["m"][0] = 0
	fmt.Println(m)
}

func f9(){
	m := make(map[string]Person)
	m["m"] = Person{
		name : "m",
	}
	// m["m"].name = "M"
	fmt.Println(m["m"].name)
}

func main(){
	f5()
}

```


