package data

import (
	"structure"

	_ "github.com/go-sql-driver/mysql"
)

func GetForUpdateLadder(idPlayer int) structure.PlayerData {
	db := GetDatabase()
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

func GetAllScoreBoard() structure.Scoreboard {
	db := GetDatabase()
	rows, err := db.Query("SELECT USERS.Username, LADDER.Rank, RATIO.Win, RATIO.Lose FROM LADDER JOIN RATIO ON LADDER.ID_USER = RATIO.ID_USER JOIN USERS ON LADDER.ID_USER = USERS.ID")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var scoreboard structure.Scoreboard
	for rows.Next() {
		var playerData structure.PlayerData
		err := rows.Scan(&playerData.Username, &playerData.Rank, &playerData.Win, &playerData.Lose)
		if err != nil {
			panic(err)
		}
		scoreboard.Players = append(scoreboard.Players, playerData)
	}
	return scoreboard
}