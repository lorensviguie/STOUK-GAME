package data

import (
	"errors"
	"fmt"
)

func Add_Dice(name string, price int, path string) (int64, error) {
	db := GetDatabase()
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM DICE WHERE Name = ?", name).Scan(&count)
	if err != nil {
		return 0, err
	}

	if count > 0 {
		return 0, errors.New("Le nom existe déjà dans la base de données")
	}

	result, err := db.Exec("INSERT INTO DICE (Name, Price, Path) VALUES (?, ?, ?)", name, price, path)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Printf("Dé ajouté avec succès. ID du nouveau dé : %d\n", lastInsertedID)
	return lastInsertedID, nil
}

func GetDiceIDByName(name string) (int64, error) {
	db := GetDatabase()

	var diceID int64
	err := db.QueryRow("SELECT ID FROM DICE WHERE Name = ?", name).Scan(&diceID)
	if err != nil {
		if err != nil {
			return 0, fmt.Errorf("dé avec le nom %s non trouvé", name)
		}
		return 0, err
	}

	return diceID, nil
}

func GetDicePathWithID(id int) (string,error) {
	db := GetDatabase()

	var dicePath string
	err := db.QueryRow("SELECT Path FROM DICE WHERE ID = ?", id).Scan(&dicePath)
	if err != nil {
		if err != nil {
			return "", fmt.Errorf("dé avec le id %s non trouvé", id)
		}
		return "", err
	}

	return dicePath, nil
}
