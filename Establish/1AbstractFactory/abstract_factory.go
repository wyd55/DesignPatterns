package _AbstractFactory

import "fmt"

// 适用场景:
// 在以下情况下可以考虑使用抽象工厂模式：
//
// (1) 一个系统不应当依赖于产品类实例如何被创建、组合和表达的细节，这对于所有类型的工厂模式都是很重要的，用户无须关心对象的创建过程，将对象的创建和使用解耦。
//
// (2) 系统中有多于一个的产品族，而每次只使用其中某一产品族。可以通过配置文件等方式来使得用户可以动态改变产品族，也可以很方便地增加新的产品族。
//
// (3) 属于同一个产品族的产品将在一起使用，这一约束必须在系统的设计中体现出来。同一个产品族中的产品可以是没有任何关系的对象，但是它们都具有一些共同的约束，如同一操作系统下的按钮和文本框，按钮与文本框之间没有直接关系，但它们都是属于某一操作系统的，此时具有一个共同的约束条件：操作系统的类型。
//
// (4) 产品等级结构稳定，设计完成之后，不会向系统中增加新的产品等级结构或者删除已有的产品等级结构。
func NewSimpleLunchFactory() LunchFactory {
	return &SimpleLunchFactory{}
}

// AbstractProduct（抽象产品）：它为每种产品声明接口，在抽象产品中声明了产品所具有的业务方法。
type Lunch interface {
	Cook()
}

// ConcreteProduct（具体产品）：它定义具体工厂生产的具体产品对象，实现抽象产品接口中声明的业务方法。
type Rise struct {
}
type Tomato struct {
}

func (r *Rise) Cook() {
	fmt.Println("it is rise.")
}
func (t *Tomato) Cook() {
	fmt.Println("it is Tomato.")
}

// AbstractFactory（抽象工厂）：它声明了一组用于创建一族产品的方法，每一个方法对应一种产品。
type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

// ConcreteFactory（具体工厂）：它实现了在抽象工厂中声明的创建产品的方法，生成一组具体产品，这些产品构成了一个产品族，每一个产品都位于某个产品等级结构中。
type SimpleLunchFactory struct {
}

func (s *SimpleLunchFactory) CreateFood() Lunch {
	return &Rise{}
}
func (s *SimpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}
