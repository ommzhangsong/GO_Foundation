package main

import "fmt"

type node struct {
	ele  any
	next *node
}
type List struct {
	len  int
	head *node
}

//NewlinkList创建空的列表

func NewlinkList() *List {
	return &List{
		len:  0,
		head: nil,
	}
}

// 头插法，简单明了
func (List *List) headadd(val any) {
	newnode := &node{val, List.head}
	List.head = newnode
	List.len++
}

// 尾插法，需要遍历找到最后一个tail，这也是为什么大部分实现链表都用头插法，最后看起来是倒序的样子，它不想数组可以随机访问，
//链表的查询效率是比较低的

func (l *List) tailadd(val any) {
	newnode := &node{val, nil}
	if l.head == nil {
		l.head = newnode
		l.len++
		return
	}
	lastnode := l.head
	for lastnode.next != nil {
		lastnode = lastnode.next
	}
	lastnode.next = newnode
	l.len++
}

func (l *List) getall() {
	cur := l.head
	if l.isempty() {
		fmt.Println("这个链表为空，还没有初始化呢")
		return
	}
	if cur.next == nil {
		fmt.Printf("只有一个头节点:%v", cur.ele)
	}

	for i := range l.len {
		fmt.Printf("第%v个节点的值是：%v\n", i+1, cur.ele)
		cur = cur.next
	}
	return
}
func (l *List) length() {
	fmt.Println("链表长度是", l.len)
}
func (l *List) isempty() bool {
	if l.head == nil {
		return true
	}
	return false
}

func main() {
	list1 := NewlinkList()

	fmt.Println(list1.isempty())
	list1.length() // 0
	list1.headadd(2)
	list1.headadd(1)
	list1.tailadd(3)
	list1.length()
	list1.getall()

}
