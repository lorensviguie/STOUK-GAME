package dice

import (
	"data"
	"structure"
)

func BuildDicePathForGame(DiceRes structure.Game_Result) structure.DicePath {
	var dicepath structure.DicePath
	diceTypes := []structure.Dice{DiceRes.Dice_type.Dice1, DiceRes.Dice_type.Dice2, DiceRes.Dice_type.Dice3, DiceRes.Dice_type.Dice4, DiceRes.Dice_type.Dice5}
	opponentDiceTypes := []structure.Dice{DiceRes.Opponent_Dice.Dice1, DiceRes.Opponent_Dice.Dice2, DiceRes.Opponent_Dice.Dice3, DiceRes.Opponent_Dice.Dice4, DiceRes.Opponent_Dice.Dice5}
	dicePaths := []*string{&dicepath.Pathdice1, &dicepath.Pathdice2, &dicepath.Pathdice3, &dicepath.Pathdice4, &dicepath.Pathdice5}
	opponentDicePaths := []*string{&dicepath.Pathdice6, &dicepath.Pathdice7, &dicepath.Pathdice8, &dicepath.Pathdice9, &dicepath.Pathdice10}
	for i, dice := range diceTypes {
		id := dice.Dice
		*dicePaths[i], _ = data.GetDicePathWithID(id)
	}
	for i, dice := range opponentDiceTypes {
		id := dice.Dice
		*opponentDicePaths[i], _ = data.GetDicePathWithID(id)
	}

	return dicepath
}
