package symmetric_tree

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

const null = -101

func fromLeft(root *TreeNode, fun func(value int)) {
	if root == nil {
		fun(null) // null 也回傳值
		return
	}
	fun(root.Val)
	fromLeft(root.Left, fun) // 不特別先檢查 null
	fromLeft(root.Right, fun)
}

func fromRight(root *TreeNode, fun func(value int)) {
	if root == nil {
		fun(null)
		return
	}
	fun(root.Val)
	fromRight(root.Right, fun)
	fromRight(root.Left, fun)
}

func toStream(root *TreeNode, fun func(root *TreeNode, fun func(value int))) <-chan int {
	streamCh := make(chan int)
	go func() {
		defer close(streamCh)
		fun(root, func(value int) {
			streamCh <- value
		})
	}()
	return streamCh
}

/*
Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	streamFromLeft := toStream(root, fromLeft)
	streamFromRight := toStream(root, fromRight)
	isSymmetric := true
	for {
		leftValue, ok := <-streamFromLeft // golang 的傳統：零值會用額外的 boolean 來判斷是否為異常
		if !ok {
			break
		}
		rightValue := <-streamFromRight
		// fmt.Printf("%v %v\n", leftValue, rightValue)
		if leftValue != rightValue {
			isSymmetric = false
			break
		}
	}
	return isSymmetric
}
