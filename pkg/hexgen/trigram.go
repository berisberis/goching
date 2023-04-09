package hexgen

import (
	"strconv"
)

const (
	earth    = iota // 0 - 000
	mountain        // 1 - 001
	water           // 2 - 010
	wind            // 3 - 011
	thunder         // 4 - 100
	fire            // 5 - 101
	lake            // 6 - 110
	heaven          // 7 - 111
)

type Trigram struct {
	trigram []int64
}

func (t Trigram) CreateTrigram() (int64, error) {

	var binaryString []byte

	for _, v := range t.trigram {
		binaryString = strconv.AppendInt(binaryString, v, 2)
	}

	trigramNumber, err := strconv.ParseInt(string(binaryString), 2, 64)

	if err != nil {
		return 2, err
	}

	return trigramNumber, nil
}

func NewTrigram(trigram []int64) *Trigram {
	return &Trigram{
		trigram: trigram,
	}
}
