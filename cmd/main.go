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
		c := corpus.ReadFromFile(os.Args[1])
		c.Sort()
		ac := c.PositionalAbstraction()
		ac.Print()
	} else {
		fmt.Println("Inferagram ~\t More than one corpus provided")
	}
}
