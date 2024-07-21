package gotrie

import (
	"testing"
)

func Test_Trie(t *testing.T) {
	t.Parallel()

	t.Run("insert", func(t *testing.T) {
		trie := NewTrie()

		trie.Insert("hello", 10)
		trie.Insert("apply", 42)
		trie.Insert("hell", 666)
		trie.Insert("helloworld", 2024)

		testCases := []struct {
			key          string
			expectedVal  int
			expectedTerm bool
		}{
			{"hello", 10, true},
			{"world", 0, false},
			{"app", 0, false},
			{"hell", 666, true},
		}

		for _, tc := range testCases {
			node, found := find(trie, tc.key)
			if found != tc.expectedTerm {
				t.Errorf("Key \"%s\" is expected to be \"%v\", got \"%v\"", tc.key, tc.expectedTerm, found)
			}

			if found && node.Value != tc.expectedVal {
				t.Errorf("Key \"%s\" is expected to have value \"%v\", got \"%v\"", tc.key, tc.expectedVal, node.Value)
			}
		}

	})

	t.Run("find", func(t *testing.T) {
		trie := NewTrie()

		trie.Insert("hello", 10)
		trie.Insert("apply", 42)
		trie.Insert("hell", 666)
		trie.Insert("helloworld", 2024)

		testCases := []struct {
			key          string
			expectedVal  int
			expectedTerm bool
		}{
			{"hello", 10, true},
			{"hhaha", 0, false},
			{"apple", 0, false},
			{"helloworld", 2024, true},
		}

		for _, tc := range testCases {
			val, found := trie.Find(tc.key)
			if found != tc.expectedTerm {
				t.Errorf("Key \"%s\" is expected to be \"%v\", got \"%v\"", tc.key, tc.expectedTerm, found)
			}
			if found && val != tc.expectedVal {
				t.Errorf("Key \"%s\" is expected to have value \"%v\", got \"%v\"", tc.key, tc.expectedVal, val)
			}
		}
	})
}

// helper to find key node in he tree
func find(t *Trie, key string) (*Node, bool) {
	cur := t.Root

	for _, c := range key {
		child, exist := cur.Children[c]
		if !exist {
			return nil, false
		}
		cur = child
	}

	return cur, cur.IsTerminal
}
