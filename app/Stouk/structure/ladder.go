package structure

type PlayerData struct {
	Username  string
	ID        int
	Rank      int
	MMR       int
	Win       int
	Lose      int
	RankMoyen int
	Rank_picture string
}


type Scoreboard struct {
	Players []PlayerData
}