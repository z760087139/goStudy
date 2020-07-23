package types

import "fmt"

// 举个例子，比如所有的Person 类型都可以执行Run 方法
type Person struct {
	name string
	age  int
}

// 方法定义跟函数相似，只是在前面先注释这个方法属于哪个类型
// (self Person) 这个用于说明这个方法属于 Person 类型，self 是Person 类型的简称，用户函数内部表示Person 类型自身
// self 的命名不固定，一般常用的有self 或者 是类型的第一个字母 (比如Person 用 p)表示
func (self Person) Run() {
	fmt.Println(self.name, "is running")
}

// 如果需要修改类型中的属性值，需要将方法中的引用对象从值改成指针形式
func (self *Person) Rename(name string) {
	self.name = name
}

// 错误例子
func (self Person) RenameFailed(name string) {
	self.name = name
}

func (self Person) Name() string {
	return self.name
}

// 关于继承
// 子级在引用父级类型的同时，会继承了父级的方法
type Student struct {
	Person
	Class int
}

func NewStudent() *Student {
	student := &Student{
		Person{
			name: "jszc",
			age:  2,
		},
		1,
	}
	return student
}

// 继承并增加内容，重写看下一个例子
func (self Student) Run() {
	self.Person.Run() // 相当于 super.Run()
	fmt.Println("this is Student method")
}
