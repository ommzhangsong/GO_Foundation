package main

/* defer 是发生在return 返回值之后，或者是panic的时候
一般用在释放空间或者说和recover处理panic
*/

// 这里演示，defer是能读到命名返回值的值，并能修改，但是匿名返回值
// 是读不到的
import "fmt"

// 这里命名返回值，开始n=0 但是defer读到了返回值res = n =0
// 最后打印出来是返回是2
func deferrun() (res int) {
	var n int = 1
	defer func() {

		res++
	}()
	return n
}

// 匿名返回，go'会创建一个临时变量，存放这个返回值，首先会把n=1的返回值给这个temp
// 然后执行defer n++，局部变量n确实变成2了，但没有意义了，有意义的是
// temp它出去了，但他就是1
func deferrun2() int {
	var n int = 1
	defer func() {
		n++
	}()
	return n
}
func main() {
	res := deferrun()
	res2 := deferrun2()
	fmt.Println(res, res2)
}
