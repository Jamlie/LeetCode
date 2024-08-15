type BrowserHistory struct {
	history *LinkedList[string]
	current *list.Element
}

func Constructor(homepage string) BrowserHistory {
	l := NewLinkedList[string]()
	curr := l.PushBack(homepage)
	return BrowserHistory{
		history: l,
		current: curr,
	}
}

func (bh *BrowserHistory) Visit(url string) {
	for bh.history.Len() > 0 && bh.history.Back() != bh.current {
		bh.history.Remove(bh.history.Back())
	}

	bh.current = bh.history.PushBack(url)
}

func (bh *BrowserHistory) Back(steps int) string {
	for steps > 0 && bh.current.Prev() != nil {
		bh.current = bh.current.Prev()
		steps--
	}

	return bh.history.Value(bh.current)
}

func (bh *BrowserHistory) Forward(steps int) string {
	for steps > 0 && bh.current.Next() != nil {
		bh.current = bh.current.Next()
		steps--
	}

	return bh.history.Value(bh.current)
}

type LinkedList[T any] struct {
	list *list.List
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{list: list.New()}
}

func (ll *LinkedList[T]) PushBack(value T) *list.Element {
	return ll.list.PushBack(value)
}

func (ll *LinkedList[T]) PushFront(value T) *list.Element {
	return ll.list.PushFront(value)
}

func (ll *LinkedList[T]) InsertAfter(value T, mark *list.Element) *list.Element {
	return ll.list.InsertAfter(value, mark)
}

func (ll *LinkedList[T]) InsertBefore(value T, mark *list.Element) *list.Element {
	return ll.list.InsertBefore(value, mark)
}

func (ll *LinkedList[T]) Remove(e *list.Element) T {
	return ll.list.Remove(e).(T)
}

func (ll *LinkedList[T]) Front() *list.Element {
	return ll.list.Front()
}

func (ll *LinkedList[T]) Back() *list.Element {
	return ll.list.Back()
}

func (ll *LinkedList[T]) Len() int {
	return ll.list.Len()
}

func (ll *LinkedList[T]) Value(e *list.Element) T {
	return e.Value.(T)
}