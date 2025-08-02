package corpus

import (
	"fmt"
	"strconv"
)

// AbstractedCorpus represents a corpus that only contains lexical classes.
type AbstractedCorpus struct {
	Abstractions [][]LexicalClass
}

// AddAbstractedCorpus created an AbstractedCorpus structure.
func AddAbstractedCorpus() *AbstractedCorpus {
	return &AbstractedCorpus{
		Abstractions: [][]LexicalClass{},
	}
}

// Print displays the tokenized corpus, sentence by sentence.
func (c *AbstractedCorpus) Print() {
	for _, sentence := range c.Abstractions {
		for _, lex := range sentence {
			fmt.Print(lex.class, " ")
		}
		fmt.Println()
	}
}

// AbstractCorpus creates an PositionalAbstraction from a Corpus. Assumes that.
// Corpus is sorted. A PositionalAbstraction is an AbstractedCorpus whose words
// are assigned a LexicalClass on its position and length of sentence.
// Ex: "Hello Go Lang" -> P31 P32 P33 (in Pnm, m denotes the m-th word of a
// sentence of length n; P denotes a Positional LexicalClass).
func (c *Corpus) PositionalAbstraction(d *Dictionary, lexes []LexicalClass) *AbstractedCorpus {
	ac := AddAbstractedCorpus()
	lcm := make(LexicalClassMap)
	lexCounter := 1

	for _, sentence := range c.Sentences {
		length := len(sentence)
		var as []LexicalClass

		for _, word := range sentence {
			if _, exists := d.entries[word]; !exists {
				lex := AddLexicalClass("P" + strconv.Itoa(length) + strconv.Itoa(lexCounter))
				lex.AddWord(word)
				lcm.Add(lex.class, word)
				lexCounter++

				d.AddEntry(word, lex)

				as = append(as, *lex)
			} else {
				as = append(as, *d.entries[word])
			}
		}

		// if !ac.IsAbstractionUnique(as) {
		ac.Abstractions = append(ac.Abstractions, as)
		// }

		lexCounter = 1
	}
	return ac
}

// ac.IsAbstractionUnique checks if given in input abstraction, it checks if that
// abstraction exists in the corpus already (removed redundancy)
func (ac *AbstractedCorpus) IsAbstractionUnique(as []LexicalClass) bool {
	for _, abstract := range ac.Abstractions {
		if len(abstract) != len(as) {
			continue
		}

		match := true
		for i := range abstract {
			if abstract[i].class != as[i].class {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}
	return false
}
