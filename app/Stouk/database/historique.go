package data

import (
	"database/sql"
	"fmt"
	"structure"
)

func Histo_Add_Game(game structure.FindGameAndResult) {
	player1_MOD := getPlayerNewRank(game.Player1res.Player_data.ID) - game.Player1res.Player_data.Rank
	player2_MOD := getPlayerNewRank(game.Player2res.Player_data.ID) - game.Player2res.Player_data.Rank
	var playerWin = 0
	if game.Player1res.Game_res.Player1_Win < game.Player1res.Game_res.Player2_Win {
		playerWin = game.Player1res.Player_data.ID
	} else {
		playerWin = game.Player2res.Player_data.ID
	}
	_, err := DB.Exec("INSERT INTO HISTORIQUE (ID_PLAYER_1, ID_PLAYER_2, ID_PLAYER_WIN, PLAYER_1_WIN, PLAYER_2_WIN, PLAYER_1_LP_MOD, PLAYER_2_LP_MOD) VALUES (?, ?, ?, ?, ?, ?, ?)", game.Player1res.Player_data.ID, game.Player2res.Player_data.ID, playerWin, game.Player1res.Game_res.Player1_Win, game.Player1res.Game_res.Player2_Win, player1_MOD, player2_MOD)
	if err != nil {
		//return err
	}
	fmt.Printf("partie stocké avec succes \n")
}

func getPlayerNewRank(userID int) int {
	var playerrank = 0
	queryLadder := `
        SELECT Rank
        FROM LADDER
        WHERE ID_USER = ?
    `
	err := DB.QueryRow(queryLadder, userID).Scan(&playerrank)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("Erreur lors de la récupération des statistiques de LADDER:", err)
		}
	}
	return playerrank
}

func GetALLDataForHistorique(account structure.Account) structure.Compte {
	nbWin := 0
	nbLose := 0
	queryLadder := `
	SELECT Win, Lose
	FROM RATIO
	WHERE ID_USER = ?
	`
	err := DB.QueryRow(queryLadder, account.Id).Scan(&nbWin, &nbLose)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("Erreur lors de la récupération des statistiques de LADDER:", err)
		}
	}

	var playerrank = 0
	queryLadder = `
	SELECT Rank
	FROM LADDER
	WHERE ID_USER = ?
	`
	err = DB.QueryRow(queryLadder, account.Id).Scan(&playerrank)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("Erreur lors de la récupération des statistiques de RANK:", err)
		}
	}
	histo, _ := GetHistoriqueByIDPlayer(account.Id)
	var Compte structure.Compte
	Compte.Account = account
	Compte.NbLoose = nbLose
	Compte.NbWin = nbWin
	Compte.Rank = playerrank
	Compte.Historique = histo
	return Compte
}

func GetHistoriqueByIDPlayer(userID int) ([]structure.Historique, error) {
	var historiqueList []structure.Historique

	rows, err := DB.Query(`
		SELECT
			H.ID_PLAYER_1,
			H.ID_PLAYER_2,
			H.ID_PLAYER_WIN,
			H.PLAYER_1_WIN,
			H.PLAYER_2_WIN,
			H.PLAYER_1_LP_MOD,
			H.PLAYER_2_LP_MOD,
			U1.Username,
			U2.Username
		FROM HISTORIQUE H
		INNER JOIN USERS U1 ON H.ID_PLAYER_1 = U1.ID
		INNER JOIN USERS U2 ON H.ID_PLAYER_2 = U2.ID
		WHERE H.ID_PLAYER_1 = ? OR H.ID_PLAYER_2 = ?
	`, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var historique structure.Historique
		var idPlayer1, idPlayer2 int
		err := rows.Scan(&idPlayer1, &idPlayer2, &historique.ID_Player_Win, &historique.PLAYER_1_WIN, &historique.PLAYER_2_WIN, &historique.PLAYER_1_LP_MOD, &historique.PLAYER_2_LP_MOD, &historique.Player1_Username, &historique.Player2_Username)
		if err != nil {
			return nil, err
		}
		// Déterminer si le joueur actuel correspond à idPlayer1 ou idPlayer2
		if idPlayer1 == userID {
			historique.Who = true // Joueur 1
		} else {
			historique.Who = false // Joueur 2
		}
		historiqueList = append(historiqueList, historique)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return historiqueList, nil
}
