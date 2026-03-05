/*
⭐⭐⭐适配器:主要解决接口不兼容的问题，
核心利用adapter，去组合新街口，并实现老接口的方法
*/

package main

import (
	"fmt"
	"strings"
)

// MessageSender  抽象层
// MessageSender 目标接口：系统统一的消息发送接口
type MessageSender interface {
	// Send 发送消息
	// 参数：
	//   to: 接收人（手机号/邮箱地址）
	//   content: 消息内容
	// 返回：发送结果
	Send(to string, content string) string
}

// SmsSender 短信发送器（已实现目标接口）
type SmsSender struct{}

func (s *SmsSender) Send(to string, content string) string {
	return fmt.Sprintf("短信发送成功：向%s发送%s", to, content)
}

// EmailService 第三方邮件服务（适配者，接口不兼容）
type EmailService struct{}

// SendEmail 第三方邮件发送方法（接口不匹配）
// 参数：
//
//	recipient: 收件人邮箱
//	subject: 邮件标题
//	body: 邮件正文
//
// 返回：发送结果
func (e *EmailService) SendEmail(recipient string, subject string, body string) string {
	return fmt.Sprintf("邮件发送成功\n向%s发送\n标题:%s\n内容:%s\n的邮件", recipient, subject, body)
}

// 这里是实现adpter的逻辑
type EmailAdapter struct {
	Es *EmailService
}

func Getsubject(content string) string {
	if has := strings.Contains(content, "|"); has {
		return content[:strings.Index(content, "|")]
	}
	return content
}
func Getbody(content string) string {
	if has := strings.Contains(content, "|"); has {
		return content[strings.Index(content, "|")+1:]
	}
	return content
}
func (emailadapter *EmailAdapter) Send(to string, content string) string {
	res := emailadapter.Es.SendEmail(to, Getsubject(content), Getbody(content))
	return res
}

func Newadpter(Es *EmailService) MessageSender {
	return &EmailAdapter{Es: Es}
}

func main() {
	es := Newadpter(new(EmailService))
	res := es.Send("123123123@qq.com", "基于深度学习的探讨|正文xxxxxxx")
	fmt.Println(res)
}
