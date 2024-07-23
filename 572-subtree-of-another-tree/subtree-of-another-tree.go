import "cmp"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return subRoot == nil
    }
    
	if root == subRoot {
		return true
	}

	if IsSameTree(root, subRoot) {
		return true
	}

	return cmp.Or(isSubtree(root.Left, subRoot), isSubtree(root.Right, subRoot))
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == q {
		return true
	}

	if p != nil && q != nil {
		return p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
	}

	return false
}