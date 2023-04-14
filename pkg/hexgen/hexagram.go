package hexgen

import (
	"crypto/rand"
)

type Trigram struct {
	FirstLine  int
	SecondLine int
	ThirdLine  int
}

type Hexagram struct {
	UpperTrigram Trigram
	LowerTrigram Trigram
}

func TossCoin() int {
	b := make([]byte, 1)
	rand.Read(b)
	if b[0]%2 == 0 {
		return 1 // Heads
	} else {
		return 0 // Tails
	}
}

func GenerateTrigram() Trigram {
	return Trigram{TossCoin(), TossCoin(), TossCoin()}
}

func GenerateHexagram() Hexagram {
	trigram1 := GenerateTrigram()
	trigram2 := GenerateTrigram()
	return Hexagram{trigram1, trigram2}
}
