package _SimpleFactory

import "fmt"

// 简单工厂模式：当你需要什么，只需要传入一个正确的参数，就可以获取你所需要的对象，而无须知道其创建细节

//在以下情况下可以考虑使用简单工厂模式：
//(1) 工厂类负责创建的对象比较少，由于创建的对象较少，不会造成工厂方法中的业务逻辑太过复杂。
//(2) 客户端只知道传入工厂类的参数，对于如何创建对象并不关心。

// 工厂角色1：工厂角色即工厂类，它是简单工厂模式的核心，负责实现创建所有产品实例的内部逻辑；工厂类可以被外界直接调用，创建所需的产品对象； 在工厂类中提供了静态的工厂方法factoryMethod()，它的返回类型为抽象产品类型Product。
type Donglaishun struct {
}

func (d *Donglaishun) GetFood() {
	fmt.Println("生产完毕，就绪。。。")
}

// 工厂角色2
type Qingfeng struct {
}

func (q *Qingfeng) GetFood() {
	fmt.Println("生产完毕，继续。。。。")
}

// 抽象产品角色：它是工厂类所创建的所有对象的父类，封装了各种产品对象的公有方法，它的引入将提高系统的灵活性，使得在工厂类中只需定义一个通用的工厂方法，因为所有创建的具体产品对象都是其子类对象。
type Restaurant interface {
	GetFood()
}

// 具体产品角色：它是简单工厂模式的创建目标，所有被创建的对象都充当这个角色的某个具体类的实例。每一个具体产品角色都继承了抽象产品角色，需要实现在抽象产品中声明的抽象方法。
func NewRestaurant(name string) Restaurant {
	switch name {
	case "d":
		return &Donglaishun{}
	case "q":
		return &Qingfeng{}
	}
	return nil
}
