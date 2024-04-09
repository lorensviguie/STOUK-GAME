package dice

import (
	"math/rand"
	"time"
)

func Roll_Power_Dice(rank int) int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1

	switch rank {
	case 1:
		// Iron: un dé 3
		return (randomNumber % 3) + 1
	case 2:
		// Bronze: un dé 3 avec le res facteur de 2
		return ((randomNumber % 3) + 1) * 2
	case 3:
		// Silver: un dé 4 avec le res facteur de 2
		return ((randomNumber % 4) + 1) * 2
	case 4:
		// Gold: un dé 5 avec le res facteur de 2
		return ((randomNumber % 5) + 1) * 2
	case 5:
		// Platinum: un dé 4 au carré
		number := (randomNumber % 4) + 1
		return number * number
	case 6:
		// Emerald: dé platine avec +1 apres
		number := (randomNumber % 4) + 1
		return number*number + 1
	case 7:
		// Diamond: dé platine +2 apres
		number := (randomNumber % 4) + 1
		return number*number + 1
	case 8:
		// Master: un dé 5 au carré -5
		number := (randomNumber % 5) + 1
		return number*number - 5
	case 9:
		// Grand Master: un d2 au cube *2
		number := (randomNumber % 2) + 1
		return (number * number * number) * 2
	case 10:
		// Challenger: un dé GM +2
		number := (randomNumber % 2) + 1
		return (number*number*number)*2 + 2
	default:
		return 0 // Gestion du cas par défaut
	}
}
