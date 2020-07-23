package types

import "fmt"

// 例子
// 基本的接口类型，其中只要实现了Eat对应方法的即可
type Animal interface {
	Eat(string)
}

// Person 接口需要在实现了Animal的接口基础上（有Eat方法） 还需要有Talk方法，但是不能重复定义Run
type Person interface {
	Talk()
	// Run() illegal
	Animal
}

type Student struct {
	name string
}

func (self Student) Run() {
	fmt.Printf("student %s is running\n", self.name)
}

func (self Student) Eat(food string) {
	fmt.Printf("student %s is eatting %s\n", self.name, food)
}

func (self Student) Talk() {
	fmt.Println("student is talking")
}
