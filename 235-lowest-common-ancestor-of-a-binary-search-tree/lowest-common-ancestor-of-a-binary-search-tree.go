/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

type Comparable int

const (
	LessThan Comparable = iota
	GreaterThan
	LessThanEq
	GreaterThanEq
)

func compareTo[T cmp.Ordered](a T) func(T, Comparable) bool {
	return func(o T, comp Comparable) bool {
		switch comp {
		case LessThan:
			return a < o
		case GreaterThan:
			return a > o
		case LessThanEq:
			return a <= o
		case GreaterThanEq:
			return a >= o
		}

		return false
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	rootComp := compareTo(root.Val)

	if rootComp(p.Val, LessThan) && rootComp(q.Val, LessThan) {
		return lowestCommonAncestor(root.Right, p, q)
	} else if rootComp(p.Val, GreaterThan) && rootComp(q.Val, GreaterThan) {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return root
}
