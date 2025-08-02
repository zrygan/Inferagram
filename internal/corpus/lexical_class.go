package corpus

import "fmt"

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

// Print displays the tokenized corpus, sentence by sentence.
func (c *AbstractedCorpus) PrintLexicalClass() {
	for _, sentence := range c.Abstractions {
		for _, lex := range sentence {
			fmt.Print(lex.class, " { ")
			i := len(lex.words)
			j := 0
			for words := range lex.words {
				fmt.Print(words)

				if j > i {
					fmt.Print(", ")
					j++
				}
			}

			j = 0
			fmt.Println(" }")
		}
	}
}

// LexicalClassMap is a mapping lexical class (non-structure) -> (words ...);
// a mapping lexical classes and words that fall under that lexical class.
type LexicalClassMap map[string]map[string]struct{}

// lcm.Add adds a lexical class (non-structure) to a LexicalClassMap (if not
// exist) and a word to that class (if provided).
func (lcm *LexicalClassMap) Add(class string, word ...string) {
	if _, exists := (*lcm)[class]; !exists {
		(*lcm)[class] = make(map[string]struct{})
	}

	for _, w := range word {
		(*lcm)[class][w] = struct{}{}
	}
}

func (lcm LexicalClassMap) IsUnique(class string, word ...string) bool {
	// if _, exists := (*lcm)[class]; exists {

	// }
	return true
}
