type TreeBuilder struct {
	preorder   []int
	inorderMap map[int]int
	preIndex   int
}

func (tb *TreeBuilder) buildTree(inStart, inEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}

	node := &TreeNode{
		Val: tb.preorder[tb.preIndex],
	}

	tb.preIndex++

	if inStart == inEnd {
		return node
	}

	inIndex := tb.inorderMap[node.Val]

	node.Left = tb.buildTree(inStart, inIndex-1)
	node.Right = tb.buildTree(inIndex+1, inEnd)
	return node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	m := make(map[int]int)

	for i, v := range inorder {
		m[v] = i
	}

	tb := &TreeBuilder{
		preorder:   preorder,
		inorderMap: m,
		preIndex:   0,
	}

	return tb.buildTree(0, len(inorder)-1)
}