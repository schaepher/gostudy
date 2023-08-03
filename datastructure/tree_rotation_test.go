package datastructure

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Left   *TreeNode
	Right  *TreeNode
	Value  int
}

// PutChild 用于简化初始化
func (n *TreeNode) PutChild(child *TreeNode) {
	if child.Value < n.Value {
		n.Left = child
	} else if child.Value > n.Value {
		n.Right = child
	}
}

// RotateRight 右旋，参数先不使用转轴
func RotateRight(root *TreeNode) {
	if root.Left == nil {
		return
	}

	// 把 A 备份出来
	a := root.Left.Right
	root.Left.Right = nil

	// 旋转。3 原先在 5 左节点，现在让 5 变成 3 的右节点
	newRoot := root.Left
	root.Left = nil
	newRoot.Right = root

	// 把 A 放回去
	root.Left = a
}

// PrintInorderIteration 中序遍历迭代法
func PrintInorderIteration(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Println(root.Value)

		root = root.Right
	}
}

func TestTree(t *testing.T) {
	// 为了与图对应，这里多申请一个位置，但只从 1 开始初始化。
	nodes := make([]*TreeNode, 10)
	for i := 1; i < 10; i++ {
		nodes[i] = &TreeNode{Value: i}
	}

	nodes[5].PutChild(nodes[3])
	nodes[5].PutChild(nodes[8])

	nodes[3].PutChild(nodes[2])
	nodes[3].PutChild(nodes[4])

	nodes[2].PutChild(nodes[1])

	nodes[8].PutChild(nodes[7])
	nodes[8].PutChild(nodes[9])

	nodes[7].PutChild(nodes[6])

	fmt.Println("BEGIN")
	PrintInorderIteration(nodes[5])
	fmt.Println("END")

	RotateRight(nodes[5])

	fmt.Println("BEGIN")
	PrintInorderIteration(nodes[3])
	fmt.Println("END")
}
