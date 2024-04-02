package dice

import (
	"data"
	"fmt"
	"structure"
)

func Dice_Game(Player1, Player2 structure.PlayerData) (structure.Game_Result, structure.Game_Result) {
	var result structure.GameStorage

	var Player1_Res structure.Game_Result
	Player1_Res.Player_data = Player1

	var Player2_Res structure.Game_Result
	Player2_Res.Player_data = Player2

	var Player1_Dice, _ = data.GetUserDice(Player1.ID)
	var Player2_Dice, _ = data.GetUserDice(Player2.ID)
	fmt.Println(Player1_Dice)
	fmt.Println(Player2_Dice)

	Player1_Game_Dice := Select5RandomDice(Player1_Dice)
	Player1_Res.Dice_type = Merge5combatDice(Player1_Game_Dice)

	Player2_Game_Dice := Select5RandomDice(Player2_Dice)
	Player2_Res.Dice_type = Merge5combatDice(Player2_Game_Dice)

	for i := 0; i < 5; i++ {
		Player1_Dice := Player1_Game_Dice[i]
		Player1_Roll := Manage_Dice_Roll(Player1_Dice)

		Player2_Dice := Player2_Game_Dice[i]
		Player2_Roll := Manage_Dice_Roll(Player2_Dice)
		for Player1_Roll == Player2_Roll {
			Player1_Roll = Manage_Dice_Roll(Player1_Dice)
			Player2_Roll = Manage_Dice_Roll(Player2_Dice)
		}

		Player1_Res.Your_Game_roll = append(Player1_Res.Your_Game_roll, Player1_Roll)
		Player1_Res.Enemy_Game_roll = append(Player1_Res.Enemy_Game_roll, Player2_Roll)

		Player2_Res.Your_Game_roll = append(Player2_Res.Your_Game_roll, Player2_Roll)
		Player2_Res.Enemy_Game_roll = append(Player2_Res.Enemy_Game_roll, Player1_Roll)

		if Player1_Roll > Player2_Roll {
			result.Player1_Win++
		} else {
			result.Player2_Win++
		}

	}
	fmt.Println(result)
	Player1_Res.Game_res = result
	Player2_Res.Game_res = result
	return Player1_Res, Player2_Res
}

func Merge5combatDice(MD []structure.Dice) structure.DiceGame {
	var res structure.DiceGame
	res.Dice1 = MD[0]
	res.Dice2 = MD[1]
	res.Dice3 = MD[2]
	res.Dice4 = MD[3]
	res.Dice5 = MD[4]
	return res
}
