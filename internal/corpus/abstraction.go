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
// Ex: "Hello Go Lang" -> Z31 Z32 Z33 (in Znm, m denotes the m-th word of a
// sentence of length n).
func (c *Corpus) PositionalAbstraction() *AbstractedCorpus {
	d := AddDictionary()
	ac := AddAbstractedCorpus()

	lexCounter := 1

	for _, sentence := range c.Sentences {
		length := len(sentence)
		var abstractedSentence []LexicalClass
		for _, word := range sentence {
			if _, exists := d.entries[word]; !exists {
				lex := AddLexicalClass("Z" + strconv.Itoa(length) + strconv.Itoa(lexCounter))
				lexCounter++

				d.AddEntry(word, lex)

				abstractedSentence = append(abstractedSentence, *lex)
			} else {
				abstractedSentence = append(abstractedSentence, *d.entries[word])
			}
		}

		ac.Abstractions = append(ac.Abstractions, abstractedSentence)
		lexCounter = 1
	}

	return ac
}
