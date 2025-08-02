package corpus

import (
	"fmt"
	"os"
	"strings"
)

type Corpus struct {
	Sentences [][]string
}

func ReadCorpus(filename string) *Corpus {
	corpus, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Inferagram ~\t File reading error:", err)
		return nil
	}

	return NewCorpus(string(corpus))
}

func NewCorpus(text string) *Corpus {
	var corpus [][]string

	lines := strings.Split(strings.TrimSpace(text), "\n")
	for _, line := range lines {
		words := strings.Fields(line)
		corpus = append(corpus, words)
	}

	return &Corpus{
		Sentences: corpus,
	}
}

func (c *Corpus) PrintCorpus() {
	for _, line := range c.Sentences {
		fmt.Println(strings.Join(line, " "))
	}
}
