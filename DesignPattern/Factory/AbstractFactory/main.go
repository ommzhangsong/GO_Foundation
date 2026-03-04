/*
抽象工厂模式⭐⭐⭐⭐⭐，
相比工厂方法模式，如果我们要新增产品的时候，我们需要新增对应的具体工厂，
所以抽象工厂模式，从产品族的方向，去用一个工厂生产多个产品
这里就要说到了产品族和产品等级，产品等级不同体系的的纵向分类
产品族，同一场景维度的产品集合
*/
/*
某电商平台需要重构支付模块，核心需求：
支持三大产品族（支付宝体系、微信支付体系、京东支付体系）；
每个产品族包含两个核心产品等级：
「支付等级」：实现订单支付功能；
「退款等级」：实现订单退款功能；
*/
package main

import "fmt"

// 抽象接口，

type payer interface {
	Pay(orderId int)
}
type refunder interface {
	Refund(orderId int)
}

type AbastractFactory interface {
	Createpayer() payer
	Createrefund() refunder
}
type Alipay struct {
}
type Alirefund struct{}

func (x *Alipay) Pay(orderId int)       { fmt.Printf("支付宝正在支付%v\n", orderId) }
func (x *Alirefund) Refund(orderId int) { fmt.Printf("支付宝正在退款%v\n", orderId) }

type AlipayFactory struct{}

func (x *AlipayFactory) Createpayer() payer     { return &Alipay{} }
func (x *AlipayFactory) Createrefund() refunder { return &Alirefund{} }

func main() {
	fc := AlipayFactory{}
	fc.Createpayer().Pay(5)
	fc.Createrefund().Refund(5)
}
