package data

import (
	"database/sql"
	"structure"
)

// GetUserDice récupère l'ensemble des dés d'un joueur par son ID utilisateur
func GetUserDice(userID int) (structure.Dice, error) {
	var db = GetDatabase()
	var Player_dice structure.Dice
	rows, err := db.Query("ID_DICE, Rank FROM USER_DICE WHERE ID_USER = ?", userID)
	if err != nil {
		return Player_dice, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&Player_dice.Dice, &Player_dice.Rank)
		if err != nil {
			return Player_dice, err
		}
	}

	if err := rows.Err(); err != nil {
		return Player_dice, err
	}

	return Player_dice, nil
}

// GetDiceByID récupère les détails d'un dé par son ID
func GetDiceByID(db *sql.DB, diceID int) (structure.Dice, error) {
	var dice structure.Dice
	err := db.QueryRow("SELECT Name, Price FROM DICE WHERE ID = ?", diceID).Scan(&dice.Dice, &dice.Rank)
	if err != nil {
		return structure.Dice{}, err
	}
	return dice, nil
}
