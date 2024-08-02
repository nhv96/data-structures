package trie

// InsertCOW implements the Copy-on-write pattern to create new nodes for modified data and return a new root.
//
// The new root will be the pointer to each "version" of the trie tree after each mutation.
func (t *Trie) InsertCOW(key string, value int) *Trie {
	cur := t.Root

	// create new root
	newRoot := NewNode()

	// find the node that need modified or to be inserted
	for _, c := range key {
		children, exist := cur.Children[string(c)]
		if exist {
			cur = children
		} else {
			// create new node and assign its parent as the newRoot?
			node := NewNode()
			newRoot.Children[string(c)] = node
			cur = newRoot.Children[string(c)]
		}
	}

	cur.IsTerminal = true
	cur.Value = value

	newTrie := NewTrie()
	newTrie.Root = newRoot
	return newTrie
}
