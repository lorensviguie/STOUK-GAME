package data

import (
	"errors"
	"fmt"
)

func Add_Dice(name string, price int) (int64, error) {
	db := GetDatabase()

	// Vérifier si le nom existe déjà
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM DICE WHERE Name = ?", name).Scan(&count)
	if err != nil {
		return 0, err
	}

	// Si le nom existe déjà, renvoyer une erreur
	if count > 0 {
		return 0, errors.New("Le nom existe déjà dans la base de données")
	}

	// Insérer le dé s'il n'existe pas déjà
	result, err := db.Exec("INSERT INTO DICE (Name, Price) VALUES (?, ?)", name, price)
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
    db := GetDatabase() // Supposons que vous ayez une fonction GetDatabase() qui retourne une connexion à la base de données

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