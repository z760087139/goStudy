package main

import "fmt"

type Person struct {
	name    string
	age     int
	friends []Person
	// friend  Person
}

// 演示Person的定义
func f1() {
	a := new(Person)
	fmt.Printf("%#v\n", a)
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

	// 引用结构体的内容
	fmt.Printf("%#v\n", b)
	fmt.Println(b.name)

	// 修改结构体内容
	b.name = "b"
	// 按顺序初始化 、 用名字和值初始化 两种方式
	b.friends = []Person{
		Person{
			name: "b",
			age:  1,
		},
		Person{
			"S",
			2,
			nil,
		},
	}
	fmt.Printf("%#v\n", b)

}

func f2(p *Person) {
	p.name = "f2"
	fmt.Println("p from f4", p)
}

// 传参例子
func f3(p Person) {
	p.name = "f3"
	fmt.Println("p from f3", p)
}

// 传参例子
func f4() {
	p := Person{
		name: "f4",
		age:  10,
	}
	f3(p)
	fmt.Println("p from f4", p)
	f2(&p)
	// f2(P)
	fmt.Println("p from f4", p)

}

// 关于嵌入和匿名成员
type Student struct {
	// 嵌入  匿名成员
	Person
	Class string
}

func f5() {
	// 初始化嵌套的结构体成员
	//  方式一
	var s1 Student
	s1.name = "s1"
	s1.age = 10
	s1.Class = "A1"

	// 通过语法糖，对没有重复名的成员在访问或者赋值时可以直接电泳

	// 方式二
	// 用名字初始化的时候，不能使用简写的语法糖，可以使用循序赋值
	s2 := Student{
		Person: Person{
			name: "s2",
			age:  10,
		},
		Class: "A1",
	}

	s3 := Student{
		Person{
			name: "s3",
			age:  10,
		},
		"A1",
	}

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

func main() {
	f4()
}
