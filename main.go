package main

import (
	"fmt"

	"github.com/berisberis/goching/pkg/hexgen"
)

func main() {
	hexagram := hexgen.GenerateHexagram()
	fmt.Println(hexagram)
}
