# 结构体
### 简介
结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。  
每个值称为结构体的成员。  
***成员名称遵循大写字母导出规则***
### 基本语法
``` go
type Name struct{
    ...
}
```
### 例子
``` go
type Person struct {
	name    string
	age     int
    friends []Person
    // friend  Person
    // friend *Person
}
```
一个命名为S的结构体类型将不能再包含S类型的成员,但可以使用*S 的指针引用 
注意例子中 friends 是一个 Preson类型的切片  
另外，结构体成员的顺序是有意义的，不同的顺序，会表示不同的结构体
### 定义、赋值结构体
方式一
``` go
a := new(Person)
fmt.Printf("%#v\n", a)
```
方式二
``` go
b := Person{
		name: "B",
		age:  10,
		// friends: nil,
		// friends: []Person{
		// 	Person{
		// 		name: "b",
		// 		age:  1,
		// 	},
		// },
	}
fmt.Printf("%#v\n", b)
```
### 引用、编辑结构体
继续引用上面定义的 b 进行内容引用和修改
``` go
fmt.Println(b.name)
b.name = "b"
b.friends = []Person{
	Person{
		name: "b",
		age:  1,
	},
}
fmt.Printf("%#v\n", b)
```
《GO语言圣经》中用结构体构建的[二叉树例子](https://books.studygolang.com/gopl-zh/ch4/ch4-04.html)  

### 传参时
由于在函数传参时，都是值拷贝。在函数中希望修改结构体内容，需要传递指针类型
