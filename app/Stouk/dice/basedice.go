package dice

import (
	"math/rand"
	"time"
)

func RollBaseDice(rank int) int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1

	switch rank {
	case 1:
		// Iron: un dé 6 classique
		return (randomNumber % 6) + 1
	case 2:
		// Bronze: un dé 7
		return (randomNumber % 7) + 1
	case 3:
		// Silver: le pourcentage de 1 passe de 14.29 à 10
		if randomNumber <= 10 {
			return 1
		}
		return (randomNumber-10)/14 + 2
	case 4:
		// Gold: le pourcentage de 1 passe à 5%
		if randomNumber <= 5 {
			return 1
		}
		return (randomNumber-5)/19 + 2
	case 5:
		// Platinum: un dé 8
		return (randomNumber % 8) + 1
	case 6:
		// Emerald: le pourcentage des chiffres entre 1 et 3 passe à 5%
		if randomNumber <= 15 {
			return randomNumber%3 + 1
		}
		return (randomNumber-5)/16 + 4
	case 7:
		// Diamond: le dé devient un d9
		return (randomNumber % 9) + 1
	case 8:
		// Master: le dé devient un d10
		return (randomNumber % 10) + 1
	case 9:
		// Grand Master: le pourcentage des chiffres entre 1 et 5 passe à 5%
		if randomNumber <= 25 {
			return randomNumber%5 + 1
		}
		return (randomNumber-5)/19 + 6
	case 10:
		// Challenger: le pourcentage de faire 10 devient 25%
		if randomNumber <= 25 {
			return 10
		}
		return (randomNumber-25)/9 + 1
	default:
		return 0 // Gestion du cas par défaut
	}
}
