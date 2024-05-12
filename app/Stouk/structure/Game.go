package structure

type DiceGame struct {
	Dice1 Dice
	Dice2 Dice
	Dice3 Dice
	Dice4 Dice
	Dice5 Dice
}

type GameStorage struct {
	Player1_Win int
	Player2_Win int
}

type Game_Result struct {
	PlayerName      string
	OpponentName 	string
	Game_res        GameStorage
	Dice_type       DiceGame
	Opponent_Dice	DiceGame
	Player_data     PlayerData
	Your_Game_roll  []int
	Enemy_Game_roll []int
}

type FrontPage struct {
	Game Game_Result
	DicePath DicePath
}
