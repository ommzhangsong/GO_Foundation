/*
工厂方法模式⭐⭐⭐⭐⭐：
简单工厂加上开闭原则，简单工厂在增加产品的时候，逻辑必须对原本的工厂进行内部修改
所以这里，实现了抽象工厂的概念，让某个产品有自己的创建逻辑，
当我们需要新增的时候，只需添加产品的类，以及创建它的工厂类就可以了
*/
/*
电商平台初期支持支付宝、微信支付两种支付方式；
产品要求：后续可能快速新增银联支付、京东支付、ApplePay等支付方式；
核心约束：新增支付方式时，绝对不能修改已有的工厂代码（避免影响线上稳定的支付宝 / 微信支付逻辑）。
*/
package main

import "fmt"

// 抽象接口，统一最后的执行
type payer interface {
	Pay(orderId int)
}

// 抽象工厂，去创建每个具体的类，这里有抽象类，去创建各个具体的工厂
type factory interface {
	create() payer
}

const (
	PayTypeAlipay = "alipay"
	PayTypeWxpay  = "wxpay"
	PayTypeJDpay  = "JDpay"
)

// 实现层
type Alipay struct {
}

type AlipayFactory struct {
	factory
}

func (x *AlipayFactory) create() payer {
	return &Alipay{}
}
func (x *Alipay) Pay(orderId int) {
	fmt.Printf("支付宝正在支付%v\n", orderId)
}

type Wxpay struct {
}
type WxpayFactory struct {
	factory
}

func (x *WxpayFactory) create() payer {
	return &Wxpay{}
}
func (x *Wxpay) Pay(orderId int) {
	fmt.Printf("wx正在支付%v\n", orderId)
}

// +++++++++++新增
type JDpay struct {
}
type JDpayFactory struct {
	factory
}

func (x *JDpayFactory) create() payer {
	return &JDpay{}
}
func (x *JDpay) Pay(orderId int) {
	fmt.Printf("京东正在支付%v\n", orderId)
}

// 包级函数
func initfac(paytype string) (factory, error) {
	switch paytype {
	case PayTypeAlipay:
		return &AlipayFactory{}, nil
	case PayTypeWxpay:
		return &WxpayFactory{}, nil
	case PayTypeJDpay:
		return &JDpayFactory{}, nil
	default:
		return nil, fmt.Errorf("没有此支付类型")
	}
}
func exeute(paytype string, orderId int) {
	fac, err := initfac(paytype)
	if err != nil {
		return
	}
	fac.create().Pay(orderId)
}
func main() {
	// 下面如何封装，实现多态，通过传参，初始化不同的factory
	//factory := &AlipayFactory{}
	//ali := factory.create()
	//ali.Pay(122)
	//
	//factory2 := &WxpayFactory{}
	//WX := factory2.create()
	//WX.Pay(122)
	exeute(PayTypeAlipay, 1543523)
	exeute(PayTypeWxpay, 1231231213)

	//新增
	exeute(PayTypeJDpay, 1231231213)
}
