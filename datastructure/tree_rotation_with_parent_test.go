package datastructure

import (
	"fmt"
	"testing"
)

type TreeNodeWithParent struct {
	Parent *TreeNodeWithParent
	Left   *TreeNodeWithParent
	Right  *TreeNodeWithParent
	Value  int
}

func (n *TreeNodeWithParent) PutChild(n1 *TreeNodeWithParent) {
	if n1.Value < n.Value {
		n.Left = n1
	} else if n1.Value > n.Value {
		n.Right = n1
	} else {
		return
	}

	n1.Parent = n
}

func PrintInorderTransversal(n *TreeNodeWithParent) {
	if n.Left != nil {
		PrintInorderTransversal(n.Left)
	}
	fmt.Println(n.Value)
	if n.Right != nil {
		PrintInorderTransversal(n.Right)
	}
}

func PrintInorderTransversalIteration(root *TreeNodeWithParent) {
	stack := make([]*TreeNodeWithParent, 0)
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

// RotateRightWithParent 右旋，参数先不使用转轴
func RotateRightWithParent(root *TreeNodeWithParent) {
	if root.Left == nil {
		return
	}

	// 把 A 备份出来，取消双向连线
	a := root.Left.Right
	root.Left.Right = nil
	if a != nil {
		a.Parent = nil
	}

	// 取到子树根节点的父节点，并取消双向连线
	rootParent := root.Parent
	root.Parent = nil
	// 如果 root 是整颗树的根节点，无需调整
	if rootParent != nil {
		if rootParent.Value > root.Value {
			rootParent.Left = nil
		} else {
			rootParent.Right = nil
		}
	}

	// 旋转。2 原先在 3 左节点，现在让 3 变成 2 的右节点
	newRoot := root.Left
	root.Left = nil
	newRoot.Parent = nil

	newRoot.Right = root
	root.Parent = newRoot

	// 设置根节点的父节点
	if rootParent != nil {
		if rootParent.Value > newRoot.Value {
			rootParent.Left = newRoot
		} else {
			rootParent.Right = newRoot
		}
	}

	// 把 A 放回去
	root.Left = a
}

// RotateRightWithParent 右旋，参数使用转轴
func RotateRightWithParentPivot(pivot *TreeNodeWithParent) {
	if pivot.Parent == nil {
		return
	}

	// 把 A 备份出来，取消双向连线
	a := pivot.Right
	pivot.Right = nil
	if a != nil {
		a.Parent = nil
	}

	// 取到子树根节点的父节点，并取消双向连线
	rootParent := pivot.Parent.Parent
	pivot.Parent.Parent = nil
	// 如果 root 是整颗树的根节点，无需调整
	if rootParent != nil {
		if rootParent.Value > pivot.Value {
			rootParent.Left = nil
		} else {
			rootParent.Right = nil
		}
	}

	// 旋转。
	root := pivot.Parent
	root.Left = nil
	pivot.Parent = nil

	pivot.Right = root
	root.Parent = pivot

	// 设置根节点的父节点
	if rootParent != nil {
		if rootParent.Value > pivot.Value {
			rootParent.Left = pivot
		} else {
			rootParent.Right = pivot
		}
	}

	// 把 A 放回去
	root.Left = a
}


func TestRotate(t *testing.T) {
	// 为了与图对应，这里多申请一个位置，但只从 1 开始初始化。
	nodes := make([]*TreeNodeWithParent, 10)
	for i := 1; i < 10; i++ {
		nodes[i] = &TreeNodeWithParent{Value: i}
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
	PrintInorderTransversalIteration(nodes[5])
	fmt.Println("END")

	RotateRightWithParent(nodes[5])

	fmt.Println("BEGIN")
	PrintInorderTransversalIteration(nodes[3])
	fmt.Println("END")
}