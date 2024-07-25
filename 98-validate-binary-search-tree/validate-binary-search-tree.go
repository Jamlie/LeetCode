func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt, math.MaxInt)
}

func validate(node *TreeNode, minimum, maximum int) bool {
	if node == nil {
		return true
	}

	if node.Val <= minimum || node.Val >= maximum {
		return false
	}

	return validate(node.Left, minimum, node.Val) && validate(node.Right, node.Val, maximum)
}