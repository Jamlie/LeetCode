type WordFinder struct {
	word  string
	board [][]byte
	rows  int
	cols  int
	path  map[[2]int]bool
}

func (wf *WordFinder) DFS(r, c, i int) bool {
	if i == len(wf.word) {
		return true
	}

	if r < 0 || c < 0 || r >= wf.rows || c >= wf.cols || wf.word[i] != wf.board[r][c] ||
		wf.path[[2]int{r, c}] {
		return false
	}

	wf.path[[2]int{r, c}] = true

	res := cmp.Or(
		wf.DFS(r+1, c, i+1),
		wf.DFS(r-1, c, i+1),
		wf.DFS(r, c+1, i+1),
		wf.DFS(r, c-1, i+1),
	)

	wf.path[[2]int{r, c}] = false

	return res
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}

	var (
		rows = len(board)
		cols = len(board[0])
	)

	wf := WordFinder{
		word:  word,
		board: board,
		rows:  rows,
		cols:  cols,
		path:  make(map[[2]int]bool),
	}

	for row := range rows {
		for col := range cols {
			if wf.DFS(row, col, 0) {
				return true
			}
		}
	}

	return false
}