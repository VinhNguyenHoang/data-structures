package trie

// InsertCOW implements the Copy-on-write pattern to create new nodes for modified data and return a new root.
//
// The new root will be the pointer to each "version" of the trie tree after each mutation.
func (t *Trie) InsertCOW(key string, value int) *Trie {
	cur := t.Root

	// create new root
	newTrie := NewTrie()
	parent := newTrie.Root

	// operation on an empty trie
	if len(cur.Children) == 0 {
		cur = parent
	} else {
		for k, node := range cur.Children {
			parent.Children[k] = node
		}
	}

	// find the node that need modified or to be inserted
	for _, c := range key {
		children, exist := cur.Children[string(c)]
		if exist {
			// create a new children with mutated data
			newchild := NewNode()

			// point all the children of the current child to new child (as parent)
			if len(children.Children) > 0 {
				for k, node := range children.Children {
					newchild.Children[k] = node
				}
			}

			// link the newly created child node to parent node
			parent.Children[string(c)] = newchild
			parent = newchild

			// move to the next inner node
			cur = newchild
		} else {
			// create new node and assign its parent
			node := NewNode()
			cur.Children[string(c)] = node

			// update nearest parent node
			parent = node

			// move to the next inner node
			cur = cur.Children[string(c)]

		}
	}

	cur.IsTerminal = true
	cur.Value = value

	return newTrie
}
