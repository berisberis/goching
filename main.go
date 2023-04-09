package main

import (
	"fmt"

	"github.com/berisberis/goching/pkg/hexgen"
)

func main() {

	var topTrigram, bottomTrigram []int64

	for i := 0; i < 3; i++ {
		c, err := hexgen.CoinToss()
		if err != nil {
			fmt.Printf("error creating CoinToss: %v \n", err)
		}
		topTrigram = append(topTrigram, c)
	}

	top, err := hexgen.NewTrigram(topTrigram).CreateTrigram()
	if err != nil {
		fmt.Printf("error creating top Trigram: %v \n", err)
	}
	fmt.Printf("\n the top tri number is : %v", top)

	for i := 0; i < 3; i++ {
		c, err := hexgen.CoinToss()
		if err != nil {
			fmt.Printf("error creating CoinToss: %v \n", err)
		}
		bottomTrigram = append(bottomTrigram, c)
	}

	bottom, err := hexgen.NewTrigram(bottomTrigram).CreateTrigram()
	if err != nil {
		fmt.Printf("error creating bottom Trigram: %v \n", err)
	}
	fmt.Printf("\n the bottom tri number is : %v", bottom)

	h := hexgen.NewHexagram(top, bottom).Generate()
	fmt.Printf("\n the hex number is : %v \n", h)
}
