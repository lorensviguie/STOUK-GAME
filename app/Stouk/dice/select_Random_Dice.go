package dice

import (
	"math/rand"
	"structure"
	"time"
)

func Select5RandomDice(ListDice structure.Dice) []structure.Dice {
	rand.Seed(time.Now().UnixNano())
	selectedDice := make([]structure.Dice, 0)
	totalDice := ListDice.Dice
	for i := 0; i < 5; i++ {
		randomIndex := rand.Intn(totalDice)
		selectedDice = append(selectedDice, structure.Dice{Dice: ListDice.Dice, Rank: randomIndex})
		totalDice--
	}

	return selectedDice
}
