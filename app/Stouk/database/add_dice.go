package data

import "fmt"

func Add_Dice(name string, price int) (int64, error) {
	db := GetDatabase()
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
