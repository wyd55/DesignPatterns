package _Singleton

func IncrementAge2() {
	p := GetInstance()
	p.IncrementAge()
}
