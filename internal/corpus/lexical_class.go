package corpus

// LexicalClass is a lexical class, with the name of that class and the words.
// attributed to it
type LexicalClass struct {
	class string
	words map[string]struct{}
}

// AddLexicalClass creates a new LexicalClass structure without words in that
// LexicalClass (use lc.AddWord to add words in the LexicalClass)
func AddLexicalClass(name string) *LexicalClass {
	return &LexicalClass{
		class: name,
		words: make(map[string]struct{}),
	}
}

// lc.AddWord adds a word to a LexicalClass (optional; may be more ergonomic
// than `E.words[word] = struct{}{}` for some LexicalClass E)
func (lc *LexicalClass) AddWord(word string) {
	lc.words[word] = struct{}{}
}
