package data

import (
	"database/sql"
	"fmt"
	"structure"
)

func GetAllPlayerDataForGame(username string) structure.PlayerData {
	var playerStat structure.PlayerData
	userID, _ := GetUserIDByUsername(username)
	playerStat.ID = int(userID)
	db := GetDatabase()

	// Récupération des statistiques de la table LADDER
	queryLadder := `
        SELECT Rank, MMR
        FROM LADDER
        WHERE ID_USER = ?
    `
	err := db.QueryRow(queryLadder, userID).Scan(&playerStat.Rank, &playerStat.MMR)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("Erreur lors de la récupération des statistiques de LADDER:", err)
		}
		// Vous pouvez gérer l'absence de données ici si nécessaire
	}

	// Récupération du rang moyen de la table RATIO
	queryRatio := `
        SELECT Win, Lose, Rank_Moyen
        FROM RATIO
        WHERE ID_USER = ?
    `
	err = db.QueryRow(queryRatio, userID).Scan(&playerStat.Win, &playerStat.Lose, &playerStat.RankMoyen)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("Erreur lors de la récupération du rang moyen de RATIO:", err)
		}
		// Vous pouvez gérer l'absence de données ici si nécessaire
	}

	return playerStat
}
