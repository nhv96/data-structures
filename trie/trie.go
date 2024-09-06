package trie

// Node represent a node in trie tree
type Node struct {
	Children   map[string]*Node // use string type as key for better debugging
	IsTerminal bool
	Value      int
}

// NewNode returns a new node
func NewNode() *Node {
	return &Node{
		Children: make(map[string]*Node),
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
		children, exist := cur.Children[string(c)]
		if exist {
			cur = children
		} else {
			cur.Children[string(c)] = NewNode()
			cur = cur.Children[string(c)]
		}
	}

	cur.IsTerminal = true
	cur.Value = value
}

// Find finds the value of key and return. Returns false if not found
func (t *Trie) Find(key string) (int, bool) {
	cur := t.Root

	for _, c := range key {
		children, exist := cur.Children[string(c)]
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

// Delete deletes a key from the tree.
// Given key "abc", check for child nodes of "a" with keys "bc", then check for child nodes of "b" with key "c",
// then check for child nodes of "c" til we find a terminal node, then update it's value to false/empty,
// and return any other child nodes if exist, or return nil to indicate we deleted the node.
func (t *Trie) Delete(cur *Node, key string) *Node {
	// case we reached the end of the key string
	if key == "" {
		if cur.IsTerminal {
			cur.IsTerminal = false
			// we don't need to update the Value here,
			// but maybe when we implement the tree with different data type, we will need to update it to nil
		}
		// check if the current node has any children
		if len(cur.Children) > 0 {
			return cur
		}
		// otherwise return nil, this will delete the node
		return nil
	}

	k := string(key[0])
	// otherwise keep traveling into the next layer and update current node.
	if childNode, exist := cur.Children[k]; exist {
		cur.Children[k] = t.Delete(childNode, key[1:])
		// the child node doesn't have any more children, we delete it that key
		if cur.Children[k] == nil {
			delete(cur.Children, k)

			// if current node is empty, we delete it too
			if cur.IsTerminal == false {
				return nil
			}
		}
	}
	return cur
}
