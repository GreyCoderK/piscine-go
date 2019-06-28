package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	BTreeApplyPostorder(root.Right, f)
	BTreeApplyPostorder(root.Left, f)
	f(root.Data)
}
