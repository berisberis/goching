package main

import (
	"fmt"

	"github.com/berisberis/goching/pkg/hexgen"
)

func main() {
	hexagrams := hexgen.GenerateHexagrams(100000)
	mostCommonHexagram, count := hexgen.MostCommonHexagram(hexagrams)
	fmt.Printf("Most common hexagram:\n%v\nCount: %v\n", mostCommonHexagram, count)
}
