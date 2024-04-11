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

func UpdateAllPlayerdataForGame(playerData structure.PlayerData) (err error) {
	db := GetDatabase()

	// Print player data for verification
	fmt.Printf("Updating player data for user %d:\n", playerData.ID)
	fmt.Printf("  Rank: %d\n", playerData.Rank)
	fmt.Printf("  MMR: %d\n", playerData.MMR)
	fmt.Printf("  Win: %d\n", playerData.Win)
	fmt.Printf("  Lose: %d\n", playerData.Lose)
	fmt.Printf("  RankMoyen: %.2f\n", playerData.RankMoyen)
	// Prepare update statements with named parameters
	queryLadder := "UPDATE LADDER SET Rank = ?, MMR = ? WHERE ID_USER = ?;"
	stmtLadder, err := db.Prepare(queryLadder)
	if err != nil {
		return fmt.Errorf("error preparing ladder update statement: %w", err)
	}
	defer stmtLadder.Close() // Ensure statement is closed even in case of errors

	queryRatio := "UPDATE RATIO SET WIN = ?, Lose = ?, RANK_MOYEN = ? WHERE ID_USER = ?;"
	stmtRatio, err := db.Prepare(queryRatio)
	if err != nil {
		return fmt.Errorf("error preparing ratio update statement: %w", err)
	}
	defer stmtRatio.Close() // Ensure statement is closed even in case of errors

	// Execute update statements with error handling
	_, err = stmtLadder.Exec(playerData.Rank, playerData.MMR, playerData.ID)
	if err != nil {
		return fmt.Errorf("error updating ladder data: %w", err)
	}

	_, err = stmtRatio.Exec(playerData.Win, playerData.Lose, playerData.RankMoyen, playerData.ID)
	if err != nil {
		return fmt.Errorf("error updating ratio data: %w", err)
	}

	fmt.Println("Player data updated successfully!")
	return nil
}


func GetAllPlayerDataForQueue(userID int) structure.PlayerData {
	var playerStat structure.PlayerData
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
