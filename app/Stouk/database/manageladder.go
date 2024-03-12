package data

import (
	"structure"

	_ "github.com/go-sql-driver/mysql"
)

func GetForUpdateLadder(idPlayer int) structure.PlayerData {
	db := GetDatabase()
	defer db.Close()
	var Player_data structure.PlayerData
	err := db.QueryRow("SELECT Rank, MMR FROM LADDER WHERE ID_USER = ?", idPlayer).Scan(&Player_data.Rank, &Player_data.MMR)
	if err != nil {
		panic(err)
	}
	err = db.QueryRow("SELECT Win, Lose, RANK_MOYEN FROM RATIO WHERE ID_USER = ?", idPlayer).Scan(&Player_data.Win, &Player_data.Lose, &Player_data.RankMoyen)
	if err != nil {
		panic(err)
	}
	return Player_data
}
