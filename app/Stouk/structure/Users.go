package structure

type Account struct {
	Id            int
	Username      string
	Email         string
	Password      string
	Balance       int
	ProfilPicture string
	CreationDate  string
}

type Historique struct {
	Who              bool
	Player1_Username string
	Player2_Username string
	ID_Player_Win    int
	PLAYER_1_WIN     int
	PLAYER_2_WIN     int
	PLAYER_1_LP_MOD  int
	PLAYER_2_LP_MOD  int
}

type AdminData struct {
	Users []Account
}

type StoreData struct {
	Dices []Dice
	User  []Account
}

type Compte struct {
	Account    Account
	NbWin      int
	NbLoose    int
	Rank       int
	Historique []Historique
}
