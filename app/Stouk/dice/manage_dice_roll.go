package dice

import "structure"

func Manage_Dice_Roll(Dice structure.Dice) int {
	switch Dice.Dice {
	case 1:
		return Roll_Base_Dice(Dice.Rank)
	case 2:
		return Roll_NormalDice(Dice.Rank)
	case 3:
		return Roll_Parabole_Dice(Dice.Rank)
	case 4:
		return Roll_Power_Dice(Dice.Rank)
	case 5:
		return Roll_Scaledice(Dice.Rank)
	case 6:
		return Roll_unscaledice(Dice.Rank)
	case 7:
		return Roll_RankDice(Dice.Rank)
	}
	return 0
}
