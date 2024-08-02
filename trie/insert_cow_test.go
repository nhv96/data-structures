package trie

import "testing"

func Test_InsertCOW(t *testing.T) {
	t.Parallel()

	t.Run("insert", func(t *testing.T) {
		trie0 := NewTrie()
		trie1 := trie0.InsertCOW("a", 1)

		node1, found1 := find(trie1, "a")

		if found1 != true {
			t.Errorf("Key \"%s\" is expected to be \"%v\", got \"%v\"", "a", true, found1)
		}

		if found1 && node1.Value != 1 {
			t.Errorf("Key \"%s\" is expected to have value \"%v\", got \"%v\"", "a", 1, node1.Value)
		}

		node0, found0 := find(trie0, "a")

		if found0 != false || node0 != nil {
			t.Errorf("Key \"%s\" is expected to be \"%v\", got \"%v\"", "a", false, found0)
		}
	})
}
