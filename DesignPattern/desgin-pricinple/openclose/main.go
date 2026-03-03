package main

import "fmt"

/*原本是一个类，下面有很多功能模块，如果我们要添加，因为各个
模块的耦合性，1.越来越多的模块会大大增加复杂性，越来越臃肿，2.如果我们要增加一个方法，我们无法保证
我们这个方法会不会影响到其他方法，也无法确定，当我们这个方法如果存在错误，会不会导致其他逻辑也无法运行
*/
/*
解决思想，openclose⭐⭐⭐⭐⭐：开闭原则
我们从原本的类，写一个抽象类（抽像接口）出来（interface），
用具体的对象去实现这个具体类，这样每个功能方法都是具体类去实现的，这样就将各个方法的耦合度大大降低
我们就可以不修改类（具体类）去达到增加模块功能的目的
并且最后我们还可以将其框架化，实现多态，调用一个方法实现不同具体类的实现
*/

/*
假设你正在开发电商平台的订单支付模块，初期产品只要求支持 支付宝 和 微信支付 两种方式。
最初的代码把所有支付逻辑都写在一个 OrderPayment 结构体里，包含 Alipay()、WechatPay() 两个方法；
现在产品提出新需求：要新增 银联支付 和 Apple Pay，未来还可能加更多支付方式；
核心要求：新增支付方式时，不能修改原有支付宝 / 微信支付的代码（避免引入 bug），且所有支付方式要能被统一调用。
*/
type Orderpayment interface {
	pay()
}

// 实现的具体类
type Alipay struct {
}

func (a *Alipay) pay() {
	fmt.Println("支付宝支付中...")
}

type WechatPay struct {
}

func (w *WechatPay) pay() {
	fmt.Println("wechat支付中...")
}

//++++++++新增逻辑

type ApplePay struct {
}

func (a *ApplePay) pay() {
	fmt.Println("Apple支付中...")
}

// 实现多态
func DoPay(o Orderpayment) {
	o.pay()
}
func main() {
	wechatPay := &WechatPay{}
	wechatPay.pay()
	alipay := &Alipay{}
	alipay.pay()
	applepay := &ApplePay{}
	applepay.pay()
	//多态
	DoPay(wechatPay)
	DoPay(alipay)
	DoPay(applepay)
}
