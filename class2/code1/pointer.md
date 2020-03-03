## 指针
**GO 语言里面虽然可以获取变量的指针地址，但是GO 语言不支持指针地址操作**
### 指针的常用命令
| 命令  | 含义                        |
| ----- | --------------------------- |
| *Type | 表示T类型的指针               |
| &V    | 表示V变量的指针             |
| *P    | 反引用，表示指针p指向的内容 |

### 例子
#### 定义一个变量，为某个类型的指针
``` go
var p1 *int
```
定义变量 p 的数据类型是int类型的指针  
定义后并不会为p1赋值  
可以尝试打印当前的p1内容
``` go
fmt.Println(*p1)
```
为p1 赋值
``` go
d1 := 1
p1 = &d1
```
不能直接使用&1,因为没有为1进行内存地址分配（会报错，可尝试）
尝试打印当前的p1和p1的内容  
#### 指针操作
再强调一遍，**go语言不支持指针地址操作**  
只能对指针所指向内容进行修改
``` go
// p1 ++ 无法执行
*p1++
```
p1作为一个变量，也有对应的内存地址，同样可以打印
``` go
fmt.Println(&p1)
```




