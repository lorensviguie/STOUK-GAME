package data

import (
	"database/sql"
	"fmt"
	"path/filepath"
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
	}

	return playerStat
}

func UpdateAllPlayerdataForGame(playerData structure.PlayerData) (err error) {
	db := GetDatabase()
	queryLadder := "UPDATE LADDER SET Rank = ?, MMR = ? WHERE ID_USER = ?;"
	stmtLadder, err := db.Prepare(queryLadder)
	if err != nil {
		return fmt.Errorf("error preparing ladder update statement: %w", err)
	}
	defer stmtLadder.Close()

	queryRatio := "UPDATE RATIO SET WIN = ?, Lose = ?, RANK_MOYEN = ? WHERE ID_USER = ?;"
	stmtRatio, err := db.Prepare(queryRatio)
	if err != nil {
		return fmt.Errorf("error preparing ratio update statement: %w", err)
	}
	defer stmtRatio.Close()

	_, err = stmtLadder.Exec(playerData.Rank, playerData.MMR, playerData.ID)
	if err != nil {
		return fmt.Errorf("error updating ladder data: %w", err)
	}

	_, err = stmtRatio.Exec(playerData.Win, playerData.Lose, playerData.RankMoyen, playerData.ID)
	if err != nil {
		return fmt.Errorf("error updating ratio data: %w", err)
	}
	queryRatio = "UPDATE USERS "

	UpdatePictureRank(playerData.ID, playerData.Rank)
	IncreaseBalance(playerData.ID)
	return nil
}

func IncreaseBalance(userID int) error {
	db := GetDatabase()
	_, err := db.Exec("UPDATE USERS SET balance = balance + 100 WHERE ID = ?", userID)
	if err != nil {
		return fmt.Errorf("error increasing balance: %w", err)
	}
	return nil
}

func UpdatePictureRank(userID int, rank int) error {
	rankThresholds := []int{0, 401, 801, 1201, 1601, 2001, 2401, 2801, 3201, 3601}
	rankImages := []string{"fer.png", "bronze.png", "argent.png", "gold.png", "platine.png", "emeraude.png", "Diamond.png", "Master.png", "Grandmaitre.png", "challenger.png"}
	var rankImage string
	for i, threshold := range rankThresholds {
		if rank < threshold {
			rankImage = rankImages[i-1]
			break
		}
	}
	imagePath := filepath.Join("./static/images/logo/", rankImage)
	db := GetDatabase()
	_, err := db.Exec("UPDATE LADDER SET RANK_picture = ? WHERE ID_USER = ?", imagePath, userID)
	if err != nil {
		return fmt.Errorf("error updating rank picture: %w", err)
	}
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
	}

	return playerStat
}
