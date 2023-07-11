package _Prototype

import "fmt"

// 原型接口
type Prototype interface {
	Clone() Prototype
}

// 具体原型
type Template struct {
	ID   int
	Name string
}

func (t *Template) Clone() Prototype {
	return &Template{
		ID:   t.ID,
		Name: t.Name,
	}
}

// 客户端
func PrototypeMain() {
	// 创建一个原型对象
	original := &Template{
		ID:   1,
		Name: "原始模板",
	}

	// 克隆原型对象
	clone := original.Clone().(*Template)

	// 修改克隆对象的属性
	clone.Name = "克隆模板"

	fmt.Println(original)
	fmt.Println(clone)
}
