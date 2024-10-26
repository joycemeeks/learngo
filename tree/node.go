package tree

import (
	"fmt"
)

type TreeNode struct {
	Left, Right *TreeNode
	Value       int
}

func createTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func (node *TreeNode) GetValue() {
	fmt.Println(node.Value)
}

func (node *TreeNode) SetValue(value int) {
	node.Value = value
}

func main() {
	root := TreeNode{Value: 0}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{nil, nil, 2}
	root.Left.Left = new(TreeNode)

	fmt.Println(root)

	root.Left.GetValue()
	root.Left.SetValue(2)
	root.Left.GetValue()
}
