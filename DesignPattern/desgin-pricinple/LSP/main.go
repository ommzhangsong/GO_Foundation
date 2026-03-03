package main

import "fmt"

/*
依赖倒转原则⭐⭐⭐⭐⭐，模块和模块的代码耦合度太高，需要理解代码的压力很大
将原本的，两层结构，具体类+实现层，变成 业务逻辑层+抽象层+实现层
核心是：面相接口的编程
业务逻辑层(实现层 )向下依赖，依赖这个抽象层，
实现层向上依赖这个抽象层，千万不能用业务逻辑层与实现层建立耦合

*/
/*
场景：假设你开发电商平台的「订单配送模块」：
初期产品要求支持 顺丰快递 和 京东物流 两种配送方式；
业务逻辑层是 OrderDispatchService（订单配送服务），负责处理订单的配送逻辑；
核心要求：
遵循依赖倒转原则：订单配送服务只能依赖抽象层，不能直接依赖顺丰 / 京东的具体实现；
扩展性：新增「圆通快递」时，无需修改订单配送服务的任何代码；
核心三层结构：业务逻辑层（订单配送服务）→ 抽象层（配送接口）→ 实现层（各快递公司）。
*/
// 抽象层
type Delivery interface {
	Deliver(orderID string)
}

// 实现层

// 顺丰快递
type SF struct {
}

func (s *SF) Deliver(orderID string) {
	fmt.Printf("SF正在送%v的快递\n", orderID)
}

// 京东快递
type JD struct {
}

func (s *JD) Deliver(orderID string) {
	fmt.Printf("京东正在送%v的快递\n", orderID)
}

// +++++++++++++新增圆通
type YT struct {
}

func (s *YT) Deliver(orderID string) {
	fmt.Printf("圆通正在送%v的快递\n", orderID)
}

// 业务逻辑层:依赖倒转原则的关键是「业务逻辑层（如 OrderDispatchService）」，你直接在 main 里操作具体实现，没有体现 “业务层依赖抽象”；
// 正确的做法是：业务逻辑层是一个独立的结构体，内部只保存抽象接口，通过 Set 方法注入具体实现。
type OrderDispatchService struct {
	// 核心：依赖抽象接口，而非具体的SF/JD/YT,或者说这里继承了一个接口类型，抽象接口
	delivery Delivery
}

// 这里去设置上面struct的delivery的具体类型是谁
func (ods *OrderDispatchService) SetDelivery(d Delivery) {
	ods.delivery = d
}
func (ods *OrderDispatchService) Dispatch(orderID string) {
	ods.delivery.Deliver(orderID) // 依赖抽象，而非具体实现！
}

func main() {
	//这里用的就是delivery抽象的借口,实现多态，这里还可以进一步封装
	//var who Delivery
	//who = &SF{}
	//who.Deliver("123")
	//who = &JD{}
	//who.Deliver("123")

	dispatchService := &OrderDispatchService{}
	dispatchService.SetDelivery(new(SF))
	dispatchService.Dispatch("123")
	dispatchService.SetDelivery(new(JD))
	dispatchService.Dispatch("123")

}
