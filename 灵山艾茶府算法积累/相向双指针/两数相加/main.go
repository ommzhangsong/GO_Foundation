package main

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
就是链表相加，嘛，返回列表
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	var num1, num2, carry int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		} else {
			num1 = 0
		}
		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		} else {
			num2 = 0
		}
		sum := num1 + num2 + carry
		carry = (sum) / 10
		sum = sum % 10
		if head == nil {
			head = &ListNode{sum, nil}
			tail = head
		} else {
			tail.Next = &ListNode{sum, nil}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{carry, nil}
	}
	return head
}
