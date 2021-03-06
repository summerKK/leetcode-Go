package utils

import "fmt"

// Node is a generic interface for all nodes in a utils
type Node struct {
	Item interface{}
	Next *Node
}

// NewNode create a new Node
func NewNode(item interface{}, next *Node) *Node {
	return &Node{item, next}
}

// LinkList ...
type LinkList struct {
	First *Node
	n     int // size
}

// NewLinkList return a linked list
func NewLinkList() *LinkList {
	return &LinkList{nil, 0}
}

// Add adds a new node to head, points to old head
func (l *LinkList) Add(item interface{}) (node *Node) {
	node = NewNode(item, l.First)
	l.First = node
	l.n++
	return
}

// Del removes head and returns its Item
func (l *LinkList) Del() (item interface{}) {
	item = l.First.Item
	l.First = l.First.Next
	l.n--
	return
}

// 删除第k个元素.并返回.如果不存在返回nil
func (l *LinkList) DeleteK(k int) (item interface{}) {
	c := 0
	pre := l.First
	for node := l.First; node != nil; node = node.Next {
		c++
		pre = node
		if c+1 == k && pre.Next != nil {
			item := pre.Next.Item
			pre.Next = pre.Next.Next
			l.n--
			return item
		}
	}

	return nil
}

// 查找元素
func (l *LinkList) Find(item interface{}) bool {
	for node := l.First; node != nil; node = node.Next {
		if node.Item == item {
			return true
		}
	}

	return false
}

// Size returns number of nodes in link list
func (l *LinkList) Size() int {
	return l.n
}

// IsEmpty reports if the link list is empty
func (l *LinkList) IsEmpty() bool {
	return l.First == nil
}

// Slice returns a slice of link list for iterating
func (l *LinkList) Slice() (items []interface{}) {
	for curr := l.First; curr != nil; curr = curr.Next {
		items = append(items, curr.Item)
	}
	return
}

// IntSlice returns a slice of int values for iterating
func (l *LinkList) IntSlice() (items []int) {
	for curr := l.First; curr != nil; curr = curr.Next {
		items = append(items, curr.Item.(int))
	}
	return
}

// StringSlice returns a slice of int values for iterating
func (l *LinkList) StringSlice() (items []string) {
	for curr := l.First; curr != nil; curr = curr.Next {
		items = append(items, curr.Item.(string))
	}
	return
}

func (l *LinkList) RemoveAfter(node *Node) {
	if node == nil {
		return
	}

	// 计算删除的长度
	c := 0
	for v := node; v.Next != nil; v = v.Next {
		c++
	}

	node.Next = nil

	l.n -= c
}

func (l *LinkList) InsertAfter(node *Node) {
	if node == nil {
		return
	}
	tail := l.First
	c := 0
	for node := l.First; node != nil; node = node.Next {
		tail = node
		c++
	}

	tail.Next = node
	l.n += c
}

// 删除跟item值一样的所有节点
func (l *LinkList) Remove(item interface{}) {
	var pre *Node = nil
	for node := l.First; node != nil; node = node.Next {
		if node.Item == item {
			// 第一个元素就是要删除的节点
			if pre == nil {
				// 删除头节点
				l.First = node.Next
			} else {
				pre.Next = node.Next
			}
			l.n--
		} else {
			pre = node
		}
	}
}

func (l *LinkList) Loop() <-chan interface{} {
	c := make(chan interface{})
	go func() {
		for node := l.First; node != nil; node = node.Next {
			c <- node.Item
		}
		close(c)
	}()

	return c
}

func (l *LinkList) String() string {
	for v := range l.Loop() {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("len:%d", l.n)
	return ""
}
