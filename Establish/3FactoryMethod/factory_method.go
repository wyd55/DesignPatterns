package _FactoryMethod

import "fmt"

// Product 是产品接口
type Product interface {
	Use() string
}

// ConcreteProduct 是具体产品
type ConcreteProduct struct{}

// Use 实现具体产品的 Use 方法
func (p *ConcreteProduct) Use() string {
	return "使用具体产品"
}

// Creator 是创建者接口
type Creator interface {
	CreateProduct() Product
}

// ConcreteCreator 是具体创建者
type ConcreteCreator struct{}

// CreateProduct 实现具体创建者的 CreateProduct 方法
func (c *ConcreteCreator) CreateProduct() Product {
	return &ConcreteProduct{}
}

func FactoryMethodMain() {
	// 创建具体创建者
	creator := &ConcreteCreator{}
	// 使用具体创建者创建产品
	product := creator.CreateProduct()
	// 使用产品
	result := product.Use()
	fmt.Println(result)
}
