package _AbstractFactory

func NewAbstractFactory() Factory {
	return &SpecificFactory{}
}

// 抽象工厂
type Factory interface {
	// 创建抽象产品
	ChanPing1() Product
	ChanPing2() Product
}

// 抽象产品
type Product interface {
	Cook()
}

// 具体工厂
type SpecificFactory struct {
}

func (*SpecificFactory) ChanPing1() Product {
	return &ChanPing1{}
}
func (*SpecificFactory) ChanPing2() Product {
	return &ChanPing2{}
}

// 具体产品
type ChanPing1 struct {
}

func (*ChanPing1) Cook() {

}

type ChanPing2 struct {
}

func (*ChanPing2) Cook() {

}
