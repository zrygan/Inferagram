package corpus

// Dictionary is a structure that contains DictionaryEntries.
type Dictionary struct {
	entries map[string]*LexicalClass
}

// AddDictionary creates an Dictionary structure.
func AddDictionary() *Dictionary {
	return &Dictionary{
		entries: make(map[string]*LexicalClass),
	}
}

// d.AddEntry adds an entry to some Dictionary d given the word and its
// LexicalClass.
func (d *Dictionary) AddEntry(word string, lex *LexicalClass) {
	d.entries[word] = lex
}
