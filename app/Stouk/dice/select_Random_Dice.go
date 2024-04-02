package dice

import (
	"math/rand"
	"structure"
	"time"
)

func Select5RandomDice(ListDice []structure.Dice) []structure.Dice {
	rand.Seed(time.Now().UnixNano())
	selectedDice := make([]structure.Dice, 0)
	totalDice := len(ListDice)
	if totalDice < 5 {
		return nil
	}

	selectedIndices := make(map[int]bool)

	for len(selectedDice) < 5 {
		randomIndex := rand.Intn(totalDice)
		if selectedIndices[randomIndex] {
			continue
		}
		selectedDice = append(selectedDice, ListDice[randomIndex])
		selectedIndices[randomIndex] = true
	}

	return selectedDice
}
