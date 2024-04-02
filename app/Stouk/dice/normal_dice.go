package dice

import (
	"math/rand"
	"time"
)

func Roll_NormalDice(rank int) int {
	rank++
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(20) + 1
}
