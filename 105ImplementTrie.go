//https://leetcode.com/problems/implement-trie-prefix-tree/description
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, ch := range word {
		i := ch - 'a'
		if curr.children[i] == nil {
			curr.children[i] = &Trie{}
		}
		curr = curr.children[i]
	}
	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, ch := range word {
		i := ch - 'a'
		if curr.children[i] == nil {
			return false
		}
		curr = curr.children[i]
	}
	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, ch := range prefix {
		i := ch - 'a'
		if curr.children[i] == nil {
			return false
		}
		curr = curr.children[i]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */