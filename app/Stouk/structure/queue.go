package structure

type Queue struct {
	ID int
	Variance int
	Username string
	Rank int
	Rank_Moyen int
}

type FindGameAndResult struct {
	Claim bool
	Player1U string
	Player2U string
	Player1res Game_Result
	Player2res Game_Result
}

var findResult []FindGameAndResult
var FindResult *[]FindGameAndResult = &findResult

var queueFile []Queue
var QueueFile *[]Queue = &queueFile
