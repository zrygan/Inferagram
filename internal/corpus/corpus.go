package corpus

import (
	"fmt"
	"os"
	"strings"
)

type Corpus struct {
	clines []string // the lines of the corpus
	clen   int      // the numer of lines in the corpus
}

func ReadCorpus(filename string) *Corpus {
	corpus, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Inferagram ~\t File reading error:", err)
		return nil
	}

	clines := strings.Split(string(corpus), "\n")

	return NewCorpus(clines, len(clines))
}

func NewCorpus(clines []string, clen int) *Corpus {
	return &Corpus{
		clines: clines,
		clen:   len(clines),
	}
}

func PrintCorpus(c *Corpus) {
	for _, line := range c.clines {
		fmt.Println(line)
	}
}
