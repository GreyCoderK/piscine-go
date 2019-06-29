package piscine

func BTreeIsBinary(root *TreeNode) bool {
        if root == nil {
                return true
        }
	
	if (root.Left == nil && root.Right != nil) || (root.Left != nil && root.Right == nil) {
		return false
	}else if root.Left != nil && root.Right != nil {
		return true
	}
        
	if root.Left.Data > root.Data || root.Right.Data < root.Data {
                return false
        }
        
	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}
