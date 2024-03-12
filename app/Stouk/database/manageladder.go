package database

func GetForUpdateLadder(idPlayer int) []int {
	db := ConnectToDb()
	defer db.Close()
	var data []int
	err := db.QueryRow("SELECT Rank, MMR FROM LADDER WHERE ID_USER = ?", idPlayer).Scan(&data[0], &data[1])
	if err != nil {
		panic(err)
	}
	err = db.QueryRow("SELECT Win, Lose, RANK_MOYEN FROM RATIO WHERE ID_USER = ?", idPlayer).Scan(&data[2], &data[3], &data[4])
	if err != nil {
		panic(err)
	}
	return data
}
