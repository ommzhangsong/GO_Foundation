package main

/*
合成复用原则⭐⭐⭐⭐可以用组合，尽量用组合（也就是我们的命名嵌入）
伪继承（匿名嵌入），这样外部的对象和继承对象的耦合程度高，外部对象继承了所有方法，
父对象可能有什么方法进行改变，外部对象也会改变，这样就导致了逻辑不清晰，而且如果有同名方法
还会有覆盖的情况，优先子方法，没有在往父节点找
*/
type sayer interface {
	say()
}
type student struct {
	sayer //匿名继承
}

type student2 struct {
	sayername sayer //命名方式
}

func main() {}
