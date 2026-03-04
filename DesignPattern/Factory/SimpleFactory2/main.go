/*
对于简单工厂模式，一些代码习惯上的优化
*/
package main

import "fmt"

// 抽象接口，去返回抽象产品
type payer interface {
	Pay(orderId int)
}

// 实现层
type Alipay struct{}

func (p *Alipay) Pay(orderId int) {
	fmt.Printf("使用支付宝支付，订单号：%v\n", orderId)
}

type Wechatpay struct{}

func (p *Wechatpay) Pay(orderId int) {
	fmt.Printf("使用wechat支付，订单号：%v\n", orderId)
}

// 用常量去代替魔法字符串：是编程中的坏味道 —— 指那些没有明确含义、直接写在代码里的字符串。大量魔法字符串会导致代码复杂性提高
const (
	PayTypeAlipay = "alipay"
	PayTypeWxpay  = "wxpay"
)

func create(paykind string) (payer, error) {
	switch paykind {
	case PayTypeAlipay:
		return new(Alipay), nil
	case PayTypeWxpay:
		return new(Wechatpay), nil
	default:
		return nil, fmt.Errorf("不支持该支付类型")

	}
}

func Execute(paykind string, orderID int) {
	payer, err := create(paykind)
	if err != nil {
		fmt.Println(err)
		return
	}
	payer.Pay(orderID)
}
func main() {
	Execute(PayTypeWxpay, 21312)
	Execute(PayTypeAlipay, 21312)
}
