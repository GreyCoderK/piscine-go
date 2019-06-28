package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	gauche := BTreeLevelCount(root.Left)
	droite := BTreeLevelCount(root.Right)

	if droite > gauche {
		return droite + 1
	}else{
		return gauche + 1
	}
}
