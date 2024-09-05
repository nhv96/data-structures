package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InsertCOW(t *testing.T) {
	t.Parallel()

	t.Run("new node must be added", func(t *testing.T) {
		is := assert.New(t)
		trie0 := NewTrie()
		trie1 := trie0.InsertCOW("a", 1)

		node1, found1 := find(trie1, "a")
		is.Equal(true, found1)
		is.Equal(true, node1.IsTerminal)
		is.Equal(1, node1.Value)

		node0, found0 := find(trie0, "a")
		is.Equal(false, found0)
		is.Nil(node0)
	})

	t.Run("new node must be added and retain old nodes", func(t *testing.T) {
		is := assert.New(t)
		trie0 := NewTrie()
		trie0.Insert("aa", 1)
		trie0.Insert("ab", 2)
		trie0.Insert("bc", 3)

		trie1 := trie0.InsertCOW("ba", 4)

		testSteps := []struct {
			name         string
			tree         *Trie
			key          string
			expectedVal  int
			expectedTerm bool
		}{
			{
				name:         "find in trie0 and must not find \"a\"",
				tree:         trie0,
				key:          "a",
				expectedVal:  0,
				expectedTerm: false,
			},
			{
				name:         "find in trie0 and must find key \"aa\"",
				tree:         trie0,
				key:          "aa",
				expectedVal:  1,
				expectedTerm: true,
			},
			{
				name:         "find in trie1 and must not find key \"a\"",
				tree:         trie1,
				key:          "a",
				expectedVal:  0,
				expectedTerm: false,
			},
			{
				name:         "find in trie1 and must find key \"aa\"",
				tree:         trie1,
				key:          "aa",
				expectedVal:  1,
				expectedTerm: true,
			},
			{
				name:         "find in trie1 and must find key \"ab\"",
				tree:         trie1,
				key:          "ab",
				expectedVal:  2,
				expectedTerm: true,
			},
			{
				name:         "find in trie1 and must find \"bc\"",
				tree:         trie1,
				key:          "bc",
				expectedVal:  3,
				expectedTerm: true,
			},
			{
				name:         "find in trie1 and must find \"ba\"",
				tree:         trie1,
				key:          "ba",
				expectedVal:  4,
				expectedTerm: true,
			},
			{
				name:         "find in trie0 and must not find \"ba\"",
				tree:         trie0,
				key:          "ba",
				expectedVal:  0,
				expectedTerm: false,
			},
		}

		for _, steps := range testSteps {
			node, found := find(steps.tree, steps.key)

			is.Equalf(steps.expectedTerm, found, "Step '%s' expected to be %v, got %v", steps.name, steps.expectedTerm, found)
			if steps.expectedTerm {
				is.Equalf(steps.expectedVal, node.Value, "Step '%s' expected to be %v, got %v", steps.name, steps.expectedVal, node.Value)
			}
		}
	})

	t.Run("update data of a node", func(t *testing.T) {
		is := assert.New(t)

		trie0 := NewTrie()
		trie0.Insert("aa", 1)
		trie0.Insert("ab", 2)
		trie0.Insert("bc", 3)

		trie1 := trie0.InsertCOW("ba", 4)

		trie2 := trie1.InsertCOW("bc", 100)

		node2, found2 := find(trie2, "bc")
		is.Equal(true, found2)
		is.Equal(100, node2.Value)

		_, found0 := find(trie0, "ba")
		is.Equal(false, found0)

		_ = trie2.InsertCOW("aa", 10)
		node0, found0 := find(trie0, "aa")
		is.Equal(true, found0)
		is.Equal(1, node0.Value)

		node2, found2 = find(trie2, "aa")
		is.Equal(true, found2)
		is.Equal(1, node2.Value)
	})
}
