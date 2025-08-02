package main

import (
	"fmt"
	"os"

	"github.com/zrygan/Inferagram/internal/corpus"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Inferagram ~\t Corpus filename is not provided")
	} else if len(os.Args) == 2 {
		corpus_struct := corpus.ReadCorpus(os.Args[1])

		corpus.PrintCorpus(corpus_struct)
	} else {
		fmt.Println("Inferagram ~\t More than one corpus provided")
	}
}
