// https://leetcode.com/problems/implement-trie-prefix-tree/
// A trie (pronounced as "try") or prefix tree is a tree data structure used to
// efficiently store and retrieve keys in a dataset of strings.
// Implement the Trie class.
//
// Constraints:
// 1 <= word.length, prefix.length <= 2000
// word and prefix consist only of lowercase English letters.

package main

type Trie struct {
	Nodes map[byte]*Trie
	IsEnd bool
}

func InitTrie() Trie {
	return Trie{Nodes: map[byte]*Trie{}}
}

func (t *Trie) Insert(word string) {
	cur := t

	for i := 0; i < len(word); i++ {
		if _, ok := cur.Nodes[word[i]]; !ok {
			cur.Nodes[word[i]] = &Trie{Nodes: map[byte]*Trie{}}
		}
		cur = cur.Nodes[word[i]]
	}

	cur.IsEnd = true
}

func (t *Trie) Search(word string) bool {
	cur := t

	for i := 0; i < len(word); i++ {
		if _, ok := cur.Nodes[word[i]]; !ok {
			return false
		}
		cur = cur.Nodes[word[i]]
	}

	return cur.IsEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t

	for i := 0; i < len(prefix); i++ {
		if _, ok := cur.Nodes[prefix[i]]; !ok {
			return false
		}
		cur = cur.Nodes[prefix[i]]
	}

	return true
}
