package tools

import "fmt"

type Object interface{}
type Node struct {
	Data Object //定义数据域
	Next *Node  //定义地址域（指向下一个表的地址）
}
type List struct {
	headNode *Node //头节点
}

/*单链表是一种顺序存储的结构。
有一个头结点，没有值域，只有连域，专门存放第一个结点的地址。
有一个尾结点，有值域，也有链域，链域值始终为NULL。
在单链表中为找第i个结点或数据元素，必须先找到第i-1结点或数据元素，而且必须知道头结点，否则整个链表无法访问*/

//1.判断链表是否为空
func (head *List) IsEmpty() bool {
	if head.headNode == nil {
		return true
	} else {
		return false
	}
}

//2.获取链表长度
func (head *List) Length() int {
	cur := head.headNode //获取链表头节点
	count := 0           //定义计时器
	for cur != nil {     //如果头节点不为空，则count++
		count++
		cur = cur.Next //对地址逐个位移
	}
	return count
}

//3.从链表头部添加元素
func (head *List) Add(data Object) *Node {
	node := &Node{Data: data}
	node.Next = head.headNode //头节点赋值给node的下一个结点
	head.headNode = node      //node赋值给头节点
	return node
}

//4.从链表尾部添加元素
func (head *List) Append(data Object) {
	node := &Node{Data: data}
	if head.IsEmpty() { //如果链表为空，直接将新元素作为头节点
		head.headNode = node
	} else {
		cur := head.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node //将新结点添加在尾节点之后
	}
}

//6.从链表指定位置添加元素
func (head *List) Insert(index int, data Object) {
	if index < 0 {
		head.Add(data)
	} else if index > head.Length() {
		head.Append(data)
	} else {
		pre := head.headNode
		count := 0
		for count < (index - 1) {
			pre = pre.Next
			count++
		}
		node := &Node{Data: data}
		node.Next = pre.Next
		pre.Next = node
	}
}

//7.删除链表指定值的元素
func (head *List) Remove(data Object) {
	pre := head.headNode  //定义变量pre，存储头节点
	if pre.Data == data { //如果是头节点，删除头节点
		head.headNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}

	}
}

//8.删除指定位置的元素
func (head *List) RemoveAtIndex(index int) {
	pre := head.headNode
	if index <= 0 {
		head.headNode = pre.Next
	} else if index > head.Length() {
		fmt.Println("超出链表长度")
		return
	} else {
		count := 0
		for count != (index-1) && pre.Next != nil {
			count++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}

//9.查看链表是否包含某个元素
func (head *List) Contain(data Object) bool {
	cur := head.headNode
	for cur != nil {
		if cur.Data == data {
			return true
		}
		cur = cur.Next
	}
	return false
}

//10.遍历链表所有节点
func (head *List) ShowList() {
	if !head.IsEmpty() {
		cur := head.headNode
		for {
			fmt.Printf("node:%v\t", cur.Data)
			if cur.Next != nil {
				cur = cur.Next
			} else {
				break
			}
		}
		fmt.Println()
	}
}

func Test() {
	list := List{}
	fmt.Println("链表长度：", list.Length())
	list.Append(1)
	list.Append(2)
	// fmt.Println("链表长度：", list.Length())
	// list.ShowList()
	// fmt.Println()
	// list.Add(3)
	// list.ShowList()
	// fmt.Println()
	// fmt.Println("链表长度：", list.Length())
	// list.Append(4)
	// list.ShowList()
	// list.Remove(4)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.ShowList()
	list.ShowList()
}
