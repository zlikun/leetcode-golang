package main

import (
	"fmt"
	"strings"
)

// 单向链表
type ListNode struct {
	Val  int       // 节点值
	Next *ListNode // 下一节点指针
}

// 遍历链表
func (head *ListNode) Traverse() {
	for head != nil {
		fmt.Print(head.Val, "\t")
		head = head.Next
	}
	fmt.Println()
}

func test1(fn func(*ListNode, *ListNode) *ListNode) {
	// 2 -> 4 -> 3
	m := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	m.Traverse()

	// 5 -> 6 -> 4
	n := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	n.Traverse()

	// 7 -> 0 -> 8
	r := fn(m, n)
	r.Traverse()
}

func test2(fn func(*ListNode, *ListNode) *ListNode) {
	// 1 -> 8
	m := &ListNode{Val: 1, Next: &ListNode{Val: 8}}
	m.Traverse()

	// 0
	n := &ListNode{}
	n.Traverse()

	// 7 -> 0 -> 8
	r := fn(m, n)
	r.Traverse()
}

func test3(fn func(*ListNode, *ListNode) *ListNode) {
	// 5
	m := &ListNode{Val: 5}
	m.Traverse()

	// 5
	n := &ListNode{Val: 5}
	n.Traverse()

	// 0 -> 1
	r := fn(m, n)
	r.Traverse()
}

func main() {
	// 第一组测试
	test1(addTwoNumbers1)
	test2(addTwoNumbers1)
	test3(addTwoNumbers1)

	fmt.Println(strings.Repeat("-", 32))

	// 第二组测试
	test1(addTwoNumbers2)
	test2(addTwoNumbers2)
	test3(addTwoNumbers2)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 时间复杂度：O(max(m,n))
// 空间复杂度：O(max(m,n))
func addTwoNumbers1(m *ListNode, n *ListNode) *ListNode {
	if m == nil {
		return n
	}
	if n == nil {
		return m
	}
	head := new(ListNode)
	curr := head
	var carry int
	for m != nil || n != nil {
		// 计算当前节点值，头节点为整数最低位，求和 >=10 则向后进一位
		var x, y int
		if m != nil {
			x = m.Val
		}
		if n != nil {
			y = n.Val
		}
		sum := x + y + carry
		curr.Val = sum % 10
		carry = sum / 10
		if m != nil {
			m = m.Next
		}
		if n != nil {
			n = n.Next
		}
		if m != nil || n != nil {
			curr.Next = new(ListNode)
			curr = curr.Next
		}
	}

	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}

	return head
}

// 根据他人解答的优解实现，比方案一性能好（多次提交有时好，有时差）
func addTwoNumbers2(m *ListNode, n *ListNode) *ListNode {
	if m == nil {
		return n
	}
	if n == nil {
		return m
	}

	head := new(ListNode) // 创建一个链表头节点
	curr := head          // 链表迭代节点

	// carry 用于记录l1,l2节点之和，同时用于记录进位
	for carry := 0; m != nil || n != nil || carry != 0; carry /= 10 {
		if m != nil {
			carry += m.Val
			m = m.Next
		}
		if n != nil {
			carry += n.Val
			n = n.Next
		}
		curr.Val = carry % 10
		if m != nil || n != nil || carry >= 10 {
			curr.Next = new(ListNode)
			curr = curr.Next
		}
	}

	return head
}
