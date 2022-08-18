package leetcode

import "fmt"

type Node struct {
	Val  int
	Next *Node
}
type ListNode struct {
	head *Node
	tail *Node
	size int
}

//初始化
func (list *ListNode) Init() {
	list.size = 0   // 此时链表是空的
	list.head = nil // 没有头
	list.tail = nil // 没有尾
}

//获取元素
func (list *ListNode) Get(i int) *Node {
	if i >= list.size {
		return nil
	}

	item := list.head
	for j := 0; j < i; j++ { // 从head数i个
		item = item.Next
	}

	return item
}

//插入数据
func (list *ListNode) Insert(i int, node *Node) bool {
	// 空的节点、索引超出范围和空链表都无法做插入操作
	if node == nil || i > list.size || list.size == 0 {
		return false
	}

	if i == 0 { // 直接排第一，也就领导小舅子才可以
		node.Next = list.head
		list.head = node
	} else {
		// 找到前一个元素
		preItem := list.head
		for j := 1; j < i; j++ { // 数前面i个元素
			preItem = preItem.Next
		}
		// 原来元素放到新元素后面,新元素放到前一个元素后面
		node.Next = preItem.Next // 新节点指向旧节点原来所指的next
		preItem.Next = node      // 原节点的next指向新节点
	}

	list.size++

	return true
}

//删除元素
func (list *ListNode) Remove(i int) bool {
	if i >= list.size {
		return false
	}

	if i == 0 { // 删除头部
		node := list.head
		list.head = node.Next
		if list.size == 1 { // 如果只有一个元素，那尾部也要调整
			list.tail = nil
		}
	} else {
		preItem := list.head
		for j := 1; j < i; j++ {
			preItem = preItem.Next
		}

		node := preItem.Next
		preItem.Next = node.Next

		if i == (list.size - 1) { // 若删除的尾部，尾部指针需要调整
			list.tail = preItem
		}
	}
	list.size--
	return true
}

//向尾部添加数据
func (list *ListNode) Append(node *Node) bool {
	if node == nil {
		return false
	}

	node.Next = nil
	// 将新元素放入单链表中
	if list.size == 0 {
		list.head = node
	} else {
		oldTail := list.tail
		oldTail.Next = node
	}

	// 调整尾部位置，及链表元素数量
	list.tail = node // node成为新的尾部
	list.size++      // 元素数量增加

	return true
}

func ListNode_Solve() {
	var list = ListNode{}
	list.Init()
	for i := 0; i < 10; i++ {
		var node = Node{Val: i}
		list.Append(&node)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(list.Get(i))
	}
	var node = list.Get(5)
	fmt.Printf("当前节点位置：%p ,数据：%d \n", node, node.Val)
	var result = list.Remove(5)
	fmt.Printf("删除是否成功%t \n", result)

	var node2 = list.Get(5)
	fmt.Printf("当前节点位置：%p ,数据：%d \n", node2, node2.Val)

}
