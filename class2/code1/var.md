## 声明及指针基础

### 声明
GO语言主要有四种类型声明语句
| name  | 用途     |
| ----- | -------- |
| var   | 定义变量 |
| const | 定义常量 |
| type  | 定义类型 |
| func  | 定义函数 |

#### 关键字
``` 
break      default       func     interface   select
case       defer         go       map         struct
chan       else          goto     package     switch
const      fallthrough   if       range       type
continue   for           import   return      var
```
#### 内建常量
``` true    false   iota    nil ```
#### 内建类型（底层类型）
```
int      int8      int16      int32      int64
uint     uint8     uint16     uint32     uint64    uintptr
float32  float64   complex128 complex64
bool     byte      rune       string     error
```
#### 内建函数
```
make    len       cap      new      copy  
complex real      imag     append
panic   recover   close    delete
```

#### 声明语法
``` go
var  name Type = ...
```

#### 声明例子

``` go
// 声明一个变量
var a1 string
// 声明多个变量
var (
	b1 string
	b2 string
	// 多个重复类型的变量
	c1, c2, c3 string = "c1", "c2", "c3"
	// 声明推导
	d1, d2 = 10, "string"
)
```

#### 其他声明方式
- 简短变量声明
``` go
// 该方式只能在非第一级使用
    e1, e2 := 10, 20
```
等价于
``` go
var e3, e4 = 30, 40
``` 
简短变量声明中，必须至少有一个是新的变量声明，如果已经有的变量则为赋值操作
``` go
e1, e5 := 11, 50
```

#### 注意事项
1.  除了全局变量以外，已经声明的变量都必须被使用，否则无法编译
2.  如果函数返回的内容不需要使用，可以用 _ 接收
3.  go 语言中 字符串表示只能使用 "" 扩起，详细内容在后面章节讲解

#### 完整例子
``` go
// 声明一个变量
var a1 string

// 声明多个变量
var (
	b1 string
	b2 string
	// 多个重复类型的变量
	c1, c2, c3 string = "c1", "c2", "c3"
	// 声明推导
	d1, d2 = 10, "string"
)

// 声明例子
func f1() {
	// 打印类型和变量值，关于fmt的打印格式会在后面章节具体讲解，这里不作详细描述
	fmt.Printf("type:%T,value:%v", d1, d1)
	// 如果是一些需要进行传递的变量类型，个人不建议用声明推导方式创建
	// 因为在传递的时候有部分函数会要求指定的变量类型，int 不能传递给int64 int32

	// 声明方式二
	// 简短变量声明
	// 该方式只能在非第一级使用
	e1, e2 := 10, 20
	// 等价于
	var e3, e4 = 30, 40

	// 简短变量声明中，必须至少有一个是新的变量声明，如果已经有的变量则为赋值操作
	e1, e5 := 11, 50
	// 可以尝试注释下面一句之后运行程序
	fmt.Println(e1, e2, e3, e4, e5)

}
```

### 赋值操作
同时进行多个数值的赋值操作。赋值语句会先对右边的表达式进行求值后统一更新左边对应的变量值
``` go
x, y = y, x
```
**GO 语言里面，所有的赋值操作，其实都是copy，copy的内容包括指针**  
不同的数据类型，在copy的过程中，实际上的内容也不一样，现在只作了解即可，会在每个数据类型介绍时候提及
| type    | syntactic sugar for                                      |
| ------- | -------------------------------------------------------- |
| array   | the array                                                |
| string  | struct holding len + a pointer to the backing array      |
| slice   | struct holding len, cap + a pointer to the backing array |
| map     | pointer to a struct                                      |
| channel | pointer to a struct                                      |

#### 举个例子
``` go
s := [5]int{1,2,3,4,5}
for _, v := range s {
    fmt.Printf("%p\t", &v)
}
fmt.Println()
for i:=0;i<len(s);i++{
    fmt.Printf("%p\t",&s[i])
}
``` 