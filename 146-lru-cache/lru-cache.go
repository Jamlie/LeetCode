type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func NewNode(key, value int) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

type LRUCache struct {
	capacity int
	left     *Node
	right    *Node
	cache    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		capacity: capacity,
		left:     NewNode(0, 0),
		right:    NewNode(0, 0),
		cache:    make(map[int]*Node),
	}

	lruCache.left.prev, lruCache.right.prev = lruCache.right, lruCache.left

	return lruCache
}

func (lc *LRUCache) Remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next, next.prev = next, prev
}

func (lc *LRUCache) Insert(node *Node) {
	prev, next := lc.right.prev, lc.right
	prev.next = node
	next.prev = node

	node.next, node.prev = next, prev
}

func (lc *LRUCache) Get(key int) int {
	if v, ok := lc.cache[key]; ok {
		lc.Remove(v)
		lc.Insert(v)

		return v.value
	}

	return -1
}

func (lc *LRUCache) Put(key int, value int) {
	if v, ok := lc.cache[key]; ok {
		lc.Remove(v)
	}

	lc.cache[key] = NewNode(key, value)
	lc.Insert(lc.cache[key])

	if len(lc.cache) > lc.capacity {
		lru := lc.left.next
		lc.Remove(lru)
		delete(lc.cache, lru.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */