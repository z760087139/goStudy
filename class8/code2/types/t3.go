package types

import "fmt"

// 类型的指针方法能被直接调用，是由于GO的隐式转换。在接口中不能使用隐式转换
type Man struct {
	name string
}

func (self *Man) Eat(food string) {
	fmt.Printf("Man is eatting %s\n", food)
}

func (self *Man) Talk() {
	fmt.Println("Man is talking")
}
