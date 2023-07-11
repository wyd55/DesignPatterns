package _Builder

import "fmt"

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
