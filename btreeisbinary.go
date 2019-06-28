package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	
	if root.Left != nil {
		return root.Left.Data < root.Data && BTreeIsBinary(root.Left)
	}else if root.Right != nil {
                return root.Right.Data > root.Data && BTreeIsBinary(root.Right)
        }else{
		return false
	}
}
