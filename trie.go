package gotrie

// Node represent a node in trie tree
type Node struct {
	Children   map[rune]*Node
	IsTerminal bool
	Value      int
}

// NewNode returns a new node
func NewNode() *Node {
	return &Node{
		Children: make(map[rune]*Node),
	}
}

// NewTrie returns a new trie
func NewTrie() *Trie {
	t := &Trie{}
	t.Root = NewNode()

	return t
}

// Trie is the implementation of trie tree
type Trie struct {
	Root *Node
}

// Insert inserts kv data into the tree
func (t *Trie) Insert(key string, value int) {
	cur := t.Root

	for _, c := range key {
		children, exist := cur.Children[c]
		if exist {
			cur = children
		} else {
			cur.Children[c] = NewNode()
			cur = cur.Children[c]
		}
	}

	cur.IsTerminal = true
	cur.Value = value
}

// Find finds the value of key and return. Returns false if not found
func (t *Trie) Find(key string) (int, bool) {
	cur := t.Root

	for _, c := range key {
		children, exist := cur.Children[c]
		if !exist {
			return 0, false
		}
		cur = children
	}

	if cur.IsTerminal {
		return cur.Value, cur.IsTerminal
	}

	return 0, false
}
