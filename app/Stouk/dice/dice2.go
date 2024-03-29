package dice

import (
	"math/rand"
	"time"
)

func dice_0(rank int) int{
	rand.Seed(time.Now().UnixNano())

	switch rank {
	case 1:
		dice := rand.Intn(6) + 1
		if (dice <= 5) {
			return 0
		} else{
			return 20
		}
	case 2:
		dice := rand.Intn(7) + 1
		if (dice <= 5) {
			return 0
		} else{
			return 20
		}
	case 3:
		dice := rand.Intn(8) + 1
		if (dice <= 5) {
			return 0
		} else{
			return 20
		}
	case 4:
		dice := rand.Intn(9) + 1
		if (dice <= 5) {
			return 0
		} else{
			return 20
		}
	case 5:
		dice := rand.Intn(10) + 1
		if (dice <= 5) {
			return 0
		} else{
			return 20
		}
	case 6:
		dice := rand.Intn(11) + 1
		if (dice <= 5) {
			return 0
		} else{
			return 20
		}
	case 7:
		dice := rand.Intn(12) + 1
		if (dice <= 5) {
			return 0
		} else{
			return 20
		}
	case 8:
		dice := rand.Intn(13) + 1
		if (dice <= 5) {
			return 0
		} else{
			return 20
		}
	case 9:
		dice := rand.Intn(14) + 1
		if (dice <= 5) {
			return 0
		} else{
			return 20
		}
	case 10:
		dice := rand.Intn(15) + 1
		if (dice <= 5) {
			return 0
		} else{
			return 20
		}
	default:
		return 0 // Gestion du cas par défaut
	}
	
}

