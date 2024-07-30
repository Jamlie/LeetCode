type TrieNode struct {
	children map[byte]*TrieNode
	isEnd    bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
		isEnd:    false,
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: newTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	node := t.root

	for _, char := range word {
		ch := byte(char)
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = newTrieNode()
		}
		node = node.children[ch]
	}

	node.isEnd = true
}

func findWords(board [][]byte, words []string) []string {
	trie := NewTrie()

	for _, word := range words {
		trie.Insert(word)
	}

	result := make(map[string]bool)
	empty := bytes.Buffer{}
	for i := range len(board) {
		for j := range len(board[0]) {
			dfs(board, i, j, trie.root, empty, result)
		}
	}

	var res []string
	for word := range result {
		res = append(res, word)
	}
	return res
}

type A struct {
	bytes.Buffer
}

func dfs(board [][]byte, i, j int, node *TrieNode, path bytes.Buffer, result map[string]bool) {
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 {
		return
	}

	ch := board[i][j]
	if ch == '#' || node.children[ch] == nil {
		return
	}

	node = node.children[ch]
	path.WriteByte(ch)
	if node.isEnd {
		result[path.String()] = true
	}

	board[i][j] = '#'

	dfs(board, i-1, j, node, path, result)
	dfs(board, i+1, j, node, path, result)
	dfs(board, i, j-1, node, path, result)
	dfs(board, i, j+1, node, path, result)

	board[i][j] = byte(ch)
}
