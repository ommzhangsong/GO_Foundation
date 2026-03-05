/*
⭐⭐⭐⭐代理模式，主要是用于创建一个新的对象去代理之前的对象，完成具体的任务需求
代理除了有这个组合的功能，也有这个装饰器的功能
核心动作：将接口当作结构体的字段去传递，实现组合
*/
package main

import "fmt"

// 抽象层
type Buyer interface {
	Buy(good string)
}

// 实现层
type chinashopping struct {
}

func (c chinashopping) Buy(good string) {
	fmt.Println("正在中国买", good)
}

type proxy struct {
	b Buyer
}

func (p *proxy) Buy(good string) {
	//额外功能：验证真伪
	fmt.Println("验证是否真假", good)
	p.b.Buy(good)
}
func newProxy(b Buyer) *proxy {
	return &proxy{b}
}

func main() {
	cs := newProxy(&chinashopping{})
	cs.Buy("apple")

}
