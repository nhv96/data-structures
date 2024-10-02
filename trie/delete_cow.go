package trie

// DeleteCOW implements the Copy-on-write pattern to create new nodes for modified data and return a new root.
//
// The new root will be the pointer to each "version" of the trie tree after each mutation.
//
// Given key "abc", check for child nodes of current root for key "a", then create
// another identical node of the node "a", copy over all it's pointers of children nodes to new node "a",
// then traverse into next key.
func (t *Trie) DeleteCOW(key string) *Trie {
	cur := &t.Root

	newRoot := NewTrie()
	newRoot.Root = *t.deletecow(cur, key)
	return newRoot
}

func (t *Trie) deletecow(cur **Node, key string) **Node {
	if key == "" {
		if (*cur).IsTerminal {
			if len((*cur).Children) > 0 {
				nn := NewNode()
				nn.Children = (*cur).Children
				return &nn
			}
			return nil
		}
		return cur
	}

	k := string(key[0])
	if childNode, exist := (*cur).Children[k]; exist {
		n := t.deletecow(&childNode, key[1:])

		// we just found that child node
		if n != &childNode {
			// if the current node has children
			if len((*cur).Children) > 0 {
				// now we create a different pointer to assign to
				// the new change
				// TODO: check if parent is an empty node and remove it instead of return empty node
				nn := NewNode()
				nn.IsTerminal = (*cur).IsTerminal
				nn.Value = (*cur).Value
				for kc, v := range (*cur).Children {
					if kc != k {
						nn.Children[kc] = v
					} else {
						if n != nil {
							nn.Children[kc] = *n
						}
					}
				}

				return &nn
			}
			// new parent
			np := NewNode()
			np.Children[k] = *n
			return &np
		}
	}
	return cur
}
