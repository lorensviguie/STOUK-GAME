package data

import (
	"fmt"
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

// func GetScoreBoard() structure.PlayerData {
// 	db := GetDatabase()
// 	var Player_data structure.PlayerData
// 	err := db.QueryRow("SELECT ID_User, Rank, MMR FROM LADDER").Scan(&Player_data.ID, &Player_data.Rank, &Player_data.MMR)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = db.QueryRow("SELECT Win, Lose, RANK_MOYEN FROM RATIO").Scan(&Player_data.Win, &Player_data.Lose, &Player_data.RankMoyen)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = db.QueryRow("SELECT Username FROM USERS", Player_data.ID).Scan(&Player_data.Username)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return Player_data
// }

func GetAllScoreBoard() structure.Scoreboard {
	db := GetDatabase()
	rows, err := db.Query("SELECT USERS.Username, LADDER.Rank, RATIO.Win, RATIO.Lose, LADDER.Rank_picture FROM LADDER JOIN RATIO ON LADDER.ID_USER = RATIO.ID_USER JOIN USERS ON LADDER.ID_USER = USERS.ID ORDER BY LADDER.Rank DESC")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var scoreboard structure.Scoreboard
	for rows.Next() {
		var playerData structure.PlayerData
		err := rows.Scan(&playerData.Username, &playerData.Rank, &playerData.Win, &playerData.Lose, &playerData.Rank_picture)
		if err != nil {
			panic(err)
		}
		scoreboard.Players = append(scoreboard.Players, playerData)
	}
	return scoreboard
}

func UpdatePictureRank(id int, rank int) {
	fmt.Println("Updating picture rank")
	db := GetDatabase()
	var err error
	switch {
	case rank >= 0 && rank < 400:
		_, err = db.Exec("UPDATE LADDER SET Rank_picture = './static/images/logo/fer.png' WHERE ID_USER = ?", id)
	case rank >= 400 && rank < 800:
		_, err = db.Exec("UPDATE LADDER SET Rank_picture = './static/images/logo/bronze.png' WHERE ID_USER = ?", id)
	case rank >= 800 && rank < 1200:
		_, err = db.Exec("UPDATE LADDER SET Rank_picture = './static/images/logo/argent.png' WHERE ID_USER = ?", id)
	case rank >= 1200 && rank < 1600:
		_, err = db.Exec("UPDATE LADDER SET Rank_picture = './static/images/logo/gold.png' WHERE ID_USER = ?", id)
	case rank >= 1600 && rank < 2000:
		_, err = db.Exec("UPDATE LADDER SET Rank_picture = './static/images/logo/platine.png' WHERE ID_USER = ?", id)
	case rank >= 2000 && rank < 2400:
		_, err = db.Exec("UPDATE LADDER SET Rank_picture = './static/images/logo/emeraude.png' WHERE ID_USER = ?", id)
	case rank >= 2400 && rank < 2800:
		_, err = db.Exec("UPDATE LADDER SET Rank_picture = './static/images/logo/Diamond.png' WHERE ID_USER = ?", id)
	case rank >= 2800 && rank < 3200:
		_, err = db.Exec("UPDATE LADDER SET Rank_picture = './static/images/logo/Master.png' WHERE ID_USER = ?", id)
	case rank >= 3200 && rank < 3600:
		_, err = db.Exec("UPDATE LADDER SET Rank_picture = './static/images/logo/Grandmaitre.png' WHERE ID_USER = ?", id)
	case rank >= 3600:
		_, err = db.Exec("UPDATE LADDER SET Rank_picture = './static/images/logo/challenger.png' WHERE ID_USER = ?", id)
	default:
		panic("Invalid rank")
	}
	if err != nil {
		panic(err)
	}
}
