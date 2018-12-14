package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	keys, root := TestTreapInsert()
	TestTreapErase(root, keys)
}

type Node struct {
	key int
	size int
	priority int
	left, right *Node
}

func NewNode(key int) *Node {
	return &Node{key: key, size: 1, priority: RandInt(math.MaxInt32), left: nil, right: nil}
}

func (targetNode *Node) ReCalculate() {
	size := 1

	if targetNode.left != nil {
		size += targetNode.left.size
	}

	if targetNode.right != nil {
		size += targetNode.right.size
	}

	targetNode.size = size
}

func (targetNode *Node) SetLeft(newNode *Node) *Node {
	targetNode.left = newNode
	targetNode.ReCalculate()
	return targetNode
}

func (targetNode *Node) SetRight(newNode *Node) *Node {
	targetNode.right = newNode
	targetNode.ReCalculate()
	return targetNode
}

func (targetNode *Node) Delete() {
	// 레퍼런스를 끊어준다.
	targetNode.left = nil
	targetNode.right = nil
}

// 이미 있는 값이 들어오는 경우 어떻게?
func Insert(root *Node, newNode *Node) *Node {
	if root == nil {
		return newNode
	}

	if newNode.priority <= root.priority {
		if newNode.key >= root.key {
			return root.SetRight(Insert(root.right, newNode))
		} else {
			return root.SetLeft(Insert(root.left, newNode))
		}
	} else {
		left, right := split(root, newNode.key)
		newNode.SetLeft(left)
		newNode.SetRight(right)
		return newNode
	}
}

func split(root *Node, key int) (*Node, *Node) {
	if root == nil {
		return nil, nil
	}

	if key > root.key {
		newLeft, newRight := split(root.right, key)
		root.SetRight(newLeft)
		return root, newRight
		/*
		right := root.right
		if right == nil {
			return root, nil
		}

		// 오른쪽 링크를 끊는다
		root.SetRight(nil)

		if key > right.key {
			// 더 쪼개야하는 경우
			newLeft, newRight := split(right, key)
			root.SetRight(newLeft)
			return root, newRight
		} else {
			// 원하는 위치를 찾은 경우
			return root, right
		}
		*/
	} else {
		newLeft, newRight := split(root.left, key)
		root.SetLeft(newRight)
		return newLeft, root
		/*
		left := root.left
		if left == nil {
			return nil, root
		}

		// 왼쪽 링크를 끊는다.
		root.SetLeft(nil)

		if key < left.key {
			// 더 쪼개야하는 경우
			newLeft, newRight := split(left, key)
			root.SetLeft(newRight)
			return newLeft, root
		} else {
			// 원하는 위치를 찾은 경우
			return left, root
		}
		*/
	}
}

func Erase(root*Node, key int) (newRoot *Node) {
	if root == nil {
		return nil
	}

	if root.key == key {
		// 왼쪽 서브트리와 오른쪽 서브트리를 합쳐 새로운 루트를 만들어 반환한다.
		newRoot := merge(root.left, root.right)
		root.Delete()
		return newRoot
	}

	// 삭제할 노드를 발견해서 새로운 노드가 생긴 경우 링크 시킨다.
	if key > root.key {
		return root.SetRight(Erase(root.right, key))
	} else {
		return root.SetLeft(Erase(root.left, key))
	}
}

func merge(a *Node, b *Node) (newRoot *Node) {
	// 한쪽 서브트리가 비어있는 경우다.
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	// 트리의 재귀적인 속성을 이용하여 머지한다.
	if a.priority > b.priority {
		return a.SetRight(merge(a.right, b))
	} else {
		return b.SetLeft(merge(a, b.left))
	}
}

func RandInt(max int) int {
	return rand.Intn(max)
}

func TestTreapInsert() ([]int, *Node) {
	keys := make([]int, 0)
	nodes := make([]*Node, 0)
	for i := 0; i < 30; i++ {
		newKey := RandInt(100)
		newNode := NewNode(newKey)
		fmt.Printf("val: %v\n", newNode.key)
		nodes = append(nodes, newNode)
		keys = append(keys, newKey)
	}

	fmt.Println("=======================")

	root := nodes[0]
	for i := 1; i < 30; i++ {
		root = Insert(root, nodes[i])
	}
	traverse(root)

	return keys, root
}

func TestTreapErase(root *Node, keys []int) {
	for i := 0; i < 30; i++ {
		root = Erase(root, keys[i])
	}
}

func traverse(root *Node) {
	if root == nil {
		return
	}

	traverse(root.left)
	fmt.Printf("key: %8v, priority: %14v\n", root.key, root.priority)
	traverse(root.right)
}