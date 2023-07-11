package _Builder

import "fmt"

/*
	参与者：
	Product(产品)(具体)：只有这一种对象，只是这个对象的内部构成由不同的具体构造者(ConcreteBuilder)来实现
	ConcreteBuilder(具体构造者)(具体)：用来构造不同结构产品的构造者
	Builder(构造器)(抽象)：具体构造者实现的接口
	Director(导演)(具体)：负责指导整个构建过程，知道构建的步骤和顺序，客户端通过他来获取一个产品

	创建对象的特点：
	只能创建一种对象，但是对象结构可以不同

	具体流程：
	1、创建一个导演和一个构造器
	2、把构造器赋值给导演
	3、然后让导演来构造需要的产品，以为导演知道构造流程，构造者知道构造方式
	4、最后通过构造器返回构造好的产品
*/

// Product 产品
type Product struct {
	partA string
	partB string
	partC string
}

// Builder 构建器接口
type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetProduct() *Product
}

// ConcreteBuilder 具体构建器
type ConcreteBuilder struct {
	product *Product
}

// NewConcreteBuilder 创建一个具体构建器实例
func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{
		product: &Product{},
	}
}

// BuildPartA 构建产品的部分 A
func (b *ConcreteBuilder) BuildPartA() {
	b.product.partA = "Part A"
}

// BuildPartB 构建产品的部分 B
func (b *ConcreteBuilder) BuildPartB() {
	b.product.partB = "Part B"
}

// BuildPartC 构建产品的部分 C
func (b *ConcreteBuilder) BuildPartC() {
	b.product.partC = "Part C"
}

// GetProduct 获取构建完成的产品
func (b *ConcreteBuilder) GetProduct() *Product {
	return b.product
}

// Director 导演者
type Director struct {
	builder Builder
}

// Construct 构造产品的过程
func (d *Director) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuildPartC()
}

func BuilderMain() {
	// 创建具体构建器
	builder := NewConcreteBuilder()

	// 创建导演者并指定构建器
	director := &Director{
		builder: builder,
	}

	// 构造产品
	director.Construct()

	// 获取构建完成的产品
	product := builder.GetProduct()

	// 打印产品的各个部分
	fmt.Println("Part A:", product.partA)
	fmt.Println("Part B:", product.partB)
	fmt.Println("Part C:", product.partC)
}
