package corpus

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// Corpus represents a tokenized corpus — each sentence is a slice of tokens (words).
type Corpus struct {
	Sentences [][]string
}

// ReadFromFile reads the contents of a text file and returns a tokenized corpus.
func ReadFromFile(filename string) *Corpus {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Inferagram ~\t File reading error:", err)
		return nil
	}
	return New(strings.TrimSpace(string(data)))
}

// c.Sort sorts the corpus by the lengths of the sentences (in ascending order).
func (c *Corpus) Sort() {
	sort.Slice(c.Sentences, func(i, j int) bool {
		return len(c.Sentences[i]) < len(c.Sentences[j])
	})
}

// New constructs a new tokenized Corpus from a raw text string.
func New(text string) *Corpus {
	var sentences [][]string
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		words := strings.Fields(line)
		sentences = append(sentences, words)
	}
	return &Corpus{Sentences: sentences}
}

// Print displays the tokenized corpus, sentence by sentence.
func (c *Corpus) Print() {
	for _, sentence := range c.Sentences {
		fmt.Println(strings.Join(sentence, " "))
	}
}
