package main

import "fmt"

type phone_interface interface {
	call()
}
type apple struct {
	name string
}

func (p apple) call() {
	fmt.Println(p.name + "is calling... ")
}

type android struct {
	name string
}

func (p *android) call() {
	fmt.Println(p.name + "is calling... ")
}

func main() {
	/*
		实现接口的type是可以赋值给interface的
		要看好是，我们调用方法是用的值类型还是引用类型
	*/

	var a = apple{"Apple1 phone"} //初始化结构体 apple,这时候结构体就已经在内存中创建了
	// ❗❗注意这里我们是值类型赋值给接口，和方法正好是适配的
	var p1 phone_interface = a
	// 值赋值给，实际做的就是接口这个结构体中存的，指向了a的副本，并不是指向a
	a.name = "modified"
	p1.call()
	//new 返回指针，相当于&apple{}，这代表的是，我的借口内部存储的是*apple类型，
	// ❗❗这里指针类型分配给接口，恰好不适配，但是可以，因为语法糖
	// ✅✅ 恰好印证了那句话，接受者是值类型的，我们值 指针都可以继承接受者的方法
	var p2 phone_interface = new(apple)
	p2.(*apple).name = "Apple2 phone" // 断言，让其变成结构体指针，他们始终指向的都是结构体初始的地方，修改的是原数据
	p2.(*apple).name = "modified"
	p2.call()

	// 这里是恰好适配
	var p3 phone_interface = new(android)
	p3.(*android).name = "Android phone"
	p3.call()
	//var p4 phone_interface = android{"Android phone"}
	//上面这句话就会报错，因为android类型并没有实现方法，实现的是*android方法。
}
