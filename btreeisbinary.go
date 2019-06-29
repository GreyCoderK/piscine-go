package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return bst(root, root.Left, root.Right)
}

func bst(root, l, r *TreeNode) bool {
	if root == nil {
		return true
	}
	if l != nil && root.Data < l.Data {
        	return false
	}
	if r != nil && root.Data > r.Data {
        	return false
	}
	return bst(root.Left, l, root) && bst(root.Right, root, r)
}
