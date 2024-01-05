package tree

// TrieNode represents a node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie represents the Trie data structure
type Trie struct {
	root *TrieNode
}

// NewTrie initializes a new Trie
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert inserts a word into the Trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

// Search checks if a word exists in the Trie
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if child, exists := node.children[char]; exists {
			node = child
		} else {
			return false
		}
	}
	return node.isEnd
}

// Autocomplete suggests possible completions based on the input
func (t *Trie) Autocomplete(prefix string) []string {
	node := t.root
	for _, char := range prefix {
		if child, exists := node.children[char]; exists {
			node = child
		} else {
			return nil
		}
	}

	var suggestions []string
	t.collectSuggestions(node, prefix, &suggestions)
	return suggestions
}

// collectSuggestions recursively collects suggestions from the Trie
func (t *Trie) collectSuggestions(node *TrieNode, currentPrefix string, suggestions *[]string) {
	if node.isEnd {
		*suggestions = append(*suggestions, currentPrefix)
	}

	for char, child := range node.children {
		t.collectSuggestions(child, currentPrefix+string(char), suggestions)
	}
}
