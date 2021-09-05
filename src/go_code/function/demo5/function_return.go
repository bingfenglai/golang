package main

func main() {

	getPhone := sayHello2("厉飞羽")

	getPhone()

}

func sayHello2(name string) (func()) {
	println("你好！",name)
	return func() {
		println("我的手机号是： 131****3901")
	}
}
