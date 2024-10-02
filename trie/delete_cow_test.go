package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DeleteCOW(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("delete a leaf node", func(t *testing.T) {
		trie0 := NewTrie()
		trie0.Insert("aa", 1)
		trie0.Insert("ab", 2)
		trie0.Insert("bce", 3)

		trie1 := trie0.DeleteCOW("bce")
		node1, found1 := find(trie1, "bce")
		is.Equal(false, found1)
		is.Nil(node1)

		node0, found0 := find(trie0, "bce")
		is.Equal(true, found0)
		is.Equal(3, node0.Value)

		trie2 := trie1.DeleteCOW("a")
		node2, found2 := find(trie2, "aa")
		is.Equal(true, found2)
		is.Equal(1, node2.Value)
	})

	t.Run("delete on an empty tree", func(t *testing.T) {
		trie0 := NewTrie()

		trie1 := trie0.DeleteCOW("a")

		node1, found1 := find(trie1, "a")
		is.Equal(false, found1)
		is.Nil(node1)
	})
}
