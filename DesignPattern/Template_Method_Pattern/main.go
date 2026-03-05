/*
模版方法模式，通过定义模版，来完成相似的流程
核心是，接口的内嵌
*/
/*
题目背景
电商系统中，不同类型的订单（普通订单、秒杀订单、预售订单）需要走统一的 “订单处理流程”，但部分步骤的具体逻辑不同：
统一流程（算法骨架）：参数校验 → 库存锁定 → 价格计算 → 生成订单 → 消息通知
差异化步骤：
库存锁定：普通订单锁定普通库存；秒杀订单锁定秒杀专属库存；预售订单锁定预售库存（且需校验定金）；
价格计算：普通订单按原价 + 优惠券；秒杀订单按秒杀价（忽略优惠券）；预售订单按定金 + 尾款计算；
消息通知：普通订单推送短信；秒杀订单推送短信 + APP 弹窗；预售订单推送短信 + 尾款提醒。
题目要求
基于 Go 语言的模板方法模式，设计并实现上述订单处理逻辑；
要求代码具备可扩展性（新增订单类型时最小化修改）；
加入必要的异常处理（如库存不足、参数非法时返回错误）；
编写测试代码验证三种订单的处理流程。
*/
package main

// 抽象类。接口，高可替换性
type OrderProcess interface {
	Stock(Order *Order)
	Price(Order *Order)
	Message(Order *Order)
}

// 实现层，订单结构
type Order struct {
	OrderID   string  // 订单ID
	OrderType string  // 订单类型：normal/seckill/preSale
	ProductID string  // 商品ID
	UserID    string  // 用户ID
	Price     float64 // 最终价格
	StockNum  int     // 购买数量
	Deposit   float64 // 预售定金（仅预售订单有效）
}

// BaseModel内嵌接口，实现差异化行为的替代
type BaseModel struct {
	orderprocess OrderProcess
}

func (bm *BaseModel) Process(order *Order) error {
	//step1 : 参数校验
	//step2 : 参数校验
	//step3 : 参数校验
	//step4 : 参数校验
	//step5 : 参数校验
	return nil
}
