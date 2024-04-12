package data

import (
	"database/sql"
	"fmt"
	"structure"
)

func UpdatePlayerDice(username string, diceName string, rank int) error {
	// Obtenir l'ID de l'utilisateur
	userID, err := GetUserIDByUsername(username)
	if err != nil {
		return err
	}

	// Obtenir l'ID du dé
	diceID, err := GetDiceIDByName(diceName)
	if err != nil {
		return err
	}

	// Vérifier si l'utilisateur possède déjà le dé
	var existingRank int
	err = DB.QueryRow("SELECT Rank FROM USER_DICE WHERE ID_USER = ? AND ID_DICE = ?", userID, diceID).Scan(&existingRank)
	if err != nil {
		if err == sql.ErrNoRows { // L'utilisateur ne possède pas encore le dé, donc nous devons l'ajouter
			_, err := DB.Exec("INSERT INTO USER_DICE (ID_USER, ID_DICE, Rank) VALUES (?, ?, ?)", userID, diceID, rank)
			if err != nil {
				return err
			}
			fmt.Printf("Nouveau dé ajouté à l'utilisateur %s avec le rang %d\n", username, rank)
		} else {
			return err
		}
	} else {
		_, err := DB.Exec("UPDATE USER_DICE SET Rank = ? WHERE ID_USER = ? AND ID_DICE = ?", rank, userID, diceID)
		if err != nil {
			return err
		}
		fmt.Printf("Rang du dé mis à jour pour l'utilisateur %s: %d\n", username, rank)
	}

	return nil
}

func GetUserDice(userID int) ([]structure.Dice, error) {
	var db = GetDatabase()
	var playerDice []structure.Dice

	rows, err := db.Query("SELECT ID_DICE, Rank FROM USER_DICE WHERE ID_USER = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dice structure.Dice
		err := rows.Scan(&dice.Dice, &dice.Rank)
		if err != nil {
			return nil, err
		}
		playerDice = append(playerDice, dice)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return playerDice, nil
}

func GetDiceByID(db *sql.DB, diceID int) (structure.Dice, error) {
	var dice structure.Dice
	err := db.QueryRow("SELECT Name, Price FROM DICE WHERE ID = ?", diceID).Scan(&dice.Dice, &dice.Rank)
	if err != nil {
		return structure.Dice{}, err
	}
	return dice, nil
}
