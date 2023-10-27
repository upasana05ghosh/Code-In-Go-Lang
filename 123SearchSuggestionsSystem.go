//https://leetcode.com/problems/search-suggestions-system/description/
func suggestedProducts(products []string, word string) [][]string {
	root := NewTrie()

	for _, v := range products {
		root.insert(v)
	}

	var res [][]string
	for _, v := range word {
		if root != nil {
			root = root.Child[v-'a']
		}

		if root == nil {
			res = append(res, []string{})
		} else {
			res = append(res, root.Suggest)
		}
	}

	return res
}

type Trie struct {
	Child   []*Trie
	Suggest []string
}

func NewTrie() *Trie {
	return &Trie{
		Child:   make([]*Trie, 26),
		Suggest: []string{},
	}
}

func (t *Trie) insert(p string) {
	for _, v := range p {
		if t.Child[v-'a'] == nil {
			t.Child[v-'a'] = NewTrie()
		}

		t = t.Child[v-'a']
		t.Suggest = append(t.Suggest, p)
		sort.Strings(t.Suggest)
		if len(t.Suggest) > 3 {
			t.Suggest = t.Suggest[:3] //pop
		}
	}
}

