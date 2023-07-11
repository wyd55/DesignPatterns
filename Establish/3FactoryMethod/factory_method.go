package _FactoryMethod

import "fmt"

/*
	参与者：
	Product（抽象产品）：具有指定操作的产品类型
	ConcreteProduct（具体产品）：具有抽象产品指定操作的具体产品
	Creator（抽象工厂）：创建抽象产品
	ConcreteFactory（具体工厂）：具有抽象工厂指定操作的具体工厂，负责创建具体产品

	创建对象的特点：
	创建的是一个抽象对象，只需要里面的操作就行了

	具体流程：
	1、创建具体工厂
	2、通过具体工厂创建具体产品但是返回的是抽象产品
	3、调用抽象产品可以调用里面的操作
*/

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
