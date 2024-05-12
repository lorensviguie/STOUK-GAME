package queue

import (
	"data"
	"dice"
	"fmt"
	"ladder"
	"logs"
	"structure"
	"time"
)

func ManageQueue() {
	fmt.Println("Traitement de la file d'attente...")
	var up structure.Queue = structure.Queue{ID: -1, Username: "up", Variance: 0, Rank: -10000, Rank_Moyen: -10000} //eviter des probleme de users seul et permet de laisser la Q alimenter en users
	var down structure.Queue = structure.Queue{ID: -1, Username: "down", Variance: 0, Rank: 1000000, Rank_Moyen: 10000000}
	*structure.QueueFile = append(*structure.QueueFile, up)
	*structure.QueueFile = append(*structure.QueueFile, down)
	antibrake := false
	for {
		for i, player := range *structure.QueueFile {
			(*structure.QueueFile)[i].Variance += 20
			if (*structure.QueueFile)[i].Variance > 400 {
				(*structure.QueueFile)[i].Variance = 0
			}
			for j, otherPlayer := range *structure.QueueFile {
				if player.ID != otherPlayer.ID && ((player.Rank+player.Variance >= otherPlayer.Rank && player.Rank <= otherPlayer.Rank) || (player.Rank-player.Variance <= otherPlayer.Rank && player.Rank >= otherPlayer.Rank)) {
					fmt.Printf("Retirer %s et %s de la file d'attente\n", player.Username, otherPlayer.Username)
					player1res, player2res := dice.Dice_Game(data.GetAllPlayerDataForGame(player.Username), data.GetAllPlayerDataForGame(otherPlayer.Username))
					var addToClaim structure.FindGameAndResult
					var histoNewGame = &addToClaim
					addToClaim.Claim = false
					addToClaim.Player1U = player.Username
					addToClaim.Player2U = otherPlayer.Username
					addToClaim.Player1res = player1res
					addToClaim.Player2res = player2res
					player1res.Player_data, player2res.Player_data = DiceApplyResult(player1res, player2res)
					*structure.QueueFile = append((*structure.QueueFile)[:i], (*structure.QueueFile)[i+1:]...)
					*structure.QueueFile = append((*structure.QueueFile)[:j-1], (*structure.QueueFile)[j:]...)
					data.Histo_Add_Game(*histoNewGame)
					*structure.FindResult = append(*structure.FindResult, addToClaim)
					antibrake = true
					break
				}
			}
			if antibrake {
				antibrake = false
				break
			}
		}
		//time.Sleep(1 * time.Second) pour gerer la viteese de la Q
	}
}

func CheckTagForUser(id int) structure.Game_Result {
	for {
		for i, game := range *structure.FindResult {
			if game.Player1res.Player_data.ID == id {
				claimPointer := &(*structure.FindResult)[i].Claim
				temp := game.Player1res
				temp.OpponentName = game.Player2U
				temp.Opponent_Dice = game.Player2res.Dice_type
				checkIFalreadyclaim(claimPointer, i)
				return temp
			} else if game.Player2res.Player_data.ID == id {
				time.Sleep(1 * time.Second)
				claimPointer := &(*structure.FindResult)[i].Claim
				temp := game.Player2res
				temp.OpponentName = game.Player1U
				temp.Opponent_Dice = game.Player1res.Dice_type
				checkIFalreadyclaim(claimPointer, i)
				return temp
			}
		}
	}
}

func checkIFalreadyclaim(claim *bool, i int) {
	fmt.Println("\n", claim)
	if !*claim {
		*claim = true
	} else {
		*structure.FindResult = append((*structure.FindResult)[:i], (*structure.FindResult)[i+1:]...)
	}
}

func ContainsID(id int) bool {
	for _, item := range *structure.QueueFile {
		if item.ID == id {
			return true
		}
	}
	return false
}

func Add_User_To_Queue(id int) {
	if ContainsID(id) {
		return
	}
	playerData := data.GetAllPlayerDataForQueue(id)
	newUserInQueue := structure.Queue{
		ID:         playerData.ID,
		Username:   data.GetUsernameByUserid(id),
		Rank:       playerData.Rank,
		Rank_Moyen: playerData.RankMoyen,
		Variance:   0,
	}
	*structure.QueueFile = append(*structure.QueueFile, newUserInQueue)
	logs.LogToFile("queue", fmt.Sprintln(playerData.ID, " Has been added to the Queue"))
}

func DiceApplyResult(player1res, player2res structure.Game_Result) (structure.PlayerData, structure.PlayerData) {
	print("\n\n--------------\nCOIN COIN COIN COIN\n----------------\n\n")
	if player1res.Game_res.Player1_Win > player1res.Game_res.Player2_Win {
		player1res.Player_data.Win += 1
		player2res.Player_data.Lose += 1
		player1res.Player_data.MMR += 20
		player2res.Player_data.MMR -= 20
		player1res.Player_data.RankMoyen = ((player1res.Player_data.Win+player1res.Player_data.Lose)*player1res.Player_data.RankMoyen + player2res.Player_data.Rank) / (player1res.Player_data.Win + player1res.Player_data.Lose + 1)
		player2res.Player_data.RankMoyen = ((player2res.Player_data.Win+player2res.Player_data.Lose)*player2res.Player_data.RankMoyen + player1res.Player_data.Rank) / (player2res.Player_data.Win + player2res.Player_data.Lose + 1)
		ladder.UpdateRankforPlayer(player1res.Player_data, true)
		ladder.UpdateRankforPlayer(player2res.Player_data, false)
	} else {
		player2res.Player_data.Win += 1
		player1res.Player_data.Lose += 1
		player1res.Player_data.MMR -= 20
		player2res.Player_data.MMR += 20
		player1res.Player_data.RankMoyen = ((player1res.Player_data.Win+player1res.Player_data.Lose)*player1res.Player_data.RankMoyen + player2res.Player_data.Rank) / (player1res.Player_data.Win + player1res.Player_data.Lose + 1)
		player2res.Player_data.RankMoyen = ((player2res.Player_data.Win+player2res.Player_data.Lose)*player2res.Player_data.RankMoyen + player1res.Player_data.Rank) / (player2res.Player_data.Win + player2res.Player_data.Lose + 1)
		ladder.UpdateRankforPlayer(player1res.Player_data, false)
		ladder.UpdateRankforPlayer(player2res.Player_data, true)
	}
	player1data := data.GetAllPlayerDataForGame(data.GetUsernameByUserid(player1res.Player_data.ID))
	player2data := data.GetAllPlayerDataForGame(data.GetUsernameByUserid(player2res.Player_data.ID))
	return player1data, player2data
}
