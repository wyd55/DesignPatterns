package _Singleton

func IncrementAge() {
	p := GetInstance()
	p.IncrementAge()
}
