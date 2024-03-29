package dice

import (
	"math/rand"
	"time"
)

// D2 1 = D4 2 = D6 3 = D8 4 = D10 5 = D12 6 = D14 7 = D16 8 = D18 9 = D20 10

func rollDice(rank int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(rank*2) + 1
}