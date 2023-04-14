package hexgen

import (
	"crypto/rand"
	"fmt"
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

func GenerateHexagrams(n int) []Hexagram {
	hexagrams := make([]Hexagram, n)
	for i := 0; i < n; i++ {
		hexagrams[i] = GenerateHexagram()
	}
	return hexagrams
}

func CountOccurrences(hexagrams []Hexagram) map[string]int {
	occurrences := make(map[string]int)
	for _, hexagram := range hexagrams {
		key := hexagram.String()
		if _, ok := occurrences[key]; ok {
			occurrences[key]++
		} else {
			occurrences[key] = 1
		}
	}
	return occurrences
}

func MostCommonHexagram(hexagrams []Hexagram) (string, int) {
	occurrences := CountOccurrences(hexagrams)
	mostCommonHexagram := ""
	mostCommonCount := 0
	for key, count := range occurrences {
		if count > mostCommonCount {
			mostCommonHexagram = key
			mostCommonCount = count
		}
	}
	return mostCommonHexagram, mostCommonCount
}

func (t Trigram) String() string {
	return fmt.Sprintf("%v%v%v", t.FirstLine, t.SecondLine, t.ThirdLine)
}

func (h Hexagram) String() string {
	return fmt.Sprintf("%v%v", h.UpperTrigram, h.LowerTrigram)
}
