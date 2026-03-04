/*

工厂模式都很简单，工厂模式的思想：复杂的创建初始化的逻辑可以从业务逻辑层剥离，由工厂去做，
业务逻辑层无需关心创建初始的逻辑，降低业务逻辑层与创建初始化逻辑的耦合度，简化业务逻辑层。
简单工厂模式⭐⭐⭐⭐⭐，简单工厂模式不属于gof的23种，因为它不符合‘开闭原则’。
它是用具体工厂，去实现的抽象产品
*/
/*
假设你正在开发一个电商平台的支付模块，平台支持多种支付方式：
支付宝支付（Alipay）：需要初始化商户 ID、私钥，支付时会打印 "使用支付宝支付，订单号：XXX"
微信支付（WeChatPay）：需要初始化 APPID、商户密钥，支付时会打印 "使用微信支付，订单号：XXX"
银联支付（UnionPay）：需要初始化机构号、终端号，支付时会打印 "使用银联支付，订单号：XXX"
需求：
业务层只需要传入支付类型（如 "alipay"、"wechat"、"unionpay"）和订单号，就能调用对应支付方式的支付功能
支付方式的创建、初始化逻辑（如配置密钥、ID 等）与业务层解耦，业务层无需关心具体初始化细节

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

type factory struct{}

func (f *factory) create(paykind string) payer {
	if paykind == "alipay" {
		return &Alipay{}
	} else if paykind == "wechatpay" {
		return &Wechatpay{}
	}
	return nil
}

func main() {
	factory := new(factory)
	alipay := factory.create("alipay")
	alipay.Pay(12321321)

	wechatpay := factory.create("wechatpay")
	wechatpay.Pay(12321321)
}
