package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type people struct {
	Height int
	Weight int
}
type home struct {
	CityName string
}
type student struct {
	Name     string `json:"username" db:"username" validate:"required"` //自定义的结构体标签，在json结构中： username:"zhang"
	Age      int    `json:"age,omitempty"`                              // 如果是零值那么不会出现在json中
	Password string `json:"-"`                                          //省略
	people          //隐式“继承”
	Hometown home   //显式继承，
	//	只有结构体的组合，没有继承
}

// 方法绑定给自定义类型，一般都是绑结构体
// 先声明接受类型，这里是student， 然后记得要写return,方法其实也就是一个函数
// 但是这个函数，它绑定到了一个固定的结构体上去了。

func (st student) GetName() string {
	return st.Name
}

// 这里演示的是，因为结构体是值类型，如果只传递值的话，改变的只是复制值
func (st student) changeAge() {
	st.Age += 1
}
func (st *student) changePassword() {
	st.Password = "1234"
}

// 封装个error懒人函数，exitonError
func exitonError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1) //1:出错退出
	}
}
func main() {

	s := student{
		Name:     "user1",
		Age:      20,
		Password: "1111",
		people:   people{179, 80}, //隐式，前后名字一样，都是类型
		Hometown: home{"xuancheng"},
	}
	s.changeAge()
	s.changePassword()
	//隐式继承，隐藏了，就相当于直接把属性拉过来用，
	fmt.Println(s, s.Hometown.CityName, s.Height)
	//	结构体json编码,这时候需要注意的是，并不是说别的包必须用我们的时候，才算导出，
	// 比如，此时json是个第三方包，但它也用我们的结构体，这个也是导出
	// 所以如果我们结构体中hometown没有大写的话，我们就不会序列化，或者说不会被json看到
	// 如果我们要传给前端，一般都是显式。
	byte, err := json.Marshal(s)
	exitonError(err)
	fmt.Println(string(byte))

	//    空结构体的应用，因为struct{}是占用0字节的，所有的空的struct指向零地址的空间

	set := make(map[int]struct{})
	set[1] = struct{}{}
	set[200] = struct{}{}
	_, ok := set[200]
	if ok {
		fmt.Println(ok)
	}

}
