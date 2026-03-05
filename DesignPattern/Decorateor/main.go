/*
⭐⭐⭐装饰器模式，和代理模式比较像，但是实现的东西很简单，只是多加一些方法，
因为实现逻辑简单，所有可以有多个装饰器，并且装饰器之间相互组合
*/
/*
这里去实现一个 手机的装饰器，手机本身的功能是show
然后我们设置多个装饰器，分别能给手机加上其他功能，比如放音乐，玩游戏
并且装饰器可以互相组合
*/
package main

import "fmt"

// 抽象层
type phone interface {
	show()
}
type decorater interface {
	phone //"继承phone的方法，必须实现原本的show方法
}

// 实现层
type apple struct {
}

func (a *apple) show() {
	fmt.Println("打开苹果手机")
}

type huawei struct {
}

func (a *huawei) show() {
	fmt.Println("打开huawei手机")
}

type playmusic_dec struct {
	decorater
}

func (p *playmusic_dec) show() {
	p.decorater.show()
	fmt.Println("正在播放")
}

func newdepm(p phone) phone {
	return &playmusic_dec{decorater: p}
}

type playgame_dec struct {
	decorater
}

func (p *playgame_dec) show() {
	p.decorater.show()
	fmt.Println("正在玩")
}

func newdepg(p phone) phone {
	return &playgame_dec{decorater: p}
}

func main() {
	a := newdepm(new(huawei))
	a.show()
	b := newdepg(a)
	b.show()

}
