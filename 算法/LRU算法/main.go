package main

/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/

type LRUCache struct {
	header   *node
	tail     *node
	capacity int
	cache    map[int]*node
}
type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

func Constructor(capacity int) LRUCache {
	header := &node{}
	tail := &node{}
	header.next = tail
	tail.prev = header
	return LRUCache{
		header:   header,
		tail:     tail,
		capacity: capacity,
		cache:    make(map[int]*node),
	}

}

func (this *LRUCache) Get(key int) int {
	if index, ok := this.cache[key]; ok {
		this.movetohead(index)
		return index.value
	}
	return -1
}
func (this *LRUCache) remove(node *node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (this *LRUCache) movetohead(node *node) {
	this.remove(node)
	node.next = this.header.next
	node.prev = this.header

	this.header.next.prev = node
	this.header.next = node
}
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.movetohead(node)
		return
	}
	node := &node{
		key:   key,
		value: value,
	}
	this.cache[key] = node
	node.next = this.header.next
	node.prev = this.header

	this.header.next.prev = node
	this.header.next = node

	if len(this.cache) > this.capacity {
		delete(this.cache, this.tail.prev.key)
		this.remove(this.tail.prev)

	}
}
