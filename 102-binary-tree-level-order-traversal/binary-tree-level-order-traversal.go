type Collection interface {
	Len() int
	IsEmpty() bool
}

type Queue[T any] struct {
	q []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		q: make([]T, 0),
	}
}

func (q *Queue[T]) Enqueue(v T) {
	q.q = append(q.q, v)
}

func (q *Queue[T]) Deque() T {
	if len(q.q) == 0 {
		var t T
		return t
	}

	v := q.q[0]
	q.q = q.q[1:]

	return v
}

func (q Queue[T]) Len() int {
	return len(q.q)
}

func (q Queue[T]) IsEmpty() bool {
	return q.Len() == 0
}

type Iter[T Collection] struct {
	v T
}

func (q *Queue[T]) Iter() Iter[*Queue[T]] {
	return Iter[*Queue[T]]{
		v: q,
	}
}

func (q Iter[T]) ForEach(fn func()) {
	for !q.v.IsEmpty() {
		fn()
	}
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	BFS(root, &res)
	return res
}

func BFS(node *TreeNode, res *[][]int) {
	queue := NewQueue[*TreeNode]()

	queue.Enqueue(node)

	queue.Iter().ForEach(func() {
		level := make([]int, 0)

		for range queue.Len() {
			node := queue.Deque()

			if node != nil {
				level = append(level, node.Val)
				queue.Enqueue(node.Left)
				queue.Enqueue(node.Right)
			}
		}

		if len(level) != 0 {
			*res = append(*res, level)
		}
	})
}
