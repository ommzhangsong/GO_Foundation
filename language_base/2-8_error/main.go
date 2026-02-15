package main

/* ERROR:支持多返回值，所以一般经常会在返回值上用到error
error 本身是一种接口类型，首先记住它是一种⚠️数据类型,
这个接口类型下面就只有一个Error()方法，可以用接口实现去实现
自定义的error

*/
import (
	"errors"
	"fmt"
)

// 一般用fmt.Errorof 或者errors.new去创建一个error对象，error对象不可比较
type Myerror struct {
	Code    int    `json:"code"` //错误码
	Message string `json:"msg"`
}

func (e Myerror) Error() string {
	// 用sprintf去拼接一下信息，方便打印，就是一个字符串
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Message)
}

func NewError(code int, msg string) error {
	// ❗这里需要注意的是，我们这里是把返回值，这个return里的内容
	// 当成是error返回出去，其实就相当于一个给接口error赋值的操作
	//如果是赋值指针的话，如下，我们在断言的时候，是断言的是指针，
	// ⚠️谁赋值给interface我们就断言什么，因为这相当于把这个类型装进了接口
	// 只有这个接口，才能取出我们想要的结果，所以断言是 err.(*Myerror)
	// 也就是Myerror的指针类型， 我们返回是返回的是值，要是地址所以有&
	// 推荐返回指针，内存占用小，
	return &Myerror{Code: code, Message: msg}
}
func main() {
	a := fmt.Errorf("worong!")
	b := errors.New("wrong")
	fmt.Println(a)
	fmt.Println(b)
	err := NewError(100, "success")
	fmt.Println(err)
	//
	fmt.Println(err.(*Myerror).Code)
}
