package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	replacement := node
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		replacement.Parent.Left = rplc
	} else {
		replacement.Parent.Right = rplc
	}
	replacement.Parent = node.Parent

	return root
}
