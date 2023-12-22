package hello

type TrieNode struct {
	Children map[rune]*TrieNode
	End      bool
}

type Trie struct {
	Root *TrieNode
}

func (t *Trie) Insert(str string) {
	root := t.Root
	for _, i := range str {
		if root.Children[i] == nil {
			root.Children[i] = &TrieNode{Children: map[rune]*TrieNode{}}
		}
		root = root.Children[i]
	}
	root.End = true
}

func (t *Trie) Contains(str string) bool {
	root := t.Root
	for _, i := range str {
		if root.Children[i] != nil {
			root = root.Children[i]
		} else {
			return false
		}
	}
	return true
}

func (t *Trie) Suffix(str string) {
	for i := 0; i < len(str); i++ {
		strrr := str[i:]
		t.Insert(strrr)
	}
}

func (t *Trie) Delete(word string) {
	t.deleteHelper(t.Root, word, 0)
}

func (t *Trie) deleteHelper(node *TrieNode, word string, index int) bool {
	if node == nil {
		return false
	}

	if index == len(word) {
		if !node.End {
			return false
		}
		node.End = false

		return len(node.Children) == 0
	}

	char := rune(word[index])
	child, exists := node.Children[char]
	if !exists || t.deleteHelper(child, word, index+1) {
		delete(node.Children, char)
		return len(node.Children) == 0
	}

	return false
}
