package types

// 接口不能定义方法
// func (self Person)Sleep(){}

// 接口不能循环嵌套或者引用自身
type I1 interface {
	// I1  // 无法使用
}

type I2 interface {
	// I3  // 不能嵌套
}

type I3 interface {
	I2
}
