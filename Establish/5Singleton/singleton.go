package _Singleton

import "sync"

var (
	p    *Pet
	once sync.Once // 只执行一次的操作，使用他来实现单例模式
)

func init() {
	// 只做一次
	once.Do(func() {
		p = &Pet{}
	})
}

func GetInstance() *Pet {
	return p
}

// Pet就是Singleton，他只能被实例化一次
type Pet struct {
	name string
	age  int
	mux  sync.Mutex
}

func (p *Pet) SetName(n string) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.name = n
}

func (p *Pet) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}

func (p *Pet) IncrementAge() {
	p.mux.Lock()
	p.age++
	p.mux.Unlock()
}

func (p *Pet) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age
}
