package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeFromSlice(slice []int, index int, null int) (node *TreeNode) {
	if index >= len(slice) {
		return
	}
	value := slice[index]
	if slice[index] == null {
		return
	}
	node = &TreeNode{Val: value}
	node.Left = NewTreeFromSlice(slice, 2*index+1, null)
	node.Right = NewTreeFromSlice(slice, 2*index+2, null)
	return
}
