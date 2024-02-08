package main

import "fmt"

const SIZE = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Len  int
	Head *Node
	Tail *Node
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Len)
	for i := 0; i < q.Len; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Len-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

type Hash map[string]*Node

type Cash struct {
	Queue Queue
	Hash  Hash
}

func NewCash() Cash {
	return Cash{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head
	return Queue{Head: head, Tail: tail}
}

func (c *Cash) Check(word string) {
	node := &Node{}
	if val, ok := c.Hash[word]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: word}
	}
	c.Add(node)
	c.Hash[word] = node

}

func (c *Cash) Remove(n *Node) *Node {
	fmt.Printf("Remove %s\n", n.Val)
	left := n.Left
	right := n.Right
	left.Right = right
	right.Left = left

	c.Queue.Len--
	delete(c.Hash, n.Val)
	return n
}

func (c *Cash) Add(n *Node) {
	fmt.Printf("Add %s\n", n.Val)
	temp := c.Queue.Head.Right
	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = temp
	temp.Left = n
	c.Queue.Len++
	if c.Queue.Len > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}

}

func (c *Cash) Display() {
	c.Queue.Display()
}

func main() {
	cash := NewCash()
	for _, word := range []string{"avocado", "bob", "sun", "avocado", "dog", "sun", "person", "cat", "dog", "grass"} {
		cash.Check(word)
		cash.Display()

	}

}
