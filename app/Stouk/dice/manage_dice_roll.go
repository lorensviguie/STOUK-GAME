package dice

import "structure"

func Manage_Dice_Roll(Dice structure.Dice)int{
	switch Dice.Dice {
	case 1:
		return RollBaseDice(Dice.Rank)
	case 2:	
	}
	return 0
}