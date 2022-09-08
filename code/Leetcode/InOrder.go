package leetcode

func main() {

}

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	inorder := make([]int, 0, 200)

	inord(root, &inorder)

	return inorder
}

func inord(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	(*arr) = append((*arr), root.Val)
	inord(root.Left, arr)
	inord(root.Right, arr)

}
