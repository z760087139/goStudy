package mock

import (
	"testing"
)

// struct -- mock
type person struct {
	// func field stub  可以通过定义 login 的方法对 Login 方法的内容进行重构
	login func(id int) (bool, error)
	// func field stub
	run func()
}

func (p person) Login(id int) (bool, error) {
	return p.login(id)
}

func (p person) Run() {
	p.run()
}

func Test_PersonCheck(t *testing.T) {

}
