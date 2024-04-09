package ladder

import (
	"data"
	"fmt"
	"structure"
)

//le rank il va de 0 a autant que le veut mais il doit avoir un max theorique de 4500 a plus on monte plus c'est dur de monté
//data : 0 rank / 1 mmr / 2 nombre  win / 3 nombre lose / 5 rankmoyen
// de base on perd 20 on gagne 20

func UpdateRankforPlayer(playerData structure.PlayerData, win bool) {
	if win {
		playerData.Rank = GainLP(playerData)
	} else {
		playerData.Rank = LooseLP(playerData)
	}
	data.UpdateAllPlayerdataForGame(playerData)
}

func GainLP(player structure.PlayerData) int {
	lpChange := 20
	coef := calculateMultiplierCoefficientForWin(player.Rank, player.MMR, player.RankMoyen)
	fmt.Println(float64(coef))
	player.Rank += int(float64(lpChange) * coef)
	return player.Rank
}

// Perd des LP en fonction des résultats
func LooseLP(player structure.PlayerData) int {
	lpChange := 20
	coef := calculateMultiplierCoefficientForLoss(player.Rank, player.MMR, player.RankMoyen)
	player.Rank -= int(float64(lpChange) * coef)
	return player.Rank
}
