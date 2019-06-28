package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil || root.Right == nil {
		return root
	}

	return BTreeMax(root.Right)
}
